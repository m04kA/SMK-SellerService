package priceservice

// CalculatePricesRequest запрос на расчёт цен
type CalculatePricesRequest struct {
	CompanyID  int64   `json:"company_id"`
	UserID     *int64  `json:"user_id,omitempty"`
	ServiceIDs []int64 `json:"service_ids"`
}

// CalculatePricesResponse ответ с рассчитанными ценами
type CalculatePricesResponse struct {
	Prices []ServicePrice `json:"prices"`
}

// ServicePrice цена на услугу
type ServicePrice struct {
	ServiceID         int64    `json:"service_id"`
	Price             *float64 `json:"price,omitempty"`
	Currency          *string  `json:"currency,omitempty"`
	PricingType       *string  `json:"pricing_type,omitempty"`
	VehicleClass      *string  `json:"vehicle_class,omitempty"`
	AppliedMultiplier *float64 `json:"applied_multiplier,omitempty"`
}

// ErrorResponse модель ошибки от PriceService
type ErrorResponse struct {
	Error string `json:"error"`
}
