package provider

import (
	"mock-shipping-provider/primitive"
	"mock-shipping-provider/repository"
)

type JNE struct {
	Price           int64
	HourPerDistance int64
	KmPerDistance   int64
}

func NewJneCalculation() repository.ProviderCalculation {
	return &JNE{
		Price:           5000,
		HourPerDistance: 200,
		KmPerDistance:   300,
	}
}

func (jne *JNE) CalculatePrice(distance float64, dimension primitive.Dimension, weight float64) int64 {

	volume := dimension.Width * dimension.Height * dimension.Depth

	var hops int64

	if distance < float64(jne.KmPerDistance) {
		hops = 1
	} else {
		hops = (int64(distance) / jne.KmPerDistance)
	}

	return (jne.Price * hops) * int64(volume)
}

func (jne *JNE) CalculateTimeOfArrival(distance float64) int64 {
	if distance < float64(jne.HourPerDistance) {
		return 1
	}
	return int64(distance / float64(jne.HourPerDistance))
}
