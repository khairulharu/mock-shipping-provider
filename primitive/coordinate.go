package primitive

// Coordinate represents the coordinate of the Earth
type Coordinate struct {
	// Longitude is the measurement east or west of the prime meridian.
	// Longitude is measured by imaginary lines that run around Earth vertically (up and down)
	// and meet at the North and South Poles. These lines are known as meridians.
	// Each meridian measures one arc degree of longitude.
	// The distance around Earth measures 360 degrees.
	Longitude float64 `json:"longitude"`

	// Latitude is the measurement of distance north or south of the Equator.
	// It is measured with 180 imaginary lines that form circles around Earth east-west,
	// parallel to the Equator. These lines are known as parallels. A circle of latitude
	// is an imaginary ring linking all points sharing a parallel.
	Latitude float64 `json:"latitude"`
}
