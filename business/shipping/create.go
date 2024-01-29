package shipping

import (
	"context"
	"strconv"

	"mock-shipping-provider/business"
	"mock-shipping-provider/primitive"
	"mock-shipping-provider/repository"
	"mock-shipping-provider/repository/provider"
)

// Create handle the business logic for
// creating shipping order
func (d *Dependency) Create(ctx context.Context, request business.CreateRequest) (business.CreateResponse, error) {
	// validate the request
	if err := ValidateCreateRequest(request); err != nil {
		return business.CreateResponse{}, err
	}

	response := business.CreateResponse{
		ReferenceNumber: request.Sender.Name,
		AirWaybill:      request.Sender.PhoneNumber,
	}

	distance, isServiceable := d.distanceCalculation.Calculate(request.Sender.Coordinate, request.Recipient.Coordinate)
	if !isServiceable {
		return business.CreateResponse{}, business.ErrNotServiceable
	}

	if request.Provider == primitive.ProviderJNE {

		JNE := provider.NewJneCalculation(&primitive.Rate{
			PerKilogram:      200,
			PerKilometer:     300,
			PerCmCubic:       400,
			KilometerPerHour: 60,
		})

		response.Price = JNE.CalculatePrice(distance, request.Dimension, request.Weight)
		response.Hours = uint64(JNE.CalculateTimeOfArrival(distance))
	}

	if request.Provider == primitive.ProviderJNT {

		JNT := provider.NewJntCalculation(&primitive.Rate{
			PerKilogram:      200,
			PerKilometer:     300,
			PerCmCubic:       400,
			KilometerPerHour: 60,
		})

		response.Price = JNT.CalculatePrice(distance, request.Dimension, request.Weight)
		response.Hours = uint64(JNT.CalculateTimeOfArrival(distance))
	}

	if request.Provider == primitive.ProviderSiCepat {

		SiCepat := provider.NewSicepatCalculation(&primitive.Rate{
			PerKilogram:      200,
			PerKilometer:     300,
			PerCmCubic:       400,
			KilometerPerHour: 60,
		})

		response.Price = SiCepat.CalculatePrice(distance, request.Dimension, request.Weight)
		response.Hours = uint64(SiCepat.CalculateTimeOfArrival(distance))
	}

	if request.Provider == primitive.ProviderAnterAja {

		AnterAja := provider.NewAnterajaCalculation(&primitive.Rate{
			PerKilogram:      200,
			PerKilometer:     300,
			PerCmCubic:       400,
			KilometerPerHour: 60,
		})

		response.Price = AnterAja.CalculatePrice(distance, request.Dimension, request.Weight)
		response.Hours = uint64(AnterAja.CalculateTimeOfArrival(distance))
	}

	orderHistory, err := d.orderLogRepository.Get(ctx, response.ReferenceNumber, response.AirWaybill)
	if err != nil {
		return business.CreateResponse{}, err
	}

	orderLog := repository.LogEntry{
		ReferenceNumber: response.ReferenceNumber,
		AirWaybill:      response.AirWaybill,
		History:         orderHistory,
	}

	if err := d.orderLogRepository.Create(ctx, orderLog); err != nil {
		return business.CreateResponse{}, err
	}

	return response, nil
}

