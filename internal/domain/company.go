package domain

import "time"

// Company представляет компанию
type Company struct {
	ID           int64
	Name         string
	Logo         *string
	Description  *string
	Tags         []string
	Addresses    []Address
	WorkingHours WorkingHours
	ManagerIDs   []int64
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// CompanyPublic представляет публичную информацию о компании
type CompanyPublic struct {
	ID           int64
	Name         string
	Logo         *string
	Description  *string
	Tags         []string
	WorkingHours WorkingHours
}

// CreateCompanyInput входные данные для создания компании
type CreateCompanyInput struct {
	Name         string
	Logo         *string
	Description  *string
	Tags         []string
	Addresses    []AddressInput
	WorkingHours WorkingHours
	ManagerIDs   []int64
}

// UpdateCompanyInput входные данные для обновления компании
type UpdateCompanyInput struct {
	Name         *string
	Logo         *string
	Description  *string
	Tags         []string
	Addresses    []AddressUpdateInput
	WorkingHours *WorkingHours
	ManagerIDs   []int64
}

// CompanyFilter фильтры для поиска компаний
type CompanyFilter struct {
	Tags  []string
	City  *string
	Page  *int // Опционально: если nil, пагинация не применяется
	Limit *int // Опционально: если nil, пагинация не применяется
}
