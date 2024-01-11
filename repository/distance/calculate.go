package distance

import (
	"math"
	"mock-shipping-provider/primitive"
	"mock-shipping-provider/repository"
)

type CalculateDistance struct {
	Radius             float64
	ServicableDistance float64
}

func NewCalculateDistance() repository.DistanceCalculation {
	return &CalculateDistance{
		Radius:             6371.0,
		ServicableDistance: 5.100,
	}
}

func (c *CalculateDistance) Calculate(from primitive.Coordinate, to primitive.Coordinate) (distance float64, serviceable bool) {
	// R := 6371.0
	// serviceableDistance := 5.100

	lat1Rad := from.Latitude * math.Pi / 180
	lon1Rad := from.Longitude * math.Pi / 180
	lat2Rad := to.Latitude * math.Pi / 180
	lon2Rad := to.Longitude * math.Pi / 180

	dlat := lat2Rad - lat1Rad
	dlon := lon2Rad - lon1Rad

	a := math.Sin(dlat/2)*math.Sin(dlat/2) + math.Cos(lat1Rad)*math.Cos(lat2Rad)*math.Sin(dlon/2)*math.Sin(dlon/2)
	cal := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	distanceCalculation := c.Radius * cal

	if distanceCalculation > c.ServicableDistance {
		return distance, false
	}

	return distance, true
}
