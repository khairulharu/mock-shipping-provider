package orderlog

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"mock-shipping-provider/repository"

	"github.com/rs/zerolog"
)

func (o *orderLogRepository) Create(ctx context.Context, logEntry repository.LogEntry) error {

	conn, err := o.db.Conn(ctx)
	if err != nil {
		return fmt.Errorf("acquiring connection from pool: %w", err)
	}

	defer func() {
		err := conn.Close()
		if err != nil && !errors.Is(err, sql.ErrConnDone) {
			log := zerolog.Ctx(ctx)
			log.Err(err).Msg("returning connection back to pool")
		}
	}()
	return nil
}
