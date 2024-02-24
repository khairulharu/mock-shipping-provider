package orderlog

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"mock-shipping-provider/repository"

	"github.com/rs/zerolog"
)

func (o *orderLogRepository) Get(ctx context.Context, referenceNumber string, airWaybill string) ([]repository.OrderHistory, error) {
	conn, err := o.db.Conn(ctx)
	if err != nil {
		return []repository.OrderHistory{}, fmt.Errorf("acquiring connection from pool: %w", err)
	}

	defer func() {
		err := conn.Close()
		if err != nil && !errors.Is(err, sql.ErrConnDone) {
			log := zerolog.Ctx(ctx)
			log.Err(err).Msg("returning connection back to pool")
		}
	}()

	return []repository.OrderHistory{}, nil
}
