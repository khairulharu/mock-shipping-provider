package provider_test

import (
	"mock-shipping-provider/primitive"
	"mock-shipping-provider/repository/provider"
	"testing"
)

func createJneCalculation(t *testing.T, distance float64, dimension primitive.Dimension, weight float64) int64 {
	JNE := provider.NewJneCalculation()

	result := JNE.CalculatePrice(distance, dimension, weight)

	return result
}

func createJneTimeArrival(t *testing.T, distance float64) int64 {
	JNE := provider.NewJneCalculation()

	result := JNE.CalculateTimeOfArrival(distance)

	return result
}

func TestProviderJne(t *testing.T) {
	distanceTest := 300.0

	dimensionTest := primitive.Dimension{
		Width:  30,
		Height: 60,
		Depth:  20,
	}

	t.Run("test calculate price", func(t *testing.T) {
		JNE := provider.NewJneCalculation()

		result := JNE.CalculatePrice(float64(distanceTest), dimensionTest, 0)

		resultOfCalculation := createJneCalculation(t, distanceTest, dimensionTest, 0)

		if result != resultOfCalculation {
			t.Errorf("must be no error, but get: %v Rp", result)
		}
	})

	t.Run("test time of arrival", func(t *testing.T) {
		JNE := provider.NewJneCalculation()

		time := JNE.CalculateTimeOfArrival(distanceTest)

		resultOfTime := createJneTimeArrival(t, distanceTest)

		if time != resultOfTime {
			t.Errorf("error get time arrival is : %v hours ", time)
		}
	})
}
