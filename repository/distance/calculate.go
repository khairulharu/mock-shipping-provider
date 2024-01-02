package distance

import (
	"math"
	"mock-shipping-provider/primitive"
)

type Calculate struct {
}

func (cal *Calculate) Calculate(from primitive.Coordinate, to primitive.Coordinate) (distance float64, serviceable bool) {
	R := 6371.0
	serviceableDistance := 5.100

	lat1Rad := from.Latitude * math.Pi / 180
	lon1Rad := from.Longitude * math.Pi / 180
	lat2Rad := to.Latitude * math.Pi / 180
	lon2Rad := to.Longitude * math.Pi / 180

	dlat := lat2Rad - lat1Rad
	dlon := lon2Rad - lon1Rad

	a := math.Sin(dlat/2)*math.Sin(dlat/2) + math.Cos(lat1Rad)*math.Cos(lat2Rad)*math.Sin(dlon/2)*math.Sin(dlon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	distanceCalculation := R * c

	if distanceCalculation > serviceableDistance {
		return distance, false
	}

	return distance, true
}
