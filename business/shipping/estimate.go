package shipping

import (
	"context"
	"errors"
	"mock-shipping-provider/business"
	"mock-shipping-provider/primitive"
)

func (d *Dependency) Estimate(ctx context.Context, request business.EstimateRequest) ([]business.EstimateResult, error) {
	// TODO implement me
	panic("implement me")
}

// ValidateEstimateRequest handle an action to
// validate the estimate request body
func ValidateEstimateRequest(request business.EstimateRequest) *business.RequestValidationError {
	// Sender
	if request.Sender == (primitive.Coordinate{}) {
		return &business.RequestValidationError{
			Reason: "sender field is required",
		}
	}

	// Recipient
	if request.Recipient == (primitive.Coordinate{}) {
		return &business.RequestValidationError{
			Reason: "recipient field is required",
		}
	}

	// Dimension
	if request.Dimension == (primitive.Dimension{}) {
		return &business.RequestValidationError{
			Reason: "dimension field is required",
		}
	}

	if err := request.Dimension.Validate(); err != nil {
		if errors.Is(err, primitive.HeightIsLowerThanZero) {
			return &business.RequestValidationError{
				Reason: "dimension.height field must be greater or equal than 0",
			}
		}
		if errors.Is(err, primitive.WidthIsLowerThanZero) {
			return &business.RequestValidationError{
				Reason: "dimension.width field must be greater or equal than 0",
			}
		}
		if errors.Is(err, primitive.DepthIsLowerThanZero) {
			return &business.RequestValidationError{
				Reason: "dimension.depth fields must be greater or equal than 0",
			}
		}
	}

	// Weight
	if request.Weight <= 0 {
		return &business.RequestValidationError{
			Reason: "weight must be greater than 0",
		}
	}

	return nil
}
