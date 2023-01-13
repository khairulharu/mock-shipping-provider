package shipping_test

import (
	"errors"
	"fmt"
	"mock-shipping-provider/business"
	"mock-shipping-provider/business/shipping"
	"mock-shipping-provider/primitive"
	"testing"
)

func TestValidateCreateRequest(t *testing.T) {
	// Provide the example correct value
	request := business.CreateRequest{
		Provider: primitive.ProviderJNT,
		Sender: primitive.Address{
			Name:        "Spiderman",
			PhoneNumber: "+62123456789",
			Address:     "Jl. Kenangan",
			City:        "Sleman",
			State:       "Yogyakarta",
			Country:     "Indonesia",
			PostalCode:  "55281",
			Coordinate: primitive.Coordinate{
				Latitude:  -7.7584860436179435,
				Longitude: 110.39994530243902,
			},
		},
		Recipient: primitive.Address{
			Name:        "Wanda",
			PhoneNumber: "+62897654321",
			Address:     "Jl. Mangga",
			City:        "Sleman",
			State:       "Yogyakarta",
			Country:     "Indonesia",
			PostalCode:  "55281",
			Coordinate: primitive.Coordinate{
				Latitude:  -7.7584860436179435,
				Longitude: 110.39994530243902,
			},
		},
		Dimension: primitive.Dimension{
			Height: 20,
			Depth:  10,
			Width:  10,
		},
		Weight:          100,
		ItemDescription: "blablablabla",
		ItemCategory:    "Electronic",
		Fragile:         true,
	}

	// positive test case
	t.Run("positive case", func(t *testing.T) {
		err := shipping.ValidateCreateRequest(request)
		if err != nil {
			t.Errorf("expect error nil, but got %T instead", err)
		}
	})

	// test Provider
	t.Run("Provider", func(t *testing.T) {
		// arrange
		mock := request
		var requestValidationError *business.RequestValidationError

		t.Run("invalid provider", func(t *testing.T) {
			// mock
			mock.Provider = 0

			// action
			err := shipping.ValidateCreateRequest(mock)

			// assert
			if err == nil {
				t.Errorf("expect error" +
					" when the given provider is invalid, but got nil instead")
			}
			if !errors.As(err, &requestValidationError) {
				t.Errorf("expect error as *business.RequestValidationError"+
					" when the given provider is invalid, but got %T instead", err)
			}
		})
	})

	// Sender.Name
	t.Run("Sender.Name", func(t *testing.T) {
		// arrange
		mock := request
		var requestValidationError *business.RequestValidationError

		t.Run("required", func(t *testing.T) {
			mock.Sender.Name = ""

			err := shipping.ValidateCreateRequest(mock)

			if err == nil {
				t.Errorf("expect error" +
					" when the given sender.name is empty, but got nil instead")
			}
			if !errors.As(err, &requestValidationError) {
				t.Errorf("expect error as *business.RequestValidationError"+
					" when the given sender.name is empty, but got %T instead", err)
			}
		})

		t.Run("less than 5 characters length", func(t *testing.T) {
			mock.Sender.Name = "a"

			err := shipping.ValidateCreateRequest(mock)

			if err == nil {
				t.Errorf("expect error" +
					" when the given sender.name is less than 5 characters length, but got nil instead")
			}
			if !errors.As(err, &requestValidationError) {
				t.Errorf("expect error as *business.RequestValidationError"+
					" when the given sender.name is less than 5 characters length, but got %T instead", err)
			}
		})

		t.Run("greater than 255 characters length", func(t *testing.T) {
			for i := 0; i < 26; i++ {
				mock.Sender.Name += "aaaaaaaaaaa"
			}

			err := shipping.ValidateCreateRequest(mock)

			if err == nil {
				t.Errorf("expect error" +
					" when the given sender.name is greater than 255 characters length, but got nil instead")
			}
			if !errors.As(err, &requestValidationError) {
				t.Errorf("expect error as *business.RequestValidationError"+
					" when the given sender.name is greater than 255 characters length, but got %T instead", err)
			}
		})

		t.Run("prohibited characters", func(t *testing.T) {
			mock.Sender.Name = "@#$%^!^&*"

			err := shipping.ValidateCreateRequest(mock)

			if err == nil {
				t.Errorf("expect error" +
					" when the given sender.name contains prohibited characters, but got nil instead")
			}
			if !errors.As(err, &requestValidationError) {
				t.Errorf("expect error as *business.RequestValidationError"+
					" when the given sender.name contains prohibited characters, but got %T instead", err)
			}
		})
	})

	// Sender.PhoneNumber
	t.Run("Sender.PhoneNumber", func(t *testing.T) {
		// arrange
		mock := request
		var requestValidationError *business.RequestValidationError

		t.Run("required", func(t *testing.T) {
			mock.Sender.PhoneNumber = ""

			err := shipping.ValidateCreateRequest(mock)

			if err == nil {
				t.Errorf("expect error" +
					" when the given sender.phone_number is empty, but got nil instead")
			}
			if !errors.As(err, &requestValidationError) {
				t.Errorf("expect error as *business.RequestValidationError"+
					" when the given sender.phone_number is empty, but got %T instead", err)
			}
		})

		t.Run("less than 20 characters length", func(t *testing.T) {
			mock.Sender.PhoneNumber = "1234672312389012301820830174102301283"

			err := shipping.ValidateCreateRequest(mock)

			if err == nil {
				t.Errorf("expect error" +
					" when the given sender.phone_number is greater than 20 characters length, but got nil instead")
			}
			if !errors.As(err, &requestValidationError) {
				t.Errorf("expect error as *business.RequestValidationError"+
					" when the given sender.phone_number greater than 20 characters length, but got %T instead", err)
			}
		})

		t.Run("valid phone number", func(t *testing.T) {
			mock.Sender.PhoneNumber = "abcdf123"

			err := shipping.ValidateCreateRequest(mock)

			if err == nil {
				t.Errorf("expect error" +
					" when the given sender.phone_number is invalid, but got nil instead")
			}
			if !errors.As(err, &requestValidationError) {
				t.Errorf("expect error as *business.RequestValidationError"+
					" when the given sender.phone_number is invalid, but got %T instead", err)
			}
		})
	})

	// test Sender.Address
	t.Run("Sender.Address", func(t *testing.T) {
		// arrange
		mock := request
		var requestValidationError *business.RequestValidationError

		t.Run("required", func(t *testing.T) {
			mock.Sender.Address = ""

			err := shipping.ValidateCreateRequest(mock)

			if err == nil {
				t.Errorf("expect error" +
					" when the given sender.address is empty, but got nil instead")
			}
			if !errors.As(err, &requestValidationError) {
				t.Errorf("expect error as *business.RequestValidationError"+
					" when the given sender.address is empty, but got %T instead", err)
			}
		})

		t.Run("less than 500 characters length", func(t *testing.T) {
			for i := 0; i < 50; i++ {
				mock.Sender.Address += "aaaaaaaaaaa"
			}

			err := shipping.ValidateCreateRequest(mock)

			if err == nil {
				t.Errorf("expect error" +
					" when the given sender.address is greater than 500 characters length, but got nil instead")
			}
			if !errors.As(err, &requestValidationError) {
				t.Errorf("expect error as *business.RequestValidationError"+
					" when the given sender.address is greater than 500 characters length, but got %T instead", err)
			}
		})

		t.Run("greater than 10 characters length", func(t *testing.T) {
			mock.Sender.Address = "abc"

			err := shipping.ValidateCreateRequest(mock)

			if err == nil {
				t.Errorf("expect error" +
					" when the given sender.address is less than 10 characters length, but got nil instead")
			}
			if !errors.As(err, &requestValidationError) {
				t.Errorf("expect error as *business.RequestValidationError"+
					" when the given sender.address is less than 10 characters length, but got %T instead", err)
			}
		})
	})

	// test Sender.City
	t.Run("Sender.City", func(t *testing.T) {
		// arrange
		mock := request
		var requestValidationError *business.RequestValidationError

		t.Run("required", func(t *testing.T) {
			mock.Sender.City = ""

			err := shipping.ValidateCreateRequest(mock)

			if err == nil {
				t.Errorf("expect error" +
					" when the given sender.city is empty, but got nil instead")
			}
			if !errors.As(err, &requestValidationError) {
				t.Errorf("expect error as *business.RequestValidationError"+
					" when the given sender.city is empty, but got %T instead", err)
			}
		})

		t.Run("less than 500 characters length", func(t *testing.T) {
			for i := 0; i < 50; i++ {
				mock.Sender.City += "aaaaaaaaaaa"
			}

			err := shipping.ValidateCreateRequest(mock)

			if err == nil {
				t.Errorf("expect error" +
					" when the given sender.city is greater than 500 characters length, but got nil instead")
			}
			if !errors.As(err, &requestValidationError) {
				t.Errorf("expect error as *business.RequestValidationError"+
					" when the given sender.city is greater than 500 characters length, but got %T instead", err)
			}
		})
	})

	// test Sender.State
	t.Run("Sender.State", func(t *testing.T) {
		// arrange
		mock := request
		var requestValidationError *business.RequestValidationError

		t.Run("required", func(t *testing.T) {
			mock.Sender.State = ""

			err := shipping.ValidateCreateRequest(mock)

			if err == nil {
				t.Errorf("expect error" +
					" when the given sender.state is empty, but got nil instead")
			}
			if !errors.As(err, &requestValidationError) {
				t.Errorf("expect error as *business.RequestValidationError"+
					" when the given sender.state is empty, but got %T instead", err)
			}
		})

		t.Run("less than 255 characters length", func(t *testing.T) {
			for i := 0; i < 26; i++ {
				mock.Sender.State += "aaaaaaaaaaaa"
			}

			err := shipping.ValidateCreateRequest(mock)

			if err == nil {
				t.Errorf("expect error" +
					" when the given sender.state is greater than 255 characters length, but got nil instead")
			}
			if !errors.As(err, &requestValidationError) {
				t.Errorf("expect error as *business.RequestValidationError"+
					" when the given sender.state is greater than 255 characters length, but got %T instead", err)
			}
		})
	})

	// test Sender.Country
	t.Run("Sender.Country", func(t *testing.T) {
		// arrange
		mock := request
		var requestValidationError *business.RequestValidationError

		t.Run("required", func(t *testing.T) {
			mock.Sender.Country = ""

			err := shipping.ValidateCreateRequest(mock)

			if err == nil {
				t.Errorf("expect error" +
					" when the given sender.country is empty, but got nil instead")
			}
			if !errors.As(err, &requestValidationError) {
				t.Errorf("expect error as *business.RequestValidationError"+
					" when the given sender.country is empty, but got %T instead", err)
			}
		})

		t.Run("less than 255 characters length", func(t *testing.T) {
			for i := 0; i < 26; i++ {
				mock.Sender.Country += "aaaaaaaaaaa"
			}

			err := shipping.ValidateCreateRequest(mock)

			if err == nil {
				t.Errorf("expect error" +
					" when the given sender.country is greater than 255 characters length, but got nil instead")
			}
			if !errors.As(err, &requestValidationError) {
				t.Errorf("expect error as *business.RequestValidationError"+
					" when the given sender.country is greater than 255 characters length, but got %T instead", err)
			}
		})
	})

	// test Sender.PostalCode
	t.Run("Sender.PostalCode", func(t *testing.T) {
		// arrange
		mock := request
		var requestValidationError *business.RequestValidationError

		t.Run("required", func(t *testing.T) {
			mock.Sender.PostalCode = ""

			err := shipping.ValidateCreateRequest(mock)

			if err == nil {
				t.Errorf("expect error" +
					" when the given sender.postal_code is empty, but got nil instead")
			}
			if !errors.As(err, &requestValidationError) {
				t.Errorf("expect error as *business.RequestValidationError"+
					" when the given sender.postal_code is empty, but got %T instead", err)
			}
		})

		t.Run("less than 10 characters length", func(t *testing.T) {
			mock.Sender.PostalCode = "1234567891312457"

			err := shipping.ValidateCreateRequest(mock)

			if err == nil {
				t.Errorf("expect error" +
					" when the given sender.postal_code is greater than 10 characters length, but got nil instead")
			}
			if !errors.As(err, &requestValidationError) {
				t.Errorf("expect error as *business.RequestValidationError"+
					" when the given sender.postal_code is greater than 10 characters length, but got %T instead", err)
			}
		})

		t.Run("valid postal code", func(t *testing.T) {
			mock.Sender.PostalCode = "ac16$"

			err := shipping.ValidateCreateRequest(mock)

			if err == nil {
				t.Errorf("expect error" +
					" when the given sender.postal_code is invalid, but got nil instead")
			}
			if !errors.As(err, &requestValidationError) {
				t.Errorf("expect error as *business.RequestValidationError"+
					" when the given sender.postal_code is invalid, but got %T instead", err)
			}
		})
	})

	// Recipient.Name
	t.Run("Recipient.Name", func(t *testing.T) {
		// arrange
		mock := request
		var requestValidationError *business.RequestValidationError

		t.Run("required", func(t *testing.T) {
			mock.Recipient.Name = ""

			err := shipping.ValidateCreateRequest(mock)

			if err == nil {
				t.Errorf("expect error" +
					" when the given recipient.name is empty, but got nil instead")
			}
			if !errors.As(err, &requestValidationError) {
				t.Errorf("expect error as *business.RequestValidationError"+
					" when the given recipient.name is empty, but got %T instead", err)
			}
		})

		t.Run("less than 5 characters length", func(t *testing.T) {
			mock.Recipient.Name = "a"

			err := shipping.ValidateCreateRequest(mock)

			if err == nil {
				t.Errorf("expect error" +
					" when the given recipient.name is less than 5 characters length, but got nil instead")
			}
			if !errors.As(err, &requestValidationError) {
				t.Errorf("expect error as *business.RequestValidationError"+
					" when the given recipient.name is less than 5 characters length, but got %T instead", err)
			}
		})

		t.Run("greater than 255 characters length", func(t *testing.T) {
			for i := 0; i < 26; i++ {
				mock.Recipient.Name += "aaaaaaaaaaa"
			}

			err := shipping.ValidateCreateRequest(mock)

			if err == nil {
				t.Errorf("expect error" +
					" when the given recipient.name is greater than 255 characters length, but got nil instead")
			}
			if !errors.As(err, &requestValidationError) {
				t.Errorf("expect error as *business.RequestValidationError"+
					" when the given recipient.name is greater than 255 characters length, but got %T instead", err)
			}
		})

		t.Run("prohibited characters", func(t *testing.T) {
			mock.Recipient.Name = "@#$%^!^&*"

			err := shipping.ValidateCreateRequest(mock)

			if err == nil {
				t.Errorf("expect error" +
					" when the given recipient.name contains prohibited characters, but got nil instead")
			}
			if !errors.As(err, &requestValidationError) {
				t.Errorf("expect error as *business.RequestValidationError"+
					" when the given recipient.name contains prohibited characters, but got %T instead", err)
			}
		})
	})

	// Recipient.PhoneNumber
	t.Run("Recipient.PhoneNumber", func(t *testing.T) {
		// arrange
		mock := request
		var requestValidationError *business.RequestValidationError

		t.Run("required", func(t *testing.T) {
			mock.Recipient.PhoneNumber = ""

			err := shipping.ValidateCreateRequest(mock)

			if err == nil {
				t.Errorf("expect error" +
					" when the given recipient.phone_number is empty, but got nil instead")
			}
			if !errors.As(err, &requestValidationError) {
				t.Errorf("expect error as *business.RequestValidationError"+
					" when the given recipient.phone_number is empty, but got %T instead", err)
			}
		})

		t.Run("less than 20 characters length", func(t *testing.T) {
			mock.Recipient.PhoneNumber = "1234672312389012301820830174102301283"

			err := shipping.ValidateCreateRequest(mock)

			if err == nil {
				t.Errorf("expect error" +
					" when the given recipient.phone_number is greater than 20 characters length, but got nil instead")
			}
			if !errors.As(err, &requestValidationError) {
				t.Errorf("expect error as *business.RequestValidationError"+
					" when the given recipient.phone_number greater than 20 characters length, but got %T instead", err)
			}
		})

		t.Run("valid phone number", func(t *testing.T) {
			mock.Recipient.PhoneNumber = "abcdf123"

			err := shipping.ValidateCreateRequest(mock)

			if err == nil {
				t.Errorf("expect error" +
					" when the given recipient.phone_number is invalid, but got nil instead")
			}
			if !errors.As(err, &requestValidationError) {
				t.Errorf("expect error as *business.RequestValidationError"+
					" when the given recipient.phone_number is invalid, but got %T instead", err)
			}
		})
	})

	// test Recipient.Address
	t.Run("Recipient.Address", func(t *testing.T) {
		// arrange
		mock := request
		var requestValidationError *business.RequestValidationError

		t.Run("required", func(t *testing.T) {
			mock.Recipient.Address = ""

			err := shipping.ValidateCreateRequest(mock)

			if err == nil {
				t.Errorf("expect error" +
					" when the given recipient.address is empty, but got nil instead")
			}
			if !errors.As(err, &requestValidationError) {
				t.Errorf("expect error as *business.RequestValidationError"+
					" when the given recipient.address is empty, but got %T instead", err)
			}
		})

		t.Run("less than 500 characters length", func(t *testing.T) {
			for i := 0; i < 50; i++ {
				mock.Recipient.Address += "aaaaaaaaaaa"
			}

			err := shipping.ValidateCreateRequest(mock)

			if err == nil {
				t.Errorf("expect error" +
					" when the given recipient.address is greater than 500 characters length, but got nil instead")
			}
			if !errors.As(err, &requestValidationError) {
				t.Errorf("expect error as *business.RequestValidationError"+
					" when the given recipient.address is greater than 500 characters length, but got %T instead", err)
			}
		})

		t.Run("greater than 10 characters length", func(t *testing.T) {
			mock.Recipient.Address = "abc"

			err := shipping.ValidateCreateRequest(mock)

			if err == nil {
				t.Errorf("expect error" +
					" when the given recipient.address is less than 10 characters length, but got nil instead")
			}
			if !errors.As(err, &requestValidationError) {
				t.Errorf("expect error as *business.RequestValidationError"+
					" when the given recipient.address is less than 10 characters length, but got %T instead", err)
			}
		})
	})

	// test Recipient.City
	t.Run("Recipient.City", func(t *testing.T) {
		// arrange
		mock := request
		var requestValidationError *business.RequestValidationError

		t.Run("required", func(t *testing.T) {
			mock.Recipient.City = ""

			err := shipping.ValidateCreateRequest(mock)

			if err == nil {
				t.Errorf("expect error" +
					" when the given recipient.city is empty, but got nil instead")
			}
			if !errors.As(err, &requestValidationError) {
				t.Errorf("expect error as *business.RequestValidationError"+
					" when the given recipient.city is empty, but got %T instead", err)
			}
		})

		t.Run("less than 500 characters length", func(t *testing.T) {
			for i := 0; i < 50; i++ {
				mock.Recipient.City += "aaaaaaaaaaa"
			}

			err := shipping.ValidateCreateRequest(mock)

			if err == nil {
				t.Errorf("expect error" +
					" when the given recipient.city is greater than 500 characters length, but got nil instead")
			}
			if !errors.As(err, &requestValidationError) {
				t.Errorf("expect error as *business.RequestValidationError"+
					" when the given recipient.city is greater than 500 characters length, but got %T instead", err)
			}
		})
	})

	// test Recipient.State
	t.Run("Recipient.State", func(t *testing.T) {
		// arrange
		mock := request
		var requestValidationError *business.RequestValidationError

		t.Run("required", func(t *testing.T) {
			mock.Recipient.State = ""

			err := shipping.ValidateCreateRequest(mock)

			if err == nil {
				t.Errorf("expect error" +
					" when the given recipient.state is empty, but got nil instead")
			}
			if !errors.As(err, &requestValidationError) {
				t.Errorf("expect error as *business.RequestValidationError"+
					" when the given recipient.state is empty, but got %T instead", err)
			}
		})

		t.Run("less than 255 characters length", func(t *testing.T) {
			for i := 0; i < 26; i++ {
				mock.Recipient.State += "aaaaaaaaaaaa"
			}

			err := shipping.ValidateCreateRequest(mock)

			if err == nil {
				t.Errorf("expect error" +
					" when the given recipient.state is greater than 255 characters length, but got nil instead")
			}
			if !errors.As(err, &requestValidationError) {
				t.Errorf("expect error as *business.RequestValidationError"+
					" when the given recipient.state is greater than 255 characters length, but got %T instead", err)
			}
		})
	})

	// test Recipient.Country
	t.Run("Recipient.Country", func(t *testing.T) {
		// arrange
		mock := request
		var requestValidationError *business.RequestValidationError

		t.Run("required", func(t *testing.T) {
			mock.Recipient.Country = ""

			err := shipping.ValidateCreateRequest(mock)

			if err == nil {
				t.Errorf("expect error" +
					" when the given recipient.country is empty, but got nil instead")
			}
			if !errors.As(err, &requestValidationError) {
				t.Errorf("expect error as *business.RequestValidationError"+
					" when the given recipient.country is empty, but got %T instead", err)
			}
		})

		t.Run("less than 255 characters length", func(t *testing.T) {
			for i := 0; i < 26; i++ {
				mock.Recipient.Country += "aaaaaaaaaaa"
			}

			err := shipping.ValidateCreateRequest(mock)

			if err == nil {
				t.Errorf("expect error" +
					" when the given recipient.country is greater than 255 characters length, but got nil instead")
			}
			if !errors.As(err, &requestValidationError) {
				t.Errorf("expect error as *business.RequestValidationError"+
					" when the given recipient.country is greater than 255 characters length, but got %T instead", err)
			}
		})
	})

	// test Recipient.PostalCode
	t.Run("Recipient.PostalCode", func(t *testing.T) {
		// arrange
		mock := request
		var requestValidationError *business.RequestValidationError

		t.Run("required", func(t *testing.T) {
			mock.Recipient.PostalCode = ""

			err := shipping.ValidateCreateRequest(mock)

			if err == nil {
				t.Errorf("expect error" +
					" when the given recipient.postal_code is empty, but got nil instead")
			}
			if !errors.As(err, &requestValidationError) {
				t.Errorf("expect error as *business.RequestValidationError"+
					" when the given recipient.postal_code is empty, but got %T instead", err)
			}
		})

		t.Run("less than 10 characters length", func(t *testing.T) {
			mock.Recipient.PostalCode = "1234567891312457"

			err := shipping.ValidateCreateRequest(mock)

			if err == nil {
				t.Errorf("expect error" +
					" when the given recipient.postal_code is greater than 10 characters length, but got nil instead")
			}
			if !errors.As(err, &requestValidationError) {
				t.Errorf("expect error as *business.RequestValidationError"+
					" when the given recipient.postal_code is greater than 10 characters length, but got %T instead", err)
			}
		})

		t.Run("valid postal code", func(t *testing.T) {
			mock.Recipient.PostalCode = "ac16$"

			err := shipping.ValidateCreateRequest(mock)

			if err == nil {
				t.Errorf("expect error" +
					" when the given recipient.postal_code is invalid, but got nil instead")
			}
			if !errors.As(err, &requestValidationError) {
				t.Errorf("expect error as *business.RequestValidationError"+
					" when the given recipient.postal_code is invalid, but got %T instead", err)
			}
		})
	})

	// test Dimension.Height
	t.Run("Dimension.Height", func(t *testing.T) {
		// arrange
		mock := request
		var requestValidationError *business.RequestValidationError

		t.Run("greater than 0", func(t *testing.T) {
			mock.Dimension.Height = -2

			err := shipping.ValidateCreateRequest(mock)

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
			mock.Dimension.Width = -2

			err := shipping.ValidateCreateRequest(mock)

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
			mock.Dimension.Depth = -2

			err := shipping.ValidateCreateRequest(mock)

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

	// test Weight
	t.Run("Weight", func(t *testing.T) {
		// arrange
		mock := request
		var requestValidationError *business.RequestValidationError

		t.Run("greater than 0", func(t *testing.T) {
			mock.Weight = 0

			err := shipping.ValidateCreateRequest(mock)

			if err == nil {
				t.Errorf("expect error" +
					" when the given weight is equal 0, but got nil instead")
			}
			if !errors.As(err, &requestValidationError) {
				t.Errorf("expect error as *business.RequestValidationError"+
					" when the given weight is equal 0, but got %T instead", err)
			}

			mock.Weight = -1

			err = shipping.ValidateCreateRequest(mock)

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

	// test ItemDescription
	t.Run("ItemDescription", func(t *testing.T) {
		// arrange
		mock := request
		var requestValidationError *business.RequestValidationError

		t.Run("required", func(t *testing.T) {
			mock.ItemDescription = ""

			err := shipping.ValidateCreateRequest(mock)

			if err == nil {
				t.Errorf("expect error" +
					" when the given item_description is empty, but got nil instead")
			}
			if !errors.As(err, &requestValidationError) {
				t.Errorf("expect error as *business.RequestValidationError"+
					" when the given item_description is empty, but got %T instead", err)
			}
		})

		t.Run("less than 500 characters length", func(t *testing.T) {
			for i := 0; i < 50; i++ {
				mock.ItemDescription += "aaaaaaaaaaaa"
			}

			err := shipping.ValidateCreateRequest(mock)

			if err == nil {
				fmt.Println("panjangnya adalah", len(mock.ItemDescription))
				t.Errorf("expect error" +
					" when the given item_description is greater than 500 characters length, but got nil instead")
			}
			if !errors.As(err, &requestValidationError) {
				t.Errorf("expect error as *business.RequestValidationError"+
					" when the given item_description is greater than 500 characters length, but got %T instead", err)
			}
		})
	})

	// test ItemCategory
	t.Run("ItemCategory", func(t *testing.T) {
		// arrange
		mock := request
		var requestValidationError *business.RequestValidationError

		t.Run("required", func(t *testing.T) {
			mock.ItemCategory = ""

			err := shipping.ValidateCreateRequest(mock)

			if err == nil {
				t.Errorf("expect error" +
					" when the given item_category is empty, but got nil instead")
			}
			if !errors.As(err, &requestValidationError) {
				t.Errorf("expect error as *business.RequestValidationError"+
					" when the given item_category is empty, but got %T instead", err)
			}
		})

		t.Run("less than 500 characters length", func(t *testing.T) {
			for i := 0; i < 50; i++ {
				mock.ItemCategory += "aaaaaaaaaaa"
			}

			err := shipping.ValidateCreateRequest(mock)

			if err == nil {
				t.Errorf("expect error" +
					" when the given item_category is greater than 500 characters length, but got nil instead")
			}
			if !errors.As(err, &requestValidationError) {
				t.Errorf("expect error as *business.RequestValidationError"+
					" when the given item_category is greater than 500 characters length, but got %T instead", err)
			}
		})
	})
}
