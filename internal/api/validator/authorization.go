package validatorutils

import (
	"github.com/gin-gonic/gin"

	"github.com/oprimogus/cardapiogo/internal/errors"
)

func IsSameUser(ctx *gin.Context, userID string) error {
	extractID, exists := ctx.Get("userID")
	if !exists {
		return errors.Forbidden("The user ID must exist in the context.")
	}
	if extractID != userID {
		return errors.Forbidden("You are not the account owner.")
	}
	return nil
}
