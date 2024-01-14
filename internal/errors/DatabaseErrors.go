package errors

import (
	"errors"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

const (
	NOT_FOUND_RECORD      = "Record not found."
	DUPLICATED_RECORD     = "There is a record with this data."
	FOREIGN_KEY_VIOLATION = "Foreign key violation."
	NULL_VIOLATION        = "Null value not allowed for column."
	VALUE_TOO_LONG        = "Input value too long for column."
	INTERNAL_SERVER_ERROR = "Internal Server Error."
)

func NewDatabaseError(err error) *ErrorResponse {
	if err == nil {
		return nil
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case "23505": // Duplicação de chave única
			return NewErrorResponse(http.StatusConflict, DUPLICATED_RECORD)
		case "23503": // Violação de chave estrangeira
			return NewErrorResponse(http.StatusBadRequest, FOREIGN_KEY_VIOLATION)
		case "23502": // Violação de campo não nulo
			return NewErrorResponse(http.StatusBadRequest, NULL_VIOLATION)
		case "22001": // Dados da string muito longos
			return NewErrorResponse(http.StatusBadRequest, VALUE_TOO_LONG)
		}
	}

	if errors.Is(err, pgx.ErrNoRows) {
		return NewErrorResponse(http.StatusNotFound, NOT_FOUND_RECORD)
	}

	// Caso padrão para outros erros
	return NewErrorResponse(http.StatusInternalServerError, INTERNAL_SERVER_ERROR)
}
