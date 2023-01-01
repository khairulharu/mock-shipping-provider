package repository

import "mock-shipping-provider/primitive"

type PriceCalculation interface {
	// Calculate price from distance, dimention, and weight combination.
	Calculate(distance float64, dimension primitive.Dimension, weight float64) int64
}
