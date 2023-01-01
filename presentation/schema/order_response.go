package schema

import "mock-shipping-provider/primitive"

type OrderResponse struct {
	StatusCode           primitive.Status `json:"status_code"`
	StatusDescription    string           `json:"status_description"`
	ReferenceNumber      string           `json:"reference_number"`
	AirWaybill           string           `json:"air_waybill"`
	Price                int64            `json:"price"`
	EstimatedHourArrival int64            `json:"estimated_hour_arrival"`
}
