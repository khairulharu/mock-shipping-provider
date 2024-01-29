package shipping

import (
	"mock-shipping-provider/repository"
)

type Dependency struct {
	orderLogRepository  repository.OrderLogRepository
	webhookClient       repository.WebhookClient
	distanceCalculation repository.DistanceCalculation
}

func NewShippingService(orderLogRepository repository.OrderLogRepository, webhookClient repository.WebhookClient, distanceCalculation repository.DistanceCalculation) (*Dependency, error) {
	// TODO: make sure orderLogRepository and webhookClient is not nil. If they are, return an error
	return &Dependency{
		orderLogRepository:  orderLogRepository,
		webhookClient:       webhookClient,
		distanceCalculation: distanceCalculation,
	}, nil
}
