package shipping

import (
	"errors"
	"mock-shipping-provider/repository"
	"mock-shipping-provider/repository/provider"
)

type Dependency struct {
	orderLogRepository  repository.OrderLogRepository
	webhookClient       repository.WebhookClient
	distanceCalculation repository.DistanceCalculation
	provider            provider.Calculation
}

func NewShippingService(orderLogRepository repository.OrderLogRepository, webhookClient repository.WebhookClient, distanceCalculation repository.DistanceCalculation,
	provider provider.Calculation) (*Dependency, error) {
	// TODO: make sure orderLogRepository and webhookClient is not nil. If they are, return an error
	if orderLogRepository == nil && webhookClient == nil {
		return &Dependency{}, errors.New("nil orderlog repository and webhook client in business layer")
	}
	return &Dependency{
		orderLogRepository:  orderLogRepository,
		webhookClient:       webhookClient,
		distanceCalculation: distanceCalculation,
		provider:            provider,
	}, nil
}
