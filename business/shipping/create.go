package shipping

import (
	"context"
	"errors"
	"mock-shipping-provider/business"
	"mock-shipping-provider/primitive"
	"regexp"
	"strconv"
)

// Create handle the business logic for
// creating shipping order
func (d *Dependency) Create(ctx context.Context, request business.CreateRequest) (business.CreateResponse, error) {
	// validate the request
	if err := ValidateCreateRequest(request); err != nil {
		return business.CreateResponse{}, err
	}

	// TODO implement me
	panic("implement me")
}

func ValidateCreateRequest(request business.CreateRequest) *business.RequestValidationError {
	// Provider
	if request.Provider == primitive.ProviderUnspecified {
		return &business.RequestValidationError{
			Reason: "invalid shipping provider",
		}
	}

	// Sender.Name
	if request.Sender.Name == "" {
		return &business.RequestValidationError{
			Reason: "sender.name is required",
		}
	}

	if len(request.Sender.Name) < 5 {
		return &business.RequestValidationError{
			Reason: "sender.name is to short, at least 5 characters length must be given",
		}
	}

	if len(request.Sender.Name) > 255 {
		return &business.RequestValidationError{
			Reason: "sender.name must be less than 255 characters length",
		}
	}

	if ok := regexp.MustCompile(primitive.AddressNamePattern).MatchString(request.Sender.Name); !ok {
		return &business.RequestValidationError{
			Reason: "sender.name contains prohibited characters",
		}
	}

	// Sender.PhoneNumber
	if request.Sender.PhoneNumber == "" {
		return &business.RequestValidationError{
			Reason: "sender.phone_number is required",
		}
	}

	if len(request.Sender.PhoneNumber) >= 20 {
		return &business.RequestValidationError{
			Reason: "sender.phone_number must be less than 20 characters length",
		}
	}

	if ok := regexp.MustCompile(primitive.AddressPhoneNumbePattern).MatchString(request.Sender.PhoneNumber); !ok {
		return &business.RequestValidationError{
			Reason: "sender.phone_number contains prohibited characters",
		}
	}

	// Sender.Address
	if request.Sender.Address == "" {
		return &business.RequestValidationError{
			Reason: "sender.address is required",
		}
	}

	if len(request.Sender.Address) > 500 {
		return &business.RequestValidationError{
			Reason: "sender.address must be less than equal 500 characters length",
		}
	}

	if len(request.Sender.Address) < 10 {
		return &business.RequestValidationError{
			Reason: "sender.address must be greater than equal 10 characters length",
		}
	}

	// Sender.City
	if request.Sender.City == "" {
		return &business.RequestValidationError{
			Reason: "sender.city is required",
		}
	}

	if len(request.Sender.City) > 500 {
		return &business.RequestValidationError{
			Reason: "sender.city too long. must be less than 500 characters length",
		}
	}

	// Sender.State
	if request.Sender.State == "" {
		return &business.RequestValidationError{
			Reason: "sender.state is required",
		}
	}

	if len(request.Sender.State) > 255 {
		return &business.RequestValidationError{
			Reason: "sender.state is too long. must be less than equal 255 characters length",
		}
	}

	// Sender.Country
	if request.Sender.Country == "" {
		return &business.RequestValidationError{
			Reason: "sender.country is required",
		}
	}

	if len(request.Sender.Country) > 255 {
		return &business.RequestValidationError{
			Reason: "sender.country is too long. must be less than equal 255 characters length",
		}
	}

	// Sender.PostalCode
	if request.Sender.PostalCode == "" {
		return &business.RequestValidationError{
			Reason: "sender.postal_code is required",
		}
	}

	if len(request.Sender.PostalCode) > 10 {
		return &business.RequestValidationError{
			Reason: "sender.postal_code is too long. must be less than equal 10 characters length",
		}
	}

	if _, err := strconv.ParseUint(request.Sender.PostalCode, 10, 64); err != nil {
		return &business.RequestValidationError{
			Reason: "sender.postal_code is not valid",
		}
	}

	// Recipient.Name
	if request.Recipient.Name == "" {
		return &business.RequestValidationError{
			Reason: "recipient.name is required",
		}
	}

	if len(request.Recipient.Name) < 5 {
		return &business.RequestValidationError{
			Reason: "recipient.name is to short, at least 5 characters length must be given",
		}
	}

	if len(request.Recipient.Name) > 255 {
		return &business.RequestValidationError{
			Reason: "recipient.name must be less than 255 characters length",
		}
	}

	if ok := regexp.MustCompile(primitive.AddressNamePattern).MatchString(request.Recipient.Name); !ok {
		return &business.RequestValidationError{
			Reason: "recipient.name contains prohibited characters",
		}
	}

	// Recipient.PhoneNumber
	if request.Recipient.PhoneNumber == "" {
		return &business.RequestValidationError{
			Reason: "recipient.phone_number is required",
		}
	}

	if len(request.Recipient.PhoneNumber) >= 20 {
		return &business.RequestValidationError{
			Reason: "recipient.phone_number must be less than 20 characters length",
		}
	}

	if ok := regexp.MustCompile(primitive.AddressPhoneNumbePattern).MatchString(request.Recipient.PhoneNumber); !ok {
		return &business.RequestValidationError{
			Reason: "recipient.phone_number contains prohibited characters",
		}
	}

	// Recipient.Address
	if request.Recipient.Address == "" {
		return &business.RequestValidationError{
			Reason: "recipient.address is required",
		}
	}

	if len(request.Recipient.Address) > 500 {
		return &business.RequestValidationError{
			Reason: "recipient.address must be less than equal 500 characters length",
		}
	}

	if len(request.Recipient.Address) < 10 {
		return &business.RequestValidationError{
			Reason: "recipient.address must be greater than equal 10 characters length",
		}
	}

	// Recipient.City
	if request.Recipient.City == "" {
		return &business.RequestValidationError{
			Reason: "recipient.city is required",
		}
	}

	if len(request.Recipient.City) > 500 {
		return &business.RequestValidationError{
			Reason: "recipient.city too long. must be less than 500 characters length",
		}
	}

	// Recipient.State
	if request.Recipient.State == "" {
		return &business.RequestValidationError{
			Reason: "recipient.state is required",
		}
	}

	if len(request.Recipient.State) > 255 {
		return &business.RequestValidationError{
			Reason: "recipient.state is too long. must be less than equal 255 characters length",
		}
	}

	// Recipient.Country
	if request.Recipient.Country == "" {
		return &business.RequestValidationError{
			Reason: "recipient.country is required",
		}
	}

	if len(request.Recipient.Country) > 255 {
		return &business.RequestValidationError{
			Reason: "recipient.country is too long. must be less than equal 255 characters length",
		}
	}

	// Recipient.PostalCode
	if request.Recipient.PostalCode == "" {
		return &business.RequestValidationError{
			Reason: "recipient.postal_code is required",
		}
	}

	if len(request.Recipient.PostalCode) > 10 {
		return &business.RequestValidationError{
			Reason: "recipient.postal_code is too long. must be less than equal 10 characters length",
		}
	}

	if _, err := strconv.ParseUint(request.Recipient.PostalCode, 10, 64); err != nil {
		return &business.RequestValidationError{
			Reason: "recipient.postal_code is not valid",
		}
	}

	// Dimentsion
	if err := request.Dimension.Validate(); err != nil {
		if errors.Is(err, primitive.HeightIsLowerThanZero) {
			return &business.RequestValidationError{
				Reason: "dimension.height must be greater than 0",
			}
		}

		if errors.Is(err, primitive.WidthIsLowerThanZero) {
			return &business.RequestValidationError{
				Reason: "dimension.width must be greater than 0",
			}
		}

		if errors.Is(err, primitive.DepthIsLowerThanZero) {
			return &business.RequestValidationError{
				Reason: "dimenstion.depth must be greater than 0",
			}
		}
	}

	// Weight
	if request.Weight <= 0 {
		return &business.RequestValidationError{
			Reason: "weight must be greater than 0",
		}
	}

	// ItemDescription
	if request.ItemDescription == "" {
		return &business.RequestValidationError{
			Reason: "item_description is required",
		}
	}

	if len(request.ItemDescription) > 500 {
		return &business.RequestValidationError{
			Reason: "item_description is too long. must be less than equal 500 characters length",
		}
	}

	// ItemCategory
	if request.ItemCategory == "" {
		return &business.RequestValidationError{
			Reason: "item_category is required",
		}
	}

	if len(request.ItemCategory) > 500 {
		return &business.RequestValidationError{
			Reason: "item_category is too long. must be less than equal 500 characters length",
		}
	}

	return nil
}
