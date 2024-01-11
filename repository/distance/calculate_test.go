package distance_test

import (
	"mock-shipping-provider/primitive"
	"mock-shipping-provider/repository/distance"
	"testing"
)

func TestDistanceCalculation(t *testing.T) {
	from := primitive.Coordinate{
		Latitude:  -4.5387718,
		Longitude: 120.3146973,
	}

	to := primitive.Coordinate{
		Latitude:  40.7128,
		Longitude: -74.0060,
	}

	fromA := primitive.Coordinate{
		Latitude:  6.2088,
		Longitude: 106.8456,
	}

	toA := primitive.Coordinate{
		Latitude:  -4.5387718,
		Longitude: 120.3146973,
	}

	distanceCalculation := distance.NewCalculateDistance()

	t.Run("test distance not service", func(t *testing.T) {
		distance, serviceable := distanceCalculation.Calculate(from, to)
		if !serviceable {
			t.Logf("not servicable: %v", distance)
		}
	})

	t.Run("test distance serviceable", func(t *testing.T) {
		distance, serviceable := distanceCalculation.Calculate(fromA, toA)
		if serviceable {
			t.Logf("servicable: %v", distance)
		}
	})

}
