package provider_test

import (
	"mock-shipping-provider/primitive"
	"mock-shipping-provider/repository/provider"
	"reflect"
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
	testCases := []struct {
		distance  float64
		dimension primitive.Dimension
		weight    float64
	}{
		{distance: 4500.0},
		{dimension: primitive.Dimension{
			Width:  30,
			Height: 20,
			Depth:  30,
		}},
		{weight: 45},
	}

	for _, testcase := range testCases {
		t.Run("test calculate price", func(t *testing.T) {
			JNE := provider.NewJneCalculation()

			result := JNE.CalculatePrice(float64(testcase.distance), testcase.dimension, testcase.weight)

			expectedPrice := createJneCalculation(t, testcase.distance, testcase.dimension, testcase.weight)

			if reflect.DeepEqual(result, expectedPrice) == false {
				t.Errorf("must not error but get %v and %v different price meaning code error", result, expectedPrice)
			}
		})
	}
	for _, testCase := range testCases {
		t.Run("test time of arrival", func(t *testing.T) {
			JNE := provider.NewJneCalculation()

			time := JNE.CalculateTimeOfArrival(testCase.distance)

			expectedTime := createJneTimeArrival(t, testCase.distance)

			if reflect.DeepEqual(time, expectedTime) == false {
				t.Errorf("error get time arrival is : %v hours ", time)
			}
		})
	}
}