func ValidateCreateRequest(request business.CreateRequest) *business.RequestValidationError {
	var issues []business.RequestValidationIssue

	// Provider
	if request.Provider == primitive.ProviderUnspecified {
		issues = append(issues, business.RequestValidationIssue{
			Code:    business.RequestValidationCodeInvalidValue,
			Field:   "provider",
			Message: "invalid shipping provider",
		})
	}

	// Sender.Name
	if request.Sender.Name == "" {
		issues = append(issues, business.RequestValidationIssue{
			Code:    business.RequestValidationCodeRequired,
			Field:   "sender.name",
			Message: "can not be empty",
		})
	} else {
		if len(request.Sender.Name) < 5 {
			issues = append(issues, business.RequestValidationIssue{
				Code:    business.RequestValidationCodeTooShort,
				Field:   "sender.name",
				Message: "minimum of 5 characters",
			})
		}

		if len(request.Sender.Name) > 255 {
			issues = append(issues, business.RequestValidationIssue{
				Code:    business.RequestValidationCodeTooLong,
				Field:   "sender.name",
				Message: "maximum of 255 characters",
			})
		}

		if ok := primitive.AddressNamePattern.MatchString(request.Sender.Name); !ok {
			issues = append(issues, business.RequestValidationIssue{
				Code:    business.RequestValidationCodeProhibitedValue,
				Field:   "sender.name",
				Message: "must be alphanumeric",
			})
		}
	}

	// Sender.PhoneNumber
	if request.Sender.PhoneNumber == "" {
		issues = append(issues, business.RequestValidationIssue{
			Code:    business.RequestValidationCodeRequired,
			Field:   "sender.phone_number",
			Message: "can not be empty",
		})
	} else {
		if len(request.Sender.PhoneNumber) >= 20 {
			issues = append(issues, business.RequestValidationIssue{
				Code:    business.RequestValidationCodeTooLong,
				Field:   "sender.phone_number",
				Message: "maximum of 20 characters",
			})
		}

		if ok := primitive.AddressPhoneNumberPattern.MatchString(request.Sender.PhoneNumber); !ok {
			issues = append(issues, business.RequestValidationIssue{
				Code:    business.RequestValidationCodeProhibitedValue,
				Field:   "sender.phone_number",
				Message: "must be numeric",
			})
		}
	}

	// Sender.Address
	if request.Sender.Address == "" {
		issues = append(issues, business.RequestValidationIssue{
			Code:    business.RequestValidationCodeRequired,
			Field:   "sender.address",
			Message: "can not be empty",
		})

	} else {

		if len(request.Sender.Address) > 500 {
			issues = append(issues, business.RequestValidationIssue{
				Code:    business.RequestValidationCodeTooLong,
				Field:   "sender.address",
				Message: "maximum of 500 characters",
			})
		}

		if len(request.Sender.Address) < 10 {
			issues = append(issues, business.RequestValidationIssue{
				Code:    business.RequestValidationCodeTooShort,
				Field:   "sender.address",
				Message: "minimum of 10 characters",
			})
		}
	}

	// Sender.City
	if request.Sender.City == "" {
		issues = append(issues, business.RequestValidationIssue{
			Code:    business.RequestValidationCodeRequired,
			Field:   "sender.city",
			Message: "can not be empty",
		})
	} else {
		if len(request.Sender.City) > 500 {
			issues = append(issues, business.RequestValidationIssue{
				Code:    business.RequestValidationCodeTooLong,
				Field:   "sender.city",
				Message: "can not be empty",
			})
		}
	}

	// Sender.State
	if request.Sender.State == "" {
		issues = append(issues, business.RequestValidationIssue{
			Code:    business.RequestValidationCodeRequired,
			Field:   "sender.state",
			Message: "sender.state is required",
		})
	} else {
		if len(request.Sender.State) > 255 {
			issues = append(issues, business.RequestValidationIssue{
				Code:    business.RequestValidationCodeTooLong,
				Field:   "sender.state",
				Message: "maximum of 255",
			})
		}

	}

	// Sender.Country
	if request.Sender.Country == "" {
		issues = append(issues, business.RequestValidationIssue{
			Code:    business.RequestValidationCodeRequired,
			Field:   "sender.country",
			Message: "can not be empty",
		})
	} else {
		if len(request.Sender.Country) > 255 {
			issues = append(issues, business.RequestValidationIssue{
				Code:    business.RequestValidationCodeTooLong,
				Field:   "sender.country",
				Message: "maximum of 255",
			})
		}
	}

	// Sender.PostalCode
	if request.Sender.PostalCode == "" {
		issues = append(issues, business.RequestValidationIssue{
			Code:    business.RequestValidationCodeRequired,
			Field:   "sender.postal_code",
			Message: "can not be empty",
		})
	} else {
		if len(request.Sender.PostalCode) > 10 {
			issues = append(issues, business.RequestValidationIssue{
				Code:    business.RequestValidationCodeTooLong,
				Field:   "sender.postal_code",
				Message: "maximum of 10 characters",
			})
		}

		if _, err := strconv.ParseUint(request.Sender.PostalCode, 10, 64); err != nil {
			issues = append(issues, business.RequestValidationIssue{
				Code:    business.RequestValidationCodeInvalidValue,
				Field:   "sender.postal_code",
				Message: "must be numeric",
			})
		}
	}

	// Recipient.Name
	if request.Recipient.Name == "" {
		issues = append(issues, business.RequestValidationIssue{
			Code:    business.RequestValidationCodeRequired,
			Field:   "recipient.name",
			Message: "can not be empty",
		})
	} else {
		if len(request.Recipient.Name) < 5 {
			issues = append(issues, business.RequestValidationIssue{
				Code:    business.RequestValidationCodeTooShort,
				Field:   "recipient.name",
				Message: "minimum of 5 characters",
			})
		}

		if len(request.Recipient.Name) > 255 {
			issues = append(issues, business.RequestValidationIssue{
				Code:    business.RequestValidationCodeTooLong,
				Field:   "recipient.name",
				Message: "maximum of 255 characters",
			})
		}

		if ok := primitive.AddressNamePattern.MatchString(request.Recipient.Name); !ok {
			issues = append(issues, business.RequestValidationIssue{
				Code:    business.RequestValidationCodeProhibitedValue,
				Field:   "recipient.name",
				Message: "must be alphanumeric",
			})
		}
	}

	// Recipient.PhoneNumber
	if request.Recipient.PhoneNumber == "" {
		issues = append(issues, business.RequestValidationIssue{
			Code:    business.RequestValidationCodeRequired,
			Field:   "recipient.phone_number",
			Message: "can not be empty",
		})
	} else {
		if len(request.Recipient.PhoneNumber) >= 20 {
			issues = append(issues, business.RequestValidationIssue{
				Code:    business.RequestValidationCodeTooLong,
				Field:   "recipient.phone_number",
				Message: "maximum of 20 characters",
			})
		}

		if ok := primitive.AddressPhoneNumberPattern.MatchString(request.Recipient.PhoneNumber); !ok {
			issues = append(issues, business.RequestValidationIssue{
				Code:    business.RequestValidationCodeProhibitedValue,
				Field:   "recipient.phone_number",
				Message: "must be numeric",
			})
		}
	}

	// Recipient.Address
	if request.Recipient.Address == "" {
		issues = append(issues, business.RequestValidationIssue{
			Code:    business.RequestValidationCodeRequired,
			Field:   "recipient.address",
			Message: "can not be empty",
		})
	} else {
		if len(request.Recipient.Address) > 500 {
			issues = append(issues, business.RequestValidationIssue{
				Code:    business.RequestValidationCodeTooLong,
				Field:   "recipient.address",
				Message: "maximum of 500 characters",
			})
		}

		if len(request.Recipient.Address) < 10 {
			issues = append(issues, business.RequestValidationIssue{
				Code:    business.RequestValidationCodeTooShort,
				Field:   "recipient.address",
				Message: "minimum of 10 characters",
			})
		}
	}

	// Recipient.City
	if request.Recipient.City == "" {
		issues = append(issues, business.RequestValidationIssue{
			Code:    business.RequestValidationCodeRequired,
			Field:   "recipient.city",
			Message: "can not be empty",
		})
	} else {
		if len(request.Recipient.City) > 500 {
			issues = append(issues, business.RequestValidationIssue{
				Code:    business.RequestValidationCodeTooLong,
				Field:   "recipient.city",
				Message: "maximum of 500 characters",
			})
		}
	}

	// Recipient.State
	if request.Recipient.State == "" {
		issues = append(issues, business.RequestValidationIssue{
			Code:    business.RequestValidationCodeRequired,
			Field:   "recipient.state",
			Message: "can not be empty",
		})
	} else {
		if len(request.Recipient.State) > 255 {
			issues = append(issues, business.RequestValidationIssue{
				Code:    business.RequestValidationCodeTooLong,
				Field:   "recipient.state",
				Message: "maximum of 255 characters",
			})
		}
	}

	// Recipient.Country
	if request.Recipient.Country == "" {
		issues = append(issues, business.RequestValidationIssue{
			Code:    business.RequestValidationCodeRequired,
			Field:   "recipient.country",
			Message: "can not be empty",
		})
	} else {
		if len(request.Recipient.Country) > 255 {
			issues = append(issues, business.RequestValidationIssue{
				Code:    business.RequestValidationCodeTooLong,
				Field:   "recipient.country",
				Message: "maximum of 255 characters",
			})
		}
	}

	// Recipient.PostalCode
	if request.Recipient.PostalCode == "" {
		issues = append(issues, business.RequestValidationIssue{
			Code:    business.RequestValidationCodeRequired,
			Field:   "recipient.postal_code",
			Message: "can not be empty",
		})
	} else {
		if len(request.Recipient.PostalCode) > 10 {
			issues = append(issues, business.RequestValidationIssue{
				Code:    business.RequestValidationCodeTooLong,
				Field:   "recipient.postal_code",
				Message: "maximum of 10 characters",
			})

		}

		if _, err := strconv.ParseUint(request.Recipient.PostalCode, 10, 64); err != nil {
			issues = append(issues, business.RequestValidationIssue{
				Code:    business.RequestValidationCodeInvalidValue,
				Field:   "recipient.postal_code",
				Message: "must be numeric",
			})
		}
	}

	// Dimension
	if err := request.Dimension.Validate(); err != nil {
		issues = append(issues, business.RequestValidationIssue{
			Code:    business.RequestValidationCodeInvalidValue,
			Field:   "dimension",
			Message: err.Error(),
		})
	}

	// Weight
	if request.Weight <= 0 {
		issues = append(issues, business.RequestValidationIssue{
			Code:    business.RequestValidationCodeInvalidValue,
			Field:   "weight",
			Message: "must be greater than 0",
		})

	}

	// ItemDescription
	if request.ItemDescription == "" {
		issues = append(issues, business.RequestValidationIssue{
			Code:    business.RequestValidationCodeRequired,
			Field:   "item_description",
			Message: "item_description is required",
		})
	} else {
		if len(request.ItemDescription) > 500 {
			issues = append(issues, business.RequestValidationIssue{
				Code:    business.RequestValidationCodeRequired,
				Field:   "item_description",
				Message: "maximum of 500 characters",
			})
		}
	}

	// ItemCategory
	if request.ItemCategory == "" {
		issues = append(issues, business.RequestValidationIssue{
			Code:    business.RequestValidationCodeRequired,
			Field:   "item_category",
			Message: "can not be empty",
		})
	} else {
		if len(request.ItemCategory) > 500 {
			issues = append(issues, business.RequestValidationIssue{
				Code:    business.RequestValidationCodeRequired,
				Field:   "item_category",
				Message: "maximum of 500 characters",
			})
		}
	}

	if len(issues) > 0 {
		return &business.RequestValidationError{Issues: issues}
	}

	return nil
}
