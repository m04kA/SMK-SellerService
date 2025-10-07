package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	"github.com/m04kA/SMK-SellerService/internal/api/handlers/create_company"
	"github.com/m04kA/SMK-SellerService/internal/api/handlers/create_service"
	"github.com/m04kA/SMK-SellerService/internal/api/handlers/delete_company"
	"github.com/m04kA/SMK-SellerService/internal/api/handlers/delete_service"
	"github.com/m04kA/SMK-SellerService/internal/api/handlers/get_company"
	"github.com/m04kA/SMK-SellerService/internal/api/handlers/get_service"
	"github.com/m04kA/SMK-SellerService/internal/api/handlers/list_companies"
	"github.com/m04kA/SMK-SellerService/internal/api/handlers/list_services"
	"github.com/m04kA/SMK-SellerService/internal/api/handlers/update_company"
	"github.com/m04kA/SMK-SellerService/internal/api/handlers/update_service"
	"github.com/m04kA/SMK-SellerService/internal/api/middleware"
	companyRepo "github.com/m04kA/SMK-SellerService/internal/infra/storage/company"
	serviceRepo "github.com/m04kA/SMK-SellerService/internal/infra/storage/service"
	companiesService "github.com/m04kA/SMK-SellerService/internal/service/companies"
	servicesService "github.com/m04kA/SMK-SellerService/internal/service/services"
	"github.com/m04kA/SMK-SellerService/pkg/logger"
)

func main() {
	// Инициализируем логгер
	log, err := logger.New("./logs/app.log")
	if err != nil {
		fmt.Printf("Failed to initialize logger: %v\n", err)
		os.Exit(1)
	}
	defer log.Close()

	log.Info("Starting SMK-SellerService...")

	// Подключаемся к базе данных
	dsn := "host=localhost port=5436 user=postgres password=postgres dbname=smk_sellerservice sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Проверяем соединение
	if err := db.Ping(); err != nil {
		log.Fatal("Failed to ping database: %v", err)
	}
	log.Info("Successfully connected to database")

	// Инициализируем репозитории
	companyRepository := companyRepo.NewRepository(db)
	serviceRepository := serviceRepo.NewRepository(db)

	// Инициализируем сервисы
	companySvc := companiesService.NewService(companyRepository)
	serviceSvc := servicesService.NewService(serviceRepository, companyRepository)

	// Инициализируем handlers для компаний
	createCompanyHandler := create_company.NewHandler(companySvc, log)
	getCompanyHandler := get_company.NewHandler(companySvc, log)
	listCompaniesHandler := list_companies.NewHandler(companySvc, log)
	updateCompanyHandler := update_company.NewHandler(companySvc, log)
	deleteCompanyHandler := delete_company.NewHandler(companySvc, log)

	// Инициализируем handlers для услуг
	createServiceHandler := create_service.NewHandler(serviceSvc, log)
	getServiceHandler := get_service.NewHandler(serviceSvc, log)
	listServicesHandler := list_services.NewHandler(serviceSvc, log)
	updateServiceHandler := update_service.NewHandler(serviceSvc, log)
	deleteServiceHandler := delete_service.NewHandler(serviceSvc, log)

	// Настраиваем роутер
	r := mux.NewRouter()

	// API prefix
	api := r.PathPrefix("/api/v1").Subrouter()

	// Public routes для компаний
	api.HandleFunc("/companies", listCompaniesHandler.Handle).Methods(http.MethodGet)
	api.HandleFunc("/companies/{id}", getCompanyHandler.Handle).Methods(http.MethodGet)

	// Public routes для услуг
	api.HandleFunc("/companies/{company_id}/services", listServicesHandler.Handle).Methods(http.MethodGet)
	api.HandleFunc("/companies/{company_id}/services/{service_id}", getServiceHandler.Handle).Methods(http.MethodGet)

	// Protected routes (требуют X-User-ID и X-User-Role)
	protected := api.PathPrefix("").Subrouter()
	protected.Use(middleware.Auth)

	// Protected routes для компаний
	protected.HandleFunc("/companies", createCompanyHandler.Handle).Methods(http.MethodPost)
	protected.HandleFunc("/companies/{id}", updateCompanyHandler.Handle).Methods(http.MethodPut)
	protected.HandleFunc("/companies/{id}", deleteCompanyHandler.Handle).Methods(http.MethodDelete)

	// Protected routes для услуг
	protected.HandleFunc("/companies/{company_id}/services", createServiceHandler.Handle).Methods(http.MethodPost)
	protected.HandleFunc("/companies/{company_id}/services/{service_id}", updateServiceHandler.Handle).Methods(http.MethodPut)
	protected.HandleFunc("/companies/{company_id}/services/{service_id}", deleteServiceHandler.Handle).Methods(http.MethodDelete)

	// Создаем HTTP сервер
	addr := ":8081"
	srv := &http.Server{
		Addr:         addr,
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Graceful shutdown
	go func() {
		log.Info("Starting server on %s", addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Server failed to start: %v", err)
		}
	}()

	// Ожидаем сигнал завершения
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Error("Server forced to shutdown: %v", err)
	}

	log.Info("Server stopped gracefully")
}
