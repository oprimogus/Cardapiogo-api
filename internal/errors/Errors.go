package errors

// ErrorResponse is the response that represents an error.
type ErrorResponse struct {
	Status       int         `json:"-"`
	ErrorMessage string      `json:"error"`
	Details      interface{} `json:"details,omitempty"`
}

func New(status int, message string, details ...interface{}) *ErrorResponse {
	return &ErrorResponse{
		Status:       status,
		ErrorMessage: message,
		Details:      details,
	}
}

// Error is required by the error interface.
func (e *ErrorResponse) Error() string {
	return e.ErrorMessage
}

// StatusCode is required by routing.HTTPError interface.
func (e *ErrorResponse) StatusCode() int {
	return e.Status
}
