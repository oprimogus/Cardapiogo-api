package xerrors

import (
	"net/http"
)

// InternalServerError creates a new error response representing an internal server error (HTTP 500)
func InternalServerError(msg string, transactionID string) *ErrorResponse {
	if msg == "" {
		msg = "We encountered an error while processing your request."
	}
	return New(http.StatusInternalServerError, msg, transactionID)
}

func ConflictError(msg string, transactionID string) *ErrorResponse {
	if msg == "" {
		msg = "We encountered an conflict error while processing your request."
	}
	return New(http.StatusConflict, msg, transactionID)
}

// NotFound creates a new error response representing a resource-not-found error (HTTP 404)
func NotFound(msg string, transactionID string) *ErrorResponse {
	if msg == "" {
		msg = "The requested resource was not found."
	}
	return New(http.StatusNotFound, msg, transactionID)
}

// Unauthorized creates a new error response representing an authentication/authorization failure (HTTP 401)
func Unauthorized(msg string, transactionID string) *ErrorResponse {
	if msg == "" {
		msg = "You are not authenticated to perform the requested action."
	}
	return New(http.StatusUnauthorized, msg, transactionID)
}

// Forbidden creates a new error response representing an authorization failure (HTTP 403)
func Forbidden(msg string, transactionID string) *ErrorResponse {
	if msg == "" {
		msg = "You are not authorized to perform the requested action."
	}
	return New(http.StatusForbidden, msg, transactionID)
}

// BadRequest creates a new error response representing a bad request (HTTP 400)
func BadRequest(msg string, transactionID string) *ErrorResponse {
	if msg == "" {
		msg = "Your request is in a bad format."
	}
	return New(http.StatusBadRequest, msg, transactionID)
}

type invalidField struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

func InvalidInput(errs map[string]string, transactionID string) *ErrorResponse {

	details := []invalidField{}
	for i, v := range errs {
		details = append(details, invalidField{
			Field: i,
			Error: v,
		})
	}
	return &ErrorResponse{
		Status:        http.StatusBadRequest,
		ErrorMessage:  "There is some problem with the data you submitted.",
		Details:       details,
		TransactionID: transactionID,
	}
}
