package services

import "errors"

var (
	// ErrServiceNotFound возвращается, когда услуга не найдена
	ErrServiceNotFound = errors.New("service not found")

	// ErrCompanyNotFound возвращается, когда компания не найдена
	ErrCompanyNotFound = errors.New("company not found")

	// ErrAccessDenied возвращается, когда у пользователя нет прав доступа к услуге/компании
	ErrAccessDenied = errors.New("access denied: user is not a manager of this company")

	// ErrInvalidInput возвращается при некорректных входных данных
	ErrInvalidInput = errors.New("invalid input data")

	// ErrInternal возвращается при внутренних ошибках сервиса
	ErrInternal = errors.New("service: internal error")
)
