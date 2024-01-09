package errors

import (
	"net/http"
)

// ErrorResponse is the response that represents an error.
type ErrorResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

func NewErrorResponse(status int, message string, details ...interface{}) *ErrorResponse {
	return &ErrorResponse{
		Status: status,
		Message: message,
		Details: details,
	}
}

// Error is required by the error interface.
func (e ErrorResponse) Error() string {
	return e.Message
}

// StatusCode is required by routing.HTTPError interface.
func (e ErrorResponse) StatusCode() int {
	return e.Status
}

// InternalServerError creates a new error response representing an internal server error (HTTP 500)
func InternalServerError(msg string) *ErrorResponse {
	if msg == "" {
		msg = "We encountered an error while processing your request."
	}
	return NewErrorResponse(http.StatusInternalServerError, msg)
}

// NotFound creates a new error response representing a resource-not-found error (HTTP 404)
func NotFound(msg string) *ErrorResponse {
	if msg == "" {
		msg = "The requested resource was not found."
	}
	return NewErrorResponse(http.StatusNotFound, msg)
}

// Unauthorized creates a new error response representing an authentication/authorization failure (HTTP 401)
func Unauthorized(msg string) *ErrorResponse {
	if msg == "" {
		msg = "You are not authenticated to perform the requested action."
	}
	return NewErrorResponse(http.StatusUnauthorized, msg)
}

// Forbidden creates a new error response representing an authorization failure (HTTP 403)
func Forbidden(msg string) *ErrorResponse {
	if msg == "" {
		msg = "You are not authorized to perform the requested action."
	}
	return NewErrorResponse(http.StatusForbidden, msg)
}

// BadRequest creates a new error response representing a bad request (HTTP 400)
func BadRequest(msg string) *ErrorResponse {
	if msg == "" {
		msg = "Your request is in a bad format."
	}
	return NewErrorResponse(http.StatusBadRequest, msg)
}

type invalidField struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

func InvalidInput(errs map[string]string) *ErrorResponse{
	details := make([]invalidField, len(errs))

	for i, v := range errs {
		details = append(details, invalidField{
			Field: i,
			Error: v,
		})
	}
	return &ErrorResponse{
		Status: http.StatusBadRequest,
		Message: "There is some problem with the data you submitted.",
		Details: details,
	}
}