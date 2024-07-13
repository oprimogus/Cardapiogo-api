package xerrors

import (
	"errors"
	"net/http"
	"os"
	"strings"
	"unicode"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

var environment = os.Getenv("GIN_MODE")

const (
	NOT_FOUND_RECORD      = "Record not found."
	DUPLICATED_RECORD     = "There is a record with this data."
	FOREIGN_KEY_VIOLATION = "Foreign key violation."
	NULL_VIOLATION        = "Null value not allowed for column."
	VALUE_TOO_LONG        = "Input value too long for column."
	INTERNAL_SERVER_ERROR = "Internal Server Error."
	INVALID_VALUE         = "Invalid value for field."
	UNKNOWN_ERROR         = "UnknoThre wn error."
	// DUPLICATED_RECORD     = "Already exist a resource with this field registered"
)

type fieldError struct {
	Field   string `json:"field"`
	Input   string `json:"input"`
	Message string `json:"message"`
	Debug interface{} `json:"debug,omitempty"`
}

func mapDatabaseErrors(err error) *ErrorResponse {
	if err == nil {
		return nil
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case "23505":
			return handleUniqueViolation(pgErr)
		case "23503": // Violação de chave estrangeira
			return New(http.StatusBadRequest, FOREIGN_KEY_VIOLATION, pgErr.Message)
		case "23502": // Violação de campo não nulo
			return New(http.StatusBadRequest, NULL_VIOLATION, pgErr.Message)
		case "22001": // Dados da string muito longos
			return New(http.StatusBadRequest, VALUE_TOO_LONG, pgErr.Message)
		case "22P02": // Inserção de dado inválido
			return New(http.StatusBadRequest, INVALID_VALUE, pgErr.Message)
		default:
			return New(http.StatusInternalServerError, UNKNOWN_ERROR, pgErr.Message)
		}
	}

	if errors.Is(err, pgx.ErrNoRows) {
		return New(http.StatusNotFound, NOT_FOUND_RECORD)
	}

	return nil
}

func snakeToCamelCase(s string) string {
	words := strings.Split(s, "_")
	for i := 1; i < len(words); i++ {
		firstChar := []rune(words[i])[0]
		words[i] = string(unicode.ToUpper(firstChar)) + words[i][1:]
	}
	return strings.Join(words, "")

}

func handleUniqueViolation(pgErr *pgconn.PgError) *ErrorResponse {
	startField := strings.Index(pgErr.Detail, "(") + 1
	endField := strings.Index(pgErr.Detail, ")=")
	field := snakeToCamelCase(pgErr.Detail[startField:endField])

	startValue := strings.Index(pgErr.Detail, "=(") + 2
	endValue := strings.LastIndex(pgErr.Detail, ")")
	value := pgErr.Detail[startValue:endValue]

	description := "This value is already in use"

	fieldError := fieldError{
		Field:   field,
		Input:   value,
		Message: description,
	}
	if environment != "release" {
		fieldError.Debug = pgErr
	}
	return New(http.StatusConflict, DUPLICATED_RECORD, fieldError)
}
