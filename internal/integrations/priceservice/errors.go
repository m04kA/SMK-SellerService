package priceservice

import "errors"

var (
	// ErrPricesNotFound возвращается, когда цены для услуг не найдены
	ErrPricesNotFound = errors.New("prices not found for services")

	// ErrInternal возвращается при внутренних ошибках клиента
	ErrInternal = errors.New("priceservice client: internal error")

	// ErrInvalidResponse возвращается при некорректном ответе от сервиса
	ErrInvalidResponse = errors.New("priceservice client: invalid response")

	// ErrServiceDegraded возвращается при применении graceful degradation
	// Указывает, что PriceService недоступен и следует вернуть данные без цен
	ErrServiceDegraded = errors.New("priceservice unavailable: graceful degradation applied")
)
