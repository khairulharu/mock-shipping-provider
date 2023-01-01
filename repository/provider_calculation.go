package repository

import "mock-shipping-provider/primitive"

type ProviderCalculation interface {
	// CalculateTimeOfArrival calculates the time of arrival (in hours) based on distance.
	CalculateTimeOfArrival(distance float64) int64
	// CalculatePrice price from distance, dimension, and weight combination.
	CalculatePrice(distance float64, dimension primitive.Dimension, weight float64) int64
}
