package domain

import "time"

// Address представляет адрес компании
type Address struct {
	ID          int64
	CompanyID   int64
	City        string
	Street      string
	Building    string
	Coordinates Coordinates
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// Coordinates представляет географические координаты
type Coordinates struct {
	Latitude  float64
	Longitude float64
}

// AddressInput входные данные для создания адреса
type AddressInput struct {
	City        string
	Street      string
	Building    string
	Coordinates Coordinates
}

// AddressUpdateInput входные данные для обновления адреса
type AddressUpdateInput struct {
	ID          *int64 // Если указан - обновление существующего
	City        string
	Street      string
	Building    string
	Coordinates Coordinates
}
