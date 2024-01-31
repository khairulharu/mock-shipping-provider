package orderlog

import (
	"context"
	"database/sql"
	"mock-shipping-provider/repository"
)

type orderLogRepository struct {
	db *sql.DB
}

func NewOrderLog(db *sql.DB) repository.OrderLogRepository {
	return &orderLogRepository{
		db: db,
	}
}

func (o *orderLogRepository) Create(ctx context.Context, logEntry repository.LogEntry) error {
	panic("unimplem")
}

func (o *orderLogRepository) Get(ctx context.Context, referenceNumber string, airWaybill string) ([]repository.OrderHistory, error) {
	panic("unimplemented")
}
