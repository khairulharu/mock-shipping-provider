package business

// RequestValidationError should be returned when a request validation error occured
type RequestValidationError struct {
	Reason string
}

func (r RequestValidationError) Error() string {
	return r.Reason
}
