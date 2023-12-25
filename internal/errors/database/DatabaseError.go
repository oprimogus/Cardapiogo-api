package errordatabase

import (
	"net/http"
	"strings"
)

type DBErrorType int

const (
	UnknownError DBErrorType = iota
	RecordNotFoundError
	DuplicateRecordError
	InvalidInputError
)

type DBError struct {
	Type       DBErrorType
	Message    string
	HttpStatus int
}

func (e *DBError) Error() string {
	return e.Message
}

func MapDBError(err error) *DBError {
	if err == nil {
		return nil
	}
	switch {
	case strings.Contains(err.Error(), "duplicate key value"):
		return &DBError{
			Type:       DuplicateRecordError,
			Message:    "There is a record with this data.",
			HttpStatus: http.StatusConflict,
		}
	case strings.Contains(err.Error(), "record not found"):
		return &DBError{
			Type:       RecordNotFoundError,
			Message:    "Record not found.",
			HttpStatus: http.StatusNotFound,
		}
	case strings.Contains(err.Error(), "value too long"):
		return &DBError{
			Type:       InvalidInputError,
			Message:    "Input value too long for column.",
			HttpStatus: http.StatusBadRequest,
		}
	default:
		return nil
	}
}
