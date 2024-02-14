package orderlog

import (
	"context"
	"fmt"
	"mock-shipping-provider/repository"
)

func (o *orderLogRepository) Get(ctx context.Context, referenceNumber string, airWaybill string) ([]repository.OrderHistory, error) {

	coon, err := o.db.Conn(ctx)

	if err != nil {
		return []repository.OrderHistory{}, fmt.Errorf("error when open new connection get: %v", err)
	}

	defer func() {
		if err := coon.Close(); err != nil {

		}
	}()
}
