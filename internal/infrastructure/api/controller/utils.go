package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	validatorutils "github.com/oprimogus/cardapiogo/internal/infrastructure/api/validator"
	"github.com/oprimogus/cardapiogo/internal/infrastructure/errors"
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

func validateIsSameUser(ctx *gin.Context, id string) {
	err := validatorutils.IsSameUser(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusForbidden, err)
		return
	}
}
