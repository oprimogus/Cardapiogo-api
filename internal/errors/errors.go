package xerrors

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Nerzal/gocloak/v13"

	"github.com/oprimogus/cardapiogo/internal/core/store"
	"github.com/oprimogus/cardapiogo/internal/core/user"
	logger "github.com/oprimogus/cardapiogo/pkg/log"
)

var log = logger.NewLogger("Error Handling")

type ErrorResponse struct {
	Status       int         `json:"-"`
	ErrorMessage string      `json:"error"`
	Details      interface{} `json:"details"`
	Debug        interface{} `json:"debug,omitempty"`
}

func New(status int, message string, details ...interface{}) *ErrorResponse {
	return &ErrorResponse{
		Status:       status,
		ErrorMessage: message,
		Details:      details,
	}
}

func Map(err error) *ErrorResponse {
	log.Error(err.Error())
	if errResp, ok := err.(*ErrorResponse); ok {
		return errResp
	}
	if errResp, ok := err.(*json.UnmarshalTypeError); ok {
		return &ErrorResponse{
			Status:       http.StatusBadRequest,
			ErrorMessage: fmt.Sprintf("Invalid JSON: field %s is not valid for type %s", errResp.Field, errResp.Value),
			Details:      errResp.Struct,
		}
	}
	if errResp, ok := err.(*gocloak.APIError); ok {
		return &ErrorResponse{
			Status:       errResp.Code,
			ErrorMessage: errResp.Message,
			Details:      errResp.Type,
		}
	}

	errResp := mapDatabaseErrors(err)
	if errResp != nil {
		return errResp
	}

	switch err {
	case user.ErrExistUserWithDocument,
		user.ErrExistUserWithEmail,
		user.ErrExistUserWithPhone:
		return &ErrorResponse{
			Status:       http.StatusConflict,
			ErrorMessage: err.Error(),
		}
	case store.ErrClosingTimeBeforeOpeningTime,
		store.ErrOpeningTimeAfterClosingTime:
		return &ErrorResponse{
			Status:       http.StatusBadRequest,
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
