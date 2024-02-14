package orderlog

import (
	"database/sql"
	"mock-shipping-provider/repository"
)

type orderLogRepository struct {
	db *sql.DB
}

func New(db *sql.DB) repository.OrderLogRepository {
	return &orderLogRepository{
		db: db,
	}
}
