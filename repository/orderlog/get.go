package orderlog

import (
	"context"
	"mock-shipping-provider/repository"
)

func (o *orderLogRepository) Get(ctx context.Context, referenceNumber string, airWaybill string) ([]repository.OrderHistory, error) {
	panic("not implement")
}
