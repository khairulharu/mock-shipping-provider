package distance_test

import (
	"mock-shipping-provider/primitive"
	"mock-shipping-provider/repository/distance"
	"testing"
)

func TestDistanceCalculation(t *testing.T) {
	distanceCalculation := distance.NewCalculateDistance()

	t.Run("test distance not service", func(t *testing.T) {
		from := primitive.Coordinate{
			Latitude:  -4.5387718,
			Longitude: 120.3146973,
		}

		to := primitive.Coordinate{
			Latitude:  40.7128,
			Longitude: -74.0060,
		}
		distance, serviceable := distanceCalculation.Calculate(from, to)
		if serviceable {
			t.Errorf("must not servicable: %v", distance)
		}
	})

	t.Run("test distance serviceable", func(t *testing.T) {
		from := primitive.Coordinate{
			Latitude:  -6.1601629,
			Longitude: 106.6793193,
		}

		to := primitive.Coordinate{
			Latitude:  -4.5387718,
			Longitude: 120.3146973,
		}
		distance, serviceable := distanceCalculation.Calculate(from, to)
		if !serviceable {
			t.Errorf("must be servicable: %v", distance)
		}
	})

}
