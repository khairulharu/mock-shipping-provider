package repository

import (
	"context"
	"time"

	"mock-shipping-provider/primitive"
)

type WebhookClient interface {
	// SendStatusUpdate should send a POST request to the destination target URL (configured somewhere else)
	// with the JSON request body of:
	//
	// {
	//   "reference_number": "string",
	//   "air_waybill": "string",
	//   "status_code": 3,
	//   "status_description": "IN_TRANSIT",
	//   "timestamp": "time.RFC3339",
	//   "note": "string"
	// }
	//
	// It should do a retry with exponential backoff if the client returns 500 status code,
	// or a timeout error happened.
	//
	// If the response status code is 400, it should throw an error and don't continue.
	SendStatusUpdate(context.Context, StatusUpdate) error
}

type StatusUpdate struct {
	ReferenceNumber string
	AirWaybill      string
	Status          primitive.Status
	Timestamp       time.Time
	Note            string
}
