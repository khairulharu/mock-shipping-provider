package shipping_test

import (
	"errors"
	"mock-shipping-provider/business"
	"mock-shipping-provider/business/shipping"
	"mock-shipping-provider/primitive"
	"testing"
)

func TestValidateEstimateRequest(t *testing.T) {
	// create example of valid request
	request := business.EstimateRequest{
		Sender: primitive.Coordinate{
			Latitude:  -7.7584860436179435,
			Longitude: 110.39994530243902,
		},
		Recipient: primitive.Coordinate{
			Latitude:  -7.7584860436179435,
			Longitude: 110.39994530243902,
		},
		Dimension: primitive.Dimension{
			Height: 10,
			Width:  10,
			Depth:  10,
		},
		Weight: 1,
	}

	// test positive case
	t.Run("positive case", func(t *testing.T) {
		err := shipping.ValidateEstimateRequest(request)
		if err != nil {
			t.Errorf("expect error nil, but got %T instead", err)
		}
	})

	// test Sender
	t.Run("Sender.latitude", func(t *testing.T) {
		// arrange
		mock := request
		var requestValidationError *business.RequestValidationError

		t.Run("required", func(t *testing.T) {
			mock.Sender = primitive.Coordinate{}

			err := shipping.ValidateEstimateRequest(mock)

			if err == nil {
				t.Errorf("expect error" +
					" when the given sender is empty object, but got nil instead")
			}
			if !errors.As(err, &requestValidationError) {
				t.Errorf("expect error as *business.RequestValidationError"+
					" when the given sender is empty object, but got %T instead", err)
			}
		})

		t.Run("given latitude is invalid", func(t *testing.T) {
			mock.Sender.Latitude = 91
			err := shipping.ValidateEstimateRequest(mock)
			if err == nil {
				t.Errorf("expect error" +
					" when the given latitude is invalid, but got nil instead")
			}

			mock.Sender.Latitude = -91
			err = shipping.ValidateEstimateRequest(mock)
			if err == nil {
				t.Errorf("expect error" +
					" when the given latitude is invalid, but got nil instead")
			}
		})
	})

	// test Sender.longitude
	t.Run("Sender.longitude", func(t *testing.T) {
		// arrange
		mock := request
		var requestValidationError *business.RequestValidationError

		t.Run("required", func(t *testing.T) {
			mock.Sender = primitive.Coordinate{}

			err := shipping.ValidateEstimateRequest(mock)

			if err == nil {
				t.Errorf("expect error" +
					" when the given sender is empty object, but got nil instead")
			}
			if !errors.As(err, &requestValidationError) {
				t.Errorf("expect error as *business.RequestValidationError"+
					" when the given sender is empty object, but got %T instead", err)
			}
		})

		t.Run("given longitude is invalid", func(t *testing.T) {
			mock.Sender.Longitude = 181
			err := shipping.ValidateEstimateRequest(mock)
			if err == nil {
				t.Errorf("expect error" +
					" when the given longitude is invalid, but got nil instead")
			}

			mock.Sender.Longitude = -181
			err = shipping.ValidateEstimateRequest(mock)
			if err == nil {
				t.Errorf("expect error" +
					" when the given longitude is invalid, but got nil instead")
			}
		})
	})

	// test Recipient
	t.Run("Recipient", func(t *testing.T) {
		// arrange
		mock := request
		var requestValidationError *business.RequestValidationError

		t.Run("required", func(t *testing.T) {
			mock.Recipient = primitive.Coordinate{}

			err := shipping.ValidateEstimateRequest(mock)

			if err == nil {
				t.Errorf("expect error" +
					" when the given recipient is empty object, but got nil instead")
			}
			if !errors.As(err, &requestValidationError) {
				t.Errorf("expect error as *business.RequestValidationError"+
					" when the given recipient is empty object, but got %T instead", err)
			}
		})
	})

	// test Recipient.latitude
	t.Run("Recipient.latitude", func(t *testing.T) {
		// arrange
		mock := request
		var requestValidationError *business.RequestValidationError

		t.Run("required", func(t *testing.T) {
			mock.Recipient = primitive.Coordinate{}

			err := shipping.ValidateEstimateRequest(mock)

			if err == nil {
				t.Errorf("expect error" +
					" when the given recipient is empty object, but got nil instead")
			}
			if !errors.As(err, &requestValidationError) {
				t.Errorf("expect error as *business.RequestValidationError"+
					" when the given recipient is empty object, but got %T instead", err)
			}
		})

		t.Run("given latitude is invalid", func(t *testing.T) {
			mock.Recipient.Latitude = 91
			err := shipping.ValidateEstimateRequest(mock)
			if err == nil {
				t.Errorf("expect error" +
					" when the given latitude is invalid, but got nil instead")
			}

			mock.Recipient.Latitude = -91
			err = shipping.ValidateEstimateRequest(mock)
			if err == nil {
				t.Errorf("expect error" +
					" when the given latitude is invalid, but got nil instead")
			}
		})
	})

	// test Recipient.longitude
	t.Run("Recipient.longitude", func(t *testing.T) {
		// arrange
		mock := request
		var requestValidationError *business.RequestValidationError

		t.Run("required", func(t *testing.T) {
			mock.Recipient = primitive.Coordinate{}

			err := shipping.ValidateEstimateRequest(mock)

			if err == nil {
				t.Errorf("expect error" +
					" when the given recipient is empty object, but got nil instead")
			}
			if !errors.As(err, &requestValidationError) {
				t.Errorf("expect error as *business.RequestValidationError"+
					" when the given recipient is empty object, but got %T instead", err)
			}
		})

		t.Run("given longitude is invalid", func(t *testing.T) {
			mock.Recipient.Longitude = 181
			err := shipping.ValidateEstimateRequest(mock)
			if err == nil {
				t.Errorf("expect error" +
					" when the given longitude is invalid, but got nil instead")
			}

			mock.Recipient.Longitude = -181
			err = shipping.ValidateEstimateRequest(mock)
			if err == nil {
				t.Errorf("expect error" +
					" when the given longitude is invalid, but got nil instead")
			}
		})
	})

	// test Dimension
	t.Run("Dimension", func(t *testing.T) {
		// arrange
		mock := request
		var requestValidationError *business.RequestValidationError

		t.Run("required", func(t *testing.T) {
			mock.Dimension = primitive.Dimension{}

			err := shipping.ValidateEstimateRequest(mock)

			if err == nil {
				t.Errorf("expect error" +
					" when the given dimension is empty object, but got nil instead")
			}
			if !errors.As(err, &requestValidationError) {
				t.Errorf("expect error as *business.RequestValidationError"+
					" when the given dimension is empty object, but got %T instead", err)
			}
		})
	})

	// test Dimension.Height
	t.Run("Dimension.Height", func(t *testing.T) {
		// arrange
		mock := request
		var requestValidationError *business.RequestValidationError

		t.Run("greater than 0", func(t *testing.T) {
			mock.Dimension.Height = -1

			err := shipping.ValidateEstimateRequest(mock)

			if err == nil {
				t.Errorf("expect error" +
					" when the given dimension.height is less than 0, but got nil instead")
			}
			if !errors.As(err, &requestValidationError) {
				t.Errorf("expect error as *business.RequestValidationError"+
					" when the given dimension.height is less than 0, but got %T instead", err)
			}
		})
	})

	// test Dimension.Width
	t.Run("Dimension.Width", func(t *testing.T) {
		// arrange
		mock := request
		var requestValidationError *business.RequestValidationError

		t.Run("greater than 0", func(t *testing.T) {
			mock.Dimension.Width = -1

			err := shipping.ValidateEstimateRequest(mock)

			if err == nil {
				t.Errorf("expect error" +
					" when the given dimension.width is less than 0, but got nil instead")
			}
			if !errors.As(err, &requestValidationError) {
				t.Errorf("expect error as *business.RequestValidationError"+
					" when the given dimension.width is less than 0, but got %T instead", err)
			}
		})
	})

	// test Dimension.Depth
	t.Run("Dimension.Depth", func(t *testing.T) {
		// arrange
		mock := request
		var requestValidationError *business.RequestValidationError

		t.Run("greater than 0", func(t *testing.T) {
			mock.Dimension.Depth = -1

			err := shipping.ValidateEstimateRequest(mock)

			if err == nil {
				t.Errorf("expect error" +
					" when the given dimension.depth is less than 0, but got nil instead")
			}
			if !errors.As(err, &requestValidationError) {
				t.Errorf("expect error as *business.RequestValidationError"+
					" when the given dimension.depth is less than 0, but got %T instead", err)
			}
		})
	})

	// test Height
	t.Run("Height", func(t *testing.T) {
		// arrange
		mock := request
		var requestValidationError *business.RequestValidationError

		t.Run("greater than 0", func(t *testing.T) {
			mock.Weight = 0

			err := shipping.ValidateEstimateRequest(mock)

			if err == nil {
				t.Errorf("expect error" +
					" when the given weight is equal with 0, but got nil instead")
			}
			if !errors.As(err, &requestValidationError) {
				t.Errorf("expect error as *business.RequestValidationError"+
					" when the given weight is equal with 0, but got %T instead", err)
			}

			mock.Weight = -1

			err = shipping.ValidateEstimateRequest(mock)

			if err == nil {
				t.Errorf("expect error" +
					" when the given weight is less than 0, but got nil instead")
			}
			if !errors.As(err, &requestValidationError) {
				t.Errorf("expect error as *business.RequestValidationError"+
					" when the given weight is less than 0, but got %T instead", err)
			}
		})
	})
}
