package postgres

import "strings"

type DBErrorType int

const (
    UnknownError DBErrorType = iota
    RecordNotFoundError
    DuplicateRecordError
)

type DBError struct {
    Type    DBErrorType
    Message string
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
			Type: DuplicateRecordError,
			Message: "There is a record with this data.",
			HttpStatus: 409,
		}
    case strings.Contains(err.Error(), "record not found"):
        return &DBError{
			Type: RecordNotFoundError,
			Message: "Record not found.",
			HttpStatus: 404,
		}
    default:
        return &DBError{
			Type: UnknownError,
			Message: "Internal Server error.",
			HttpStatus: 500,
		}
    }
}