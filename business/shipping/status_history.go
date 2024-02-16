package shipping

import (
	"context"
	"mock-shipping-provider/business"
)

func (d *Dependency) StatusHistory(ctx context.Context, request business.StatusRequest) (business.StatusHistoryResponse, error) {

	if err := ValidateStatusHistoryRequest(request); err != nil {
		return business.StatusHistoryResponse{}, err
	}

	var history []business.StatusHistory

	orderHistory, err := d.orderLogRepository.Get(ctx, request.ReferenceNumber, request.AirWaybill)
	if err != nil {
		return business.StatusHistoryResponse{}, err
	}

	for _, val := range orderHistory {
		history = append(history, business.StatusHistory{
			Status:    val.StatusCode,
			Timestamp: val.Timestamp,
			Note:      val.Note,
		})
	}
	return business.StatusHistoryResponse{
		ReferenceNumber: request.ReferenceNumber,
		AirWaybill:      request.AirWaybill,
		History:         history,
	}, nil

}

func ValidateStatusHistoryRequest(request business.StatusRequest) *business.RequestValidationError {
	var issues []business.RequestValidationIssue

	//Airwaybill
	if request.AirWaybill == ("") {
		issues = append(issues, business.RequestValidationIssue{
			Code:    business.RequestValidationCodeRequired,
			Field:   "airWaybill",
			Message: "can not be empty",
		})
	}

	if request.ReferenceNumber == ("") {
		issues = append(issues, business.RequestValidationIssue{
			Code:    business.RequestValidationCodeRequired,
			Field:   "referenceNumber",
			Message: "can not be empty",
		})
	}

	//Reference Number

	if len(issues) > 0 {
		return &business.RequestValidationError{
			Issues: issues,
		}
	}

	return nil
}
