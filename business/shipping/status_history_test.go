package shipping_test

import (
	"mock-shipping-provider/business"
	"mock-shipping-provider/business/shipping"
	"testing"
)

func TestValidateStatusHistoryRequest(t *testing.T) {

	testCases := []struct {
		whenNil    business.StatusRequest
		whenDefine business.StatusRequest
	}{
		{whenDefine: business.StatusRequest{
			ReferenceNumber: "referenceNumber",
			AirWaybill:      "airWaybill",
		}},
	}

	for _, testCase := range testCases {
		t.Run("when request nil", func(t *testing.T) {
			if err := shipping.ValidateStatusHistoryRequest(testCase.whenNil); err == nil {
				t.Error("should get errorr but nil")
			}
		})
	}

	for _, testCase := range testCases {
		t.Run("request are difine", func(t *testing.T) {
			if err := shipping.ValidateStatusHistoryRequest(testCase.whenDefine); err != nil {
				t.Errorf("must not get erorr but get instead: %v", err)
			}
		})
	}

}
