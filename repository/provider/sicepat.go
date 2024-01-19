package provider

import (
	"mock-shipping-provider/primitive"
	"mock-shipping-provider/repository"
)

type Sicepat struct {
	Price           int64
	HourPerDistance int64
	KmPerDistance   int64
}

func NewSicepatCalculation(sicepatProvider Sicepat) repository.ProviderCalculation {
	return &Sicepat{
		Price:           sicepatProvider.Price,
		HourPerDistance: sicepatProvider.HourPerDistance,
		KmPerDistance:   sicepatProvider.KmPerDistance,
	}
}

func (sicepat *Sicepat) CalculatePrice(distance float64, dimension primitive.Dimension, weight float64) int64 {
	volume := dimension.Width * dimension.Height * dimension.Depth

	var hops int64

	if distance < 1 {
		hops = 1
	} else {
		hops = int64(distance) / sicepat.KmPerDistance
	}

	return sicepat.Price * hops * int64(volume)
}

func (sicepat *Sicepat) CalculateTimeOfArrival(distance float64) int64 {
	return int64(distance) / sicepat.HourPerDistance
}
