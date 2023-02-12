package presentation

import (
	"encoding/json"
	"errors"
	"mock-shipping-provider/business"
	"mock-shipping-provider/presentation/schema"
	"mock-shipping-provider/primitive"
	"net/http"
)

func (p *Presenter) OrderHandler(w http.ResponseWriter, r *http.Request) {
	// Bind the request body
	var requestBody schema.OrderRequest
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
	orderRequest := business.CreateRequest{
		Provider: providerMap[requestBody.Provider],
		Sender: primitive.Address{
			Name:        requestBody.Sender.Name,
			PhoneNumber: requestBody.Sender.PhoneNumber,
			Address:     requestBody.Sender.Address,
			City:        requestBody.Sender.City,
			State:       requestBody.Sender.State,
			Country:     requestBody.Sender.Country,
			PostalCode:  requestBody.Sender.PostalCode,
			Coordinate:  requestBody.Sender.Coordinate,
		},
		Recipient: primitive.Address{
			Name:        requestBody.Sender.Name,
			PhoneNumber: requestBody.Sender.PhoneNumber,
			Address:     requestBody.Sender.Address,
			City:        requestBody.Sender.City,
			State:       requestBody.Sender.State,
			Country:     requestBody.Sender.Country,
			PostalCode:  requestBody.Sender.PostalCode,
			Coordinate:  requestBody.Sender.Coordinate,
		},
		Dimension: primitive.Dimension{
			Height: requestBody.Dimension.Height,
			Width:  requestBody.Dimension.Width,
			Depth:  requestBody.Dimension.Depth,
		},
		Weight:          requestBody.Weight,
		ItemDescription: requestBody.ItemDescription,
		ItemCategory:    requestBody.ItemCategory,
		Fragile:         requestBody.Fragile,
	}

	// Call the service
	orderResponse, err := p.shippingService.Create(r.Context(), orderRequest)
	// handle error
	if err != nil {
		var requestValidationError *business.RequestValidationError
		if errors.As(err, &requestValidationError) {
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
	}

	// Convert into common order response
	responseBody, err := json.Marshal(schema.OrderResponse{
		StatusCode:           primitive.StatusOrderPlaced,
		StatusDescription:    "order created",
		ReferenceNumber:      orderResponse.ReferenceNumber,
		AirWaybill:           orderResponse.AirWaybill,
		Price:                orderResponse.Price,
		EstimatedHourArrival: orderResponse.Hours,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// if response is success
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(responseBody)
}
