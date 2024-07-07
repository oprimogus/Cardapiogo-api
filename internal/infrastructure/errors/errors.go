package xerrors

import (
	"net/http"

	"github.com/Nerzal/gocloak/v13"

	"github.com/oprimogus/cardapiogo/internal/application/user"
)

// ErrorResponse is the response that represents an error.
type ErrorResponse struct {
	Status       int         `json:"-"`
	ErrorMessage string      `json:"error"`
	Details      interface{} `json:"details"`
}

func New(status int, message string, details ...interface{}) *ErrorResponse {
	return &ErrorResponse{
		Status:       status,
		ErrorMessage: message,
		Details:      details,
	}
}

func Map(err error) *ErrorResponse {
	if errResp, ok := err.(*ErrorResponse); ok {
		return errResp
	}
	if errResp, ok := err.(*gocloak.APIError); ok {
		return &ErrorResponse{
			Status:       errResp.Code,
			ErrorMessage: errResp.Message,
			Details:      errResp.Type,
		}
	}

	switch err {
	case user.ErrExistUserWithDocument,
		user.ErrExistUserWithEmail,
		user.ErrExistUserWithPhone:
		return &ErrorResponse{
			Status:       http.StatusConflict,
			ErrorMessage: err.Error(),
		}

	default:
		return &ErrorResponse{
			Status:       http.StatusInternalServerError,
			ErrorMessage: err.Error(),
		}
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
