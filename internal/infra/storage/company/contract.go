package company

import (
	"context"
	"database/sql"

	"github.com/m04kA/SMK-SellerService/pkg/dbmetrics"
)

// Переиспользуем интерфейсы из dbmetrics
type DBExecutor = dbmetrics.DBExecutor
type TxExecutor = dbmetrics.TxExecutor

// TxBeginner интерфейс для начала транзакций (поддерживает *sql.DB и *dbmetrics.DB)
type TxBeginner interface {
	BeginTx(ctx context.Context, opts *sql.TxOptions) (TxExecutor, error)
}
