package repository

import (
	"context"
	"time"

	"mock-shipping-provider/primitive"
)

type OrderLogRepository interface {
	// Create will insert multiple row into the database.
	// Because we are a mock server, nothing should be done through any
	// administrator dashboard or some sort that enable real people to
	// change the shipping status. Future events can be inserted here.
	Create(ctx context.Context, logEntry LogEntry) error
	// Get will return only the order history that are before or equal to
	// current time. If there are any order history which event happen
	// after current time, it will not be shown.
	Get(ctx context.Context, referenceNumber string, airWaybill string) ([]OrderHistory, error)
}

type OrderHistory struct {
	StatusCode primitive.Status
	Timestamp  time.Time
	Note       string
}

type LogEntry struct {
	ReferenceNumber string
	AirWaybill      string
	History         []OrderHistory
}
