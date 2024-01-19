package provider_test

import (
	"mock-shipping-provider/primitive"
	"mock-shipping-provider/repository/provider"
	"testing"
)

func createSicepatCalculation(t *testing.T, distance float64, dimension primitive.Dimension, weight float64) int64 {
	Sicepat := provider.NewSicepatCalculation(provider.Sicepat{
		Price:           20000,
		HourPerDistance: 200,
		KmPerDistance:   300,
	})

	result := Sicepat.CalculatePrice(distance, dimension, weight)

	return result
}

func createSicepatTimeArrival(t *testing.T, distance float64) int64 {
	Sicepat := provider.NewSicepatCalculation(provider.Sicepat{
		Price:           20000,
		HourPerDistance: 200,
		KmPerDistance:   300,
	})

	result := Sicepat.CalculateTimeOfArrival(distance)

	return result
}

func TestProviderSicepat(t *testing.T) {
	distanceTest := 4500.0
	dimensionTest := primitive.Dimension{
		Width:  30,
		Height: 50,
		Depth:  40,
	}

	t.Run("test calculate price", func(t *testing.T) {
		Sicepat := provider.NewSicepatCalculation(provider.Sicepat{
			Price:           20000,
			HourPerDistance: 200,
			KmPerDistance:   300,
		})

		result := Sicepat.CalculatePrice(20, dimensionTest, 0)

		expectedResult := createSicepatCalculation(t, distanceTest, dimensionTest, 0)

		if result != expectedResult {
			t.Errorf("get error must price is: %v and %v", result, expectedResult)
		}
	})

	t.Run("test time of arrival", func(t *testing.T) {
		Sicepat := provider.NewSicepatCalculation(provider.Sicepat{
			Price:           20000,
			HourPerDistance: 200,
			KmPerDistance:   300,
		})

		result := Sicepat.CalculateTimeOfArrival(distanceTest)

		expectedResultTime := createSicepatTimeArrival(t, distanceTest)

		if result != expectedResultTime {
			t.Errorf("must be not error but get result time: %v, end expected: %v", result, expectedResultTime)
		}
	})

	t.Run("test failed or direction", func(t *testing.T) {
		Sicepat := provider.NewSicepatCalculation(provider.Sicepat{
			Price:           20000,
			HourPerDistance: 200,
			KmPerDistance:   300,
		})

		result := Sicepat.CalculateTimeOfArrival(distanceTest)

		expectedResultTime := createSicepatTimeArrival(t, 3)

		if result != expectedResultTime {
			t.Errorf("must be not error but get result time: %v, end expected: %v", result, expectedResultTime)
		}
	})
}
