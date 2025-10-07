package service

import "errors"

var (
	// ErrServiceNotFound возвращается, когда услуга не найдена в БД
	ErrServiceNotFound = errors.New("repository: service not found")

	// ErrCompanyNotFound возвращается, когда компания не найдена в БД
	ErrCompanyNotFound = errors.New("repository: company not found")

	// ErrBuildQuery возвращается при ошибке построения SQL запроса
	ErrBuildQuery = errors.New("repository: failed to build SQL query")

	// ErrExecQuery возвращается при ошибке выполнения SQL запроса
	ErrExecQuery = errors.New("repository: failed to execute SQL query")

	// ErrScanRow возвращается при ошибке сканирования строки из БД
	ErrScanRow = errors.New("repository: failed to scan row")

	// ErrTransaction возвращается при ошибке работы с транзакцией
	ErrTransaction = errors.New("repository: transaction error")
)
