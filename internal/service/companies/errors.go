package companies

import "errors"

var (
	// ErrCompanyNotFound возвращается, когда компания не найдена
	ErrCompanyNotFound = errors.New("company not found")

	// ErrAccessDenied возвращается, когда у пользователя нет прав доступа к компании
	ErrAccessDenied = errors.New("access denied: user is not a manager of this company")

	// ErrOnlySuperuser возвращается, когда операцию может выполнить только superuser
	ErrOnlySuperuser = errors.New("access denied: only superuser can create companies")

	// ErrInvalidInput возвращается при некорректных входных данных
	ErrInvalidInput = errors.New("invalid input data")

	// ErrInternal возвращается при внутренних ошибках сервиса
	ErrInternal = errors.New("service: internal error")
)
