package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	validatorutils "github.com/oprimogus/cardapiogo/internal/api/validator"
	"github.com/oprimogus/cardapiogo/internal/domain/profile"
	"github.com/oprimogus/cardapiogo/internal/domain/user"
	"github.com/oprimogus/cardapiogo/internal/errors"
)

type ProfileController struct {
	service   *profile.Service
	validator *validatorutils.Validator
}

func NewProfileController(repository profile.Repository, userRepository user.Repository, validator *validatorutils.Validator) *ProfileController {
	return &ProfileController{service: profile.NewService(repository, userRepository), validator: validator}
}

func (c *ProfileController) CreateProfileHandler(ctx *gin.Context) {
	var createProfileParams profile.CreateProfileParams
	err := ctx.BindJSON(&createProfileParams)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errors.BadRequest(err.Error()))
		return
	}

	validationErr := c.validator.Validate(createProfileParams)
	if validationErr != nil {
		ctx.JSON(http.StatusBadRequest, validationErr)
	}

	userID := ctx.GetString("userID")

	err = c.service.CreateProfile(ctx, userID, createProfileParams)
	validateErrorResponse(ctx, err)
	ctx.Status(http.StatusOK)
}
