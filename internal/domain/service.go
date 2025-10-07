package domain

import "time"

// Service представляет услугу
type Service struct {
	ID              int64
	CompanyID       int64
	Name            string
	Description     *string
	AverageDuration *int
	AddressIDs      []int64
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

// ServicePublic представляет публичную информацию об услуге
type ServicePublic struct {
	ID              int64
	CompanyID       int64
	Name            string
	Description     *string
	AverageDuration *int
	Addresses       []Address
}

// CreateServiceInput входные данные для создания услуги
type CreateServiceInput struct {
	Name            string
	Description     *string
	AverageDuration *int
	AddressIDs      []int64
}

// UpdateServiceInput входные данные для обновления услуги
type UpdateServiceInput struct {
	Name            *string
	Description     *string
	AverageDuration *int
	AddressIDs      []int64
}
