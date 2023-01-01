package primitive_test

import (
	"testing"

	"mock-shipping-provider/primitive"
)

func TestDimension_Validate(t *testing.T) {
	t.Run("Happy", func(t *testing.T) {
		dimension := primitive.Dimension{
			Height: 0,
			Width:  0,
			Depth:  0,
		}

		err := dimension.Validate()
		if err != nil {
			t.Errorf("unexpected error: %s", err.Error())
		}
	})
}
