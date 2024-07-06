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

func NewAuthController(validator *validatorutils.Validator, factory repository.Factory) *AuthController {
	return &AuthController{
		validator: validator,
		signIn:    authentication.NewSignIn(factory.NewAuthenticationRepository()),
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
//	@Success		200		{object}	entity.JWT
//	@Failure		400		{object}	xerrors.ErrorResponse
//	@Failure		500		{object}	xerrors.ErrorResponse
//	@Failure		502		{object}	xerrors.ErrorResponse
//	@Router			/v1/auth/sign-in [post]
func (c *AuthController) SignIn(ctx *gin.Context) {
	var loginParams authentication.SignInParams
	err := ctx.BindJSON(&loginParams)
	if err != nil {
		xerror := xerrors.Map(err)
		ctx.JSON(xerror.Status, xerror)
		return
	}
	errValidate := c.validator.Validate(loginParams)
	if errValidate != nil {
		xerror := xerrors.Map(errValidate)
		ctx.JSON(xerror.Status, xerror)
		return
	}
	jwt, err := c.signIn.Execute(ctx, loginParams.Email, loginParams.Password)
	if err != nil {
		xerror := xerrors.Map(err)
		ctx.JSON(xerror.Status, xerror)
		return
	}

	ctx.JSON(http.StatusOK, jwt)
}

// TestAuthenticationRoute godoc
//
//	@Summary		Route for test authentication middleware
//	@Description	Route for test authentication middleware
//	@Tags			Authentication
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	entity.JWT
//	@Failure		400	{object}	xerrors.ErrorResponse
//	@Failure		500	{object}	xerrors.ErrorResponse
//	@Failure		502	{object}	xerrors.ErrorResponse
//	@Security		Bearer Token
//	@Router			/v1/auth/test/authentication [get]
func (c *AuthController) ProtectedRoute(ctx *gin.Context) {
	userID := ctx.GetString("userID")
	log.Infof(userID)
	ctx.JSON(http.StatusOK, gin.H{"message": "Alright, you access the protected endpoint!"})
}

// TestAuthorizationRoute godoc
//
//	@Summary		Route for test authorization middleware
//	@Description	Route for test authorization middleware
//	@Tags			Authentication
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	entity.JWT
//	@Failure		400	{object}	xerrors.ErrorResponse
//	@Failure		500	{object}	xerrors.ErrorResponse
//	@Failure		502	{object}	xerrors.ErrorResponse
//	@Security		Bearer Token
//	@Router			/v1/auth/test/authorization [get]
func (c *AuthController) ProtectedRouteForRoles(ctx *gin.Context) {
	userRoles := ctx.GetStringSlice("userRoles")
	log.Infof("%v", userRoles)
	ctx.JSON(http.StatusOK, gin.H{"message": "Alright, you access the role protected endpoint!"})
}
