package service

import (
	"context"
	"database/sql"
)

// DBExecutor интерфейс для выполнения SQL запросов (поддерживает *sql.DB и *sql.Tx)
type DBExecutor interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
}
