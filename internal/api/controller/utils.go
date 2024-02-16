package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/oprimogus/cardapiogo/internal/errors"
)

// validateErrorResponse validate if error is an ErrorResponse and return a http response
func validateErrorResponse(ctx *gin.Context, err error) {
	if err != nil {
		errorResponse, ok := err.(*errors.ErrorResponse)
		if ok {
			ctx.JSON(errorResponse.Status, errorResponse)
			return
		}
		ctx.JSON(http.StatusInternalServerError, errors.New(http.StatusInternalServerError, err.Error()))
		return
	}

}
