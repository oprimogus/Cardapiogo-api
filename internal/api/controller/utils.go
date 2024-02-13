package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/oprimogus/cardapiogo/internal/errors"
)

// returnError validate if error is an ErrorResponse and return a http response
func returnError(ctx *gin.Context, err error) {
	if err != nil {
		errorResponse, ok := err.(*errors.ErrorResponse)
		if ok && errorResponse != nil {
			ctx.JSON(errorResponse.Status, errorResponse)
			return
		}
		ctx.JSON(http.StatusInternalServerError, errors.New(http.StatusInternalServerError, err.Error()))
		return
	}

}
