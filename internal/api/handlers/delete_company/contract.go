package delete_company

import "context"

type CompanyService interface {
	Delete(ctx context.Context, id int64, userID int64, userRole string) error
}

type Logger interface {
	Info(format string, v ...interface{})
	Warn(format string, v ...interface{})
	Error(format string, v ...interface{})
}
