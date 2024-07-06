package validatorutils

import (
	"github.com/gin-gonic/gin"

	"github.com/oprimogus/cardapiogo/internal/infrastructure/errors"
)

func IsSameUser(ctx *gin.Context, userID string) error {
	extractID, exists := ctx.Get("userID")
	if !exists {
		return xerrors.Forbidden("The user ID must exist in the context.")
	}
	if extractID != userID {
		return xerrors.Forbidden("You are not the account owner.")
	}
	return nil
}
