package repository

import "mock-shipping-provider/primitive"

type DistanceCalculation interface {
	// Calculate distance between two coordinate. It returns distance in kilometers.
	Calculate(from primitive.Coordinate, to primitive.Coordinate) (distance float64, serviceable bool)
}
