package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	validatorutils "github.com/oprimogus/cardapiogo/internal/api/validator"
	"github.com/oprimogus/cardapiogo/internal/core/authentication"
	xerrors "github.com/oprimogus/cardapiogo/internal/errors"
)

// UserController struct
type AuthController struct {
	validator            *validatorutils.Validator
	authenticationModule authentication.AuthenticationModule
}

func NewAuthController(validator *validatorutils.Validator, authRepository authentication.Repository) *AuthController {
	return &AuthController{
		validator:            validator,
		authenticationModule: authentication.NewAuthenticationModule(authRepository),
	}
}

// SignIn godoc
//
//	@Summary		Sign-In with email and password
//	@Description	Authenticate a user using email and password and issue a JWT on successful login.
//	@Tags			Authentication
//	@Accept			json
//	@Produce		json
//	@Param			request	body		authentication.SignInParams	false	"SignInParams"
//	@Success		200		{object}	authentication.JWT
//	@Failure		400		{object}	xerrors.ErrorResponse
//	@Failure		500		{object}	xerrors.ErrorResponse
//	@Failure		502		{object}	xerrors.ErrorResponse
//	@Router			/v1/auth/sign-in [post]
func (c *AuthController) SignIn(ctx *gin.Context) {
	var params authentication.SignInParams
	err := ctx.BindJSON(&params)
	if err != nil {
		xerror := xerrors.Map(err)
		ctx.JSON(xerror.Status, xerror)
		return
	}
	errValidate := c.validator.Validate(params)
	if errValidate != nil {
		xerror := xerrors.Map(errValidate)
		ctx.JSON(xerror.Status, xerror)
		return
	}
	jwt, err := c.authenticationModule.SignIn.Execute(ctx, params.Email, params.Password)
	if err != nil {
		xerror := xerrors.Map(err)
		ctx.JSON(xerror.Status, xerror)
		return
	}

	ctx.JSON(http.StatusOK, jwt)
}

// RefreshUserToken godoc
//
//	@Summary		Refresh token when access token expires
//	@Description	Refresh token when access token expires
//	@Tags			Authentication
//	@Accept			json
//	@Produce		json
//	@Param			request	body		authentication.RefreshParams	true	"RefreshParams"
//	@Success		200		{object}	authentication.JWT
//	@Failure		400		{object}	xerrors.ErrorResponse
//	@Failure		500		{object}	xerrors.ErrorResponse
//	@Failure		502		{object}	xerrors.ErrorResponse
//	@Router			/v1/auth/refresh [post]
func (c *AuthController) RefreshUserToken(ctx *gin.Context) {
	var params authentication.RefreshParams
	err := ctx.BindJSON(&params)
	if err != nil {
		xerror := xerrors.Map(err)
		ctx.JSON(xerror.Status, xerror)
		return
	}
	errValidate := c.validator.Validate(params)
	if errValidate != nil {
		xerror := xerrors.Map(errValidate)
		ctx.JSON(xerror.Status, xerror)
		return
	}
	jwt, err := c.authenticationModule.Refresh.Execute(ctx, params.RefreshToken)
	if err != nil {
		xerror := xerrors.Map(err)
		ctx.JSON(xerror.Status, xerror)
		return
	}

	ctx.JSON(http.StatusOK, jwt)
}
