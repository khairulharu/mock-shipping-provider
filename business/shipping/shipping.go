package shipping

import (
	"context"

	"mock-shipping-provider/business"
	"mock-shipping-provider/repository"
)

type Dependency struct {
	orderLogRepository  repository.OrderLogRepository
	webhookClient       repository.WebhookClient
	distanceCalculation repository.DistanceCalculation
	priceCalculation    repository.PriceCalculation
}

// TODO: move these into different file and implement them

func (d *Dependency) Estimate(ctx context.Context, request business.EstimateRequest) (business.EstimateResult, error) {
	// TODO implement me
	panic("implement me")
}

func (d *Dependency) Create(ctx context.Context, request business.CreateRequest) (business.CreateResponse, error) {
	// TODO implement me
	panic("implement me")
}

func (d *Dependency) StatusHistory(ctx context.Context, request business.StatusRequest) ([]business.StatusHistoryResponse, error) {
	// TODO implement me
	panic("implement me")
}

func NewShippingService(orderLogRepository repository.OrderLogRepository, webhookClient repository.WebhookClient) (*Dependency, error) {
	// TODO: make sure orderLogRepository and webhookClient is not nil. If they are, return an error
	return &Dependency{
		orderLogRepository:  orderLogRepository,
		webhookClient:       webhookClient,
		distanceCalculation: nil, // implement later
		priceCalculation:    nil,
	}, nil
}
