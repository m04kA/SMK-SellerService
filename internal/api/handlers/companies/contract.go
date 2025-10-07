package companies

import (
	"context"

	"github.com/m04kA/SMK-SellerService/internal/service/companies/models"
)

type CompanyService interface {
	Create(ctx context.Context, userID int64, userRole string, req *models.CreateCompanyRequest) (*models.CompanyResponse, error)
	GetByID(ctx context.Context, id int64) (*models.CompanyResponse, error)
	List(ctx context.Context, req *models.CompanyFilterRequest) (*models.CompanyListResponse, error)
	Update(ctx context.Context, id int64, userID int64, userRole string, req *models.UpdateCompanyRequest) (*models.CompanyResponse, error)
	Delete(ctx context.Context, id int64, userID int64, userRole string) error
}
