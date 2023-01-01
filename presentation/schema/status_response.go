package schema

import (
	"mock-shipping-provider/primitive"
)

type StatusHistory struct {
	StatusCode        primitive.Status `json:"status_code"`
	StatusDescription string           `json:"status_description"`
	Timestamp         string           `json:"timestamp"`
	Note              string           `json:"note"`
}

type StatusResponse struct {
	ReferenceNumber string          `json:"reference_number"`
	AirWaybill      string          `json:"air_waybill"`
	History         []StatusHistory `json:"history"`
}
