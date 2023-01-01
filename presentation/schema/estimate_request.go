package schema

import "mock-shipping-provider/primitive"

type EstimateRequest struct {
	Sender    primitive.Coordinate `json:"sender"`
	Recipient primitive.Coordinate `json:"recipient"`
	Dimension primitive.Dimension  `json:"dimension"`
	Weight    float64              `json:"weight"`
}
