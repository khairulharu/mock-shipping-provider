package business

import (
	"context"
	"errors"
	"time"

	"mock-shipping-provider/primitive"
)

type Shipping interface {
	// Estimate calculates the two coordinate distance and returns a price and estimated arrival hours.
	// If it's too far, it returns ErrNotServiceable.
	Estimate(context.Context, EstimateRequest) ([]EstimateResult, error)
	// Create a new shipping order, recalculate the price and estimated arrival hours. If the location of
	// sender and recipient is too far, it returns ErrNotServiceable. If everything is fine,
	// it will return a reference number and an air waybill number.
	//
	// By calling this function, we activate a background worker that will do webhook calls on every status change
	// to the designated target URL.
	Create(context.Context, CreateRequest) (CreateResponse, error)
	// StatusHistory returns status history of the given reference number and air waybill.
	StatusHistory(context.Context, StatusRequest) (StatusHistoryResponse, error)
}

// ErrNotServiceable indicates the distance is too far or the service is not available on that specific coordinate.
// For example, it should not be able to deliver something to the middle of the pacific ocean.
var ErrNotServiceable = errors.New("not serviceable")

type EstimateRequest struct {
	Sender    primitive.Coordinate
	Recipient primitive.Coordinate
	Dimension primitive.Dimension
	Weight    float64
}

type EstimateResult struct {
	Provider primitive.Provider
	Price    int64
	Hours    uint64
}

// OrderRequest holds the order request
// data coming from the presentation layer
type CreateRequest struct {
	Provider        primitive.Provider
	Sender          primitive.Address
	Recipient       primitive.Address
	Dimension       primitive.Dimension
	Weight          float64
	ItemDescription string
	ItemCategory    string
	Fragile         bool
}

type CreateResponse struct {
	ReferenceNumber string
	// AirWaybill is a document that accompanies goods shipped by an international air courier
	// to provide detailed information about the shipment and allow it to be tracked.
	// The bill has multiple copies so that each party involved in the shipment can document it.
	// An air waybill (AWB), also known as an air consignment note, is a type of bill of lading.
	AirWaybill string
	Price      int64
	Hours      uint64
}

type StatusRequest struct {
	ReferenceNumber string
	AirWaybill      string
}

type StatusHistory struct {
	Status    primitive.Status
	Timestamp time.Time
	Note      string
}

type StatusHistoryResponse struct {
	ReferenceNumber string
	AirWaybill      string
	History         []StatusHistory
}
