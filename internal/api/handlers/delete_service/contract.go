package delete_service

import "context"

type ServiceService interface {
	Delete(ctx context.Context, companyID int64, serviceID int64, userID int64, userRole string) error
}

type Logger interface {
	Info(format string, v ...interface{})
	Warn(format string, v ...interface{})
	Error(format string, v ...interface{})
}
