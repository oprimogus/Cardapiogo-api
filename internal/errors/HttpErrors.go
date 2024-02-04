package errors

import (
	"net/http"
)

// InternalServerError creates a new error response representing an internal server error (HTTP 500)
func InternalServerError(msg string) *ErrorResponse {
	if msg == "" {
		msg = "We encountered an error while processing your request."
	}
	return New(http.StatusInternalServerError, msg)
}

// NotFound creates a new error response representing a resource-not-found error (HTTP 404)
func NotFound(msg string) *ErrorResponse {
	if msg == "" {
		msg = "The requested resource was not found."
	}
	return New(http.StatusNotFound, msg)
}

// Unauthorized creates a new error response representing an authentication/authorization failure (HTTP 401)
func Unauthorized(msg string) *ErrorResponse {
	if msg == "" {
		msg = "You are not authenticated to perform the requested action."
	}
	return New(http.StatusUnauthorized, msg)
}

// Forbidden creates a new error response representing an authorization failure (HTTP 403)
func Forbidden(msg string) *ErrorResponse {
	if msg == "" {
		msg = "You are not authorized to perform the requested action."
	}
	return New(http.StatusForbidden, msg)
}

// BadRequest creates a new error response representing a bad request (HTTP 400)
func BadRequest(msg string) *ErrorResponse {
	if msg == "" {
		msg = "Your request is in a bad format."
	}
	return New(http.StatusBadRequest, msg)
}

type invalidField struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

func InvalidInput(errs map[string]string) *ErrorResponse {

	details := []invalidField{}
	for i, v := range errs {
		details = append(details, invalidField{
			Field: i,
			Error: v,
		})
	}
	return &ErrorResponse{
		Status:       http.StatusBadRequest,
		ErrorMessage: "There is some problem with the data you submitted.",
		Details:      details,
	}
}
