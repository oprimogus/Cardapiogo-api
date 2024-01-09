package httpResponses

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/oprimogus/cardapiogo/internal/errors"
)

type httpErrorResponse struct {
	Error   string
	Details interface{}
}

func newErrorResponse(errs *errors.ErrorResponse) *httpErrorResponse {
	return &httpErrorResponse{
		Error: errs.Error(),
	}
}

func newErrorResponseWithMessage(errs error) *httpErrorResponse {
	return &httpErrorResponse{
		Error: errs.Error(),
	}
}

func newErrorResponseWithDetails(errs *errors.ErrorResponse) *httpErrorResponse {
	return &httpErrorResponse{
		Error: errs.Error(),
		Details: errs.Details,
	}
}

// InternalServerErrorResponse creates a new error response representing an internal server error (HTTP 500)
func InternalServerErrorResponse(ctx *gin.Context, err *errors.ErrorResponse) {
	ctx.JSON(err.StatusCode(), newErrorResponse(err))
}

// InternalServerError creates a new error response representing an internal server error (HTTP 500)
func NotFoundResponse(ctx *gin.Context, err *errors.ErrorResponse) {
	ctx.JSON(err.StatusCode(), newErrorResponse(err))
}

// Unauthorized creates a new error response representing an authentication/authorization failure (HTTP 401)
func Unauthorized(ctx *gin.Context, err *errors.ErrorResponse) {
	ctx.JSON(err.StatusCode(), newErrorResponse(err))
}

// Forbidden creates a new error response representing an authorization failure (HTTP 403)
func Forbidden(ctx *gin.Context, err *errors.ErrorResponse) {
	ctx.JSON(err.StatusCode(), newErrorResponse(err))
}

// BadRequest creates a new error response representing a bad request (HTTP 400)
func BadRequest(ctx *gin.Context, err *errors.ErrorResponse){
	ctx.JSON(err.StatusCode(), newErrorResponseWithDetails(err))
}

// BadRequest creates a new error response representing a bad request (HTTP 400)
func BadRequestMessage(ctx *gin.Context, err error){
	ctx.JSON(http.StatusBadRequest, newErrorResponseWithMessage(err))
}

