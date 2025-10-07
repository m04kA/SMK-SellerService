package create_company

import (
	"context"

	"github.com/m04kA/SMK-SellerService/internal/service/companies/models"
)

type CompanyService interface {
	Create(ctx context.Context, userID int64, userRole string, req *models.CreateCompanyRequest) (*models.CompanyResponse, error)
}

type Logger interface {
	Info(format string, v ...interface{})
	Warn(format string, v ...interface{})
	Error(format string, v ...interface{})
}
