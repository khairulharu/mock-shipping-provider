package provider_test

import (
	"mock-shipping-provider/primitive"
	"mock-shipping-provider/repository/provider"
	"testing"
)

func createCalculation(t *testing.T, disatance float64, dimension primitive.Dimension, weight float64) int64 {
	JNE := provider.NewJneCalculation()

	result := JNE.CalculatePrice(disatance, primitive.Dimension{
		Width:  dimension.Width,
		Height: dimension.Height,
		Depth:  dimension.Depth,
	}, weight)

	return result
}

func createTimeArrival(t *testing.T, distance float64) int64 {
	JNE := provider.NewJneCalculation()

	result := JNE.CalculateTimeOfArrival(distance)

	return result
}

func TestCalculate(t *testing.T) {
	var distanceTest float64 = 300

	dimensionTest := primitive.Dimension{
		Width:  30,
		Height: 60,
		Depth:  20,
	}

	t.Run("test calculate", func(t *testing.T) {
		JNE := provider.NewJneCalculation()

		result := JNE.CalculatePrice(float64(distanceTest), primitive.Dimension{
			Width:  dimensionTest.Width,
			Height: dimensionTest.Height,
			Depth:  dimensionTest.Depth,
		}, 0)

		resultOfCalculation := createCalculation(t, distanceTest, dimensionTest, 0)

		if result != resultOfCalculation {
			t.Errorf("must be not error get: %v Rp", result)
		}
	})

	t.Run("test for time of arrival", func(t *testing.T) {
		JNE := provider.NewJneCalculation()

		time := JNE.CalculateTimeOfArrival(distanceTest)

		resultOfTime := createTimeArrival(t, distanceTest)

		if time != resultOfTime {
			t.Errorf("must error because time arrival is : %v hours ", time)
		}
	})
}
