package presentation

import (
	"encoding/json"
	"errors"
	"mock-shipping-provider/business"
	"mock-shipping-provider/presentation/schema"
	"net/http"
)

func (p *Presenter) EstimateHandler(w http.ResponseWriter, r *http.Request) {
	// bind the request body
	var requestBody schema.EstimateRequest
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		responseBody, e := json.Marshal(schema.Error{
			StatusCode:    http.StatusBadRequest,
			StatusMessage: "mailformed JSON",
		})

		if e != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(responseBody)
		return
	}

	// Convert into business schema
	estimateRequest := business.EstimateRequest{
		Sender:    requestBody.Sender,
		Recipient: requestBody.Recipient,
		Dimension: requestBody.Dimension,
		Weight:    requestBody.Weight,
	}

	//  Call appropiate service
	estimateResponse, err := p.shippingService.Estimate(r.Context(), estimateRequest)

	// handle error
	if err != nil {
		// if error is not serviceable or
		// request validation error
		var requestValidatonError *business.RequestValidationError
		if errors.Is(err, business.ErrNotServiceable) || errors.As(err, &requestValidatonError) {
			responseBody, e := json.Marshal(schema.Error{
				StatusCode:    http.StatusBadRequest,
				StatusMessage: err.Error(),
			})

			if e != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write(responseBody)
			return
		}

		// TODO: send to logger

		// internal server error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Convert into presentation response
	responseSchema := schema.EstimateResponse{}
	for _, val := range estimateResponse {
		responseSchema.Estimation = append(responseSchema.Estimation, schema.Estimation{
			Provider:             val.Provider.String(),
			EstimatedPrice:       val.Price,
			EstimatedHourArrival: val.Hours,
		})
	}

	responseBody, err := json.Marshal(responseSchema)
	if err != nil {
		// TODO: send to logger

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "applicatino/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseBody)
}
