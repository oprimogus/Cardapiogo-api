package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/oprimogus/cardapiogo/internal/application/authentication"
	"github.com/oprimogus/cardapiogo/internal/domain/repository"
	validatorutils "github.com/oprimogus/cardapiogo/internal/infrastructure/api/validator"
	xerrors "github.com/oprimogus/cardapiogo/internal/infrastructure/errors"
)

// UserController struct
type AuthController struct {
	validator *validatorutils.Validator
	signIn    authentication.SignIn
}

func NewAuthController(validator *validatorutils.Validator, authRepository repository.AuthenticationRepository) *AuthController {
	return &AuthController{
		validator: validator,
		signIn:    authentication.NewSignIn(authRepository),
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
//	@Success		200		{object}	object.JWT
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
	jwt, err := c.signIn.Execute(ctx, params.Email, params.Password)
	if err != nil {
		xerror := xerrors.Map(err)
		ctx.JSON(xerror.Status, xerror)
		return
	}

	ctx.JSON(http.StatusOK, jwt)
}
