package provider

import (
	"mock-shipping-provider/primitive"
	"mock-shipping-provider/repository"
)

type JNE struct {
	RatePerKilogram  int64
	RatePerKilometer int64
	RatePerCmCubic   int64
	KilometerPerHour int64
}

func NewJneCalculation() repository.ProviderCalculation {
	return &JNE{
		RatePerKilogram:  3120,
		RatePerKilometer: 4150,
		RatePerCmCubic:   1245,
		KilometerPerHour: 60,
	}
}

func (jne *JNE) CalculatePrice(distance float64, dimension primitive.Dimension, weight float64) int64 {

	volume := dimension.Width * dimension.Height * dimension.Depth

	distanceCost := distance * float64(jne.RatePerKilometer)

	weightCost := weight * float64(jne.RatePerKilogram)

	volumeCost := volume * float64(jne.RatePerCmCubic)

	return int64(distanceCost + weightCost + volumeCost)
}

func (jne *JNE) CalculateTimeOfArrival(distance float64) int64 {
	if int64(distance) < jne.KilometerPerHour {
		return 1
	}

	return int64(distance) / jne.KilometerPerHour
}
