package primitive

import "errors"

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
		return errors.New("height is lower than 0")
	}

	if d.Width < 0 {
		return errors.New("width is lower than 0")
	}

	if d.Depth < 0 {
		return errors.New("depth is lower than 0")
	}

	return nil
}
