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
	INVALID_VALUE         = "Invalid value for field."
	UNKNOWN_ERROR         = "Unknown error."
)

func NewDatabaseError(err error) *ErrorResponse {
	if err == nil {
		return nil
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case "23505": // Duplicação de chave única
			return New(http.StatusConflict, DUPLICATED_RECORD, pgErr.Message)
		case "23503": // Violação de chave estrangeira
			return New(http.StatusBadRequest, FOREIGN_KEY_VIOLATION, pgErr.Message)
		case "23502": // Violação de campo não nulo
			return New(http.StatusBadRequest, NULL_VIOLATION, pgErr.Message)
		case "22001": // Dados da string muito longos
			return New(http.StatusBadRequest, VALUE_TOO_LONG, pgErr.Message)
		case "22P02": // Inserção de dado inválido
			return New(http.StatusBadRequest, INVALID_VALUE, pgErr.Message)
		default:
			return New(http.StatusBadRequest, UNKNOWN_ERROR, pgErr.Message)
		}
	}

	if errors.Is(err, pgx.ErrNoRows) {
		return New(http.StatusNotFound, NOT_FOUND_RECORD)
	}

	// Caso padrão para outros erros
	return New(http.StatusInternalServerError, INTERNAL_SERVER_ERROR)
}
