package primitive

import "errors"

var HeightIsLowerThanZero = errors.New("height is lower than 0")
var WidthIsLowerThanZero = errors.New("width is lower than 0")
var DepthIsLowerThanZero = errors.New("depth is lower than 0")

// Dimension specifies the dimension of an object.
// Every field represent the length in centimeters.
type Dimension struct {
	Height float64 `json:"height"`
	Width  float64 `json:"width"`
	Depth  float64 `json:"depth"`
}

// Validate will return an error if any of the field on Dimension
// is lower than zero.
func (d Dimension) Validate() error {
	if d.Height < 0 {
		return HeightIsLowerThanZero
	}

	if d.Width < 0 {
		return WidthIsLowerThanZero
	}

	if d.Depth < 0 {
		return DepthIsLowerThanZero
	}

	return nil
}
