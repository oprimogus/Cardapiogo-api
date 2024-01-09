package errordatabase

import (
	"errors"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type DatabaseErrorType int

const (
	ForeignKeyViolationError DatabaseErrorType = iota + 4
	SyntaxError
	UndefinedTableError
	UndefinedColumnError
	StringDataRightTruncationError
	DuplicateRecordError
	InvalidInputError
	RecordNotFoundError
	UnknownError
	NullViolationError
)

const (
	NOT_FOUND_RECORD      = "Record not found."
	DUPLICATED_RECORD     = "There is a record with this data."
	FOREIGN_KEY_VIOLATION = "Foreign key violation."
	NULL_VIOLATION        = "Null value not allowed for column."
	VALUE_TOO_LONG        = "Input value too long for column."
)

type DatabaseError struct {
	Type       DatabaseErrorType
	Message    string
	HttpStatus int
}

func NewDBError(errorType DatabaseErrorType, message string, httpStatus int) *DatabaseError {
	return &DatabaseError{
		Type:       errorType,
		Message:    message,
		HttpStatus: httpStatus,
	}
}

func (e *DatabaseError) Error() string {
	return e.Message
}

func MapDBError(err error) *DatabaseError {
	if err == nil {
		return nil
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case "23505": // Duplicação de chave única
			return NewDBError(DuplicateRecordError, DUPLICATED_RECORD, http.StatusConflict)
		case "23503": // Violação de chave estrangeira
			return NewDBError(ForeignKeyViolationError, FOREIGN_KEY_VIOLATION, http.StatusBadRequest)
		case "23502": // Violação de campo não nulo
			return NewDBError(NullViolationError, NULL_VIOLATION, http.StatusBadRequest)
		case "22001": // Dados da string muito longos
			return NewDBError(StringDataRightTruncationError, VALUE_TOO_LONG, http.StatusBadRequest)
		}
	}

	if errors.Is(err, pgx.ErrNoRows) {
		return NewDBError(RecordNotFoundError, NOT_FOUND_RECORD, http.StatusNotFound)
	}

	// Caso padrão para outros erros
	return NewDBError(UnknownError, "Unknown database error", http.StatusInternalServerError)
}
