package schema

type Estimation struct {
	Provider             string `json:"provider"`
	EstimatedPrice       int64  `json:"estimated_price"`
	EstimatedHourArrival uint64 `json:"estimated_hour_arrival"`
}

type EstimateResponse struct {
	Estimation []Estimation `json:"estimation"`
}
