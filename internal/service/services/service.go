package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/m04kA/SMK-SellerService/internal/service"
	"github.com/m04kA/SMK-SellerService/internal/service/services/models"
	companyRepo "github.com/m04kA/SMK-SellerService/internal/infra/storage/company"
	serviceRepo "github.com/m04kA/SMK-SellerService/internal/infra/storage/service"
)

type Service struct {
	serviceRepo ServiceRepository
	companyRepo CompanyRepository
}

func NewService(serviceRepo ServiceRepository, companyRepo CompanyRepository) *Service {
	return &Service{
		serviceRepo: serviceRepo,
		companyRepo: companyRepo,
	}
}

// Create создает новую услугу для компании
func (s *Service) Create(ctx context.Context, companyID int64, userID int64, userRole string, req *models.CreateServiceRequest) (*models.ServiceResponse, error) {
	// Проверка прав доступа к компании
	if err := s.checkAccess(ctx, companyID, userID, userRole); err != nil {
		return nil, err
	}

	input := req.ToDomainCreateInput()
	service, err := s.serviceRepo.Create(ctx, companyID, input)
	if err != nil {
		return nil, fmt.Errorf("%w: Create - repository error: %v", ErrInternal, err)
	}

	return models.FromDomainService(service), nil
}

// GetByID получает услугу по ID
func (s *Service) GetByID(ctx context.Context, companyID int64, serviceID int64) (*models.ServiceResponse, error) {
	service, err := s.serviceRepo.GetByID(ctx, companyID, serviceID)
	if err != nil {
		if errors.Is(err, serviceRepo.ErrServiceNotFound) {
			return nil, ErrServiceNotFound
		}
		return nil, fmt.Errorf("%w: GetByID - repository error: %v", ErrInternal, err)
	}

	return models.FromDomainService(service), nil
}

// ListByCompany получает список услуг компании
func (s *Service) ListByCompany(ctx context.Context, companyID int64) (*models.ServiceListResponse, error) {
	services, err := s.serviceRepo.ListByCompany(ctx, companyID)
	if err != nil {
		return nil, fmt.Errorf("%w: ListByCompany - repository error: %v", ErrInternal, err)
	}

	return models.FromDomainServiceList(services), nil
}

// Update обновляет услугу
func (s *Service) Update(ctx context.Context, companyID int64, serviceID int64, userID int64, userRole string, req *models.UpdateServiceRequest) (*models.ServiceResponse, error) {
	// Проверка прав доступа к компании
	if err := s.checkAccess(ctx, companyID, userID, userRole); err != nil {
		return nil, err
	}

	input := req.ToDomainUpdateInput()
	service, err := s.serviceRepo.Update(ctx, companyID, serviceID, input)
	if err != nil {
		if errors.Is(err, serviceRepo.ErrServiceNotFound) {
			return nil, ErrServiceNotFound
		}
		return nil, fmt.Errorf("%w: Update - repository error: %v", ErrInternal, err)
	}

	return models.FromDomainService(service), nil
}

// Delete удаляет услугу
func (s *Service) Delete(ctx context.Context, companyID int64, serviceID int64, userID int64, userRole string) error {
	// Проверка прав доступа к компании
	if err := s.checkAccess(ctx, companyID, userID, userRole); err != nil {
		return err
	}

	if err := s.serviceRepo.Delete(ctx, companyID, serviceID); err != nil {
		if errors.Is(err, serviceRepo.ErrServiceNotFound) {
			return ErrServiceNotFound
		}
		return fmt.Errorf("%w: Delete - repository error: %v", ErrInternal, err)
	}

	return nil
}

// checkAccess проверяет права доступа пользователя к компании
func (s *Service) checkAccess(ctx context.Context, companyID int64, userID int64, userRole string) error {
	// Superuser имеет полный доступ
	if userRole == service.RoleSuperuser {
		return nil
	}

	// Обычный пользователь должен быть менеджером компании
	isManager, err := s.companyRepo.IsManager(ctx, companyID, userID)
	if err != nil {
		if errors.Is(err, companyRepo.ErrCompanyNotFound) {
			return ErrCompanyNotFound
		}
		return fmt.Errorf("%w: checkAccess - repository error: %v", ErrInternal, err)
	}

	if !isManager {
		return ErrAccessDenied
	}

	return nil
}
