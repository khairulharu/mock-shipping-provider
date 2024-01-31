package provider

import (
	"mock-shipping-provider/primitive"
	"mock-shipping-provider/repository"
)

type ProviderAll struct {
	JNE      repository.ProviderCalculation
	JNT      repository.ProviderCalculation
	SiCepat  repository.ProviderCalculation
	AnterAja repository.ProviderCalculation
}

func GetProviderCalculation() (*ProviderAll, error) {
	jneRate := primitive.Rate{
		PerKilogram:      700,
		PerKilometer:     600,
		PerCmCubic:       500,
		KilometerPerHour: 60,
	}
	if err := jneRate.Validate(); err != nil {
		return &ProviderAll{}, err
	}
	jneCalculation := NewJneCalculation(&jneRate)

	jntRate := primitive.Rate{
		PerKilogram:      700,
		PerKilometer:     600,
		PerCmCubic:       500,
		KilometerPerHour: 60,
	}
	if err := jntRate.Validate(); err != nil {
		return &ProviderAll{}, err
	}
	jntCalculation := NewJneCalculation(&jntRate)

	anterajaRate := primitive.Rate{
		PerKilogram:      700,
		PerKilometer:     600,
		PerCmCubic:       500,
		KilometerPerHour: 60,
	}
	if err := anterajaRate.Validate(); err != nil {
		return &ProviderAll{}, err
	}
	anterajaCalculation := NewJneCalculation(&anterajaRate)

	sicepatRate := primitive.Rate{
		PerKilogram:      700,
		PerKilometer:     600,
		PerCmCubic:       500,
		KilometerPerHour: 60,
	}
	if err := sicepatRate.Validate(); err != nil {
		return &ProviderAll{}, err
	}
	sicepatCalculation := NewJneCalculation(&sicepatRate)

	return &ProviderAll{
		JNE:      jneCalculation,
		JNT:      jntCalculation,
		SiCepat:  sicepatCalculation,
		AnterAja: anterajaCalculation,
	}, nil
}
