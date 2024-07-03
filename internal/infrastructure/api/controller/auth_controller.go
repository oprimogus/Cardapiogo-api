package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/oprimogus/cardapiogo/internal/application/authentication"
	"github.com/oprimogus/cardapiogo/internal/domain/repository"
	validatorutils "github.com/oprimogus/cardapiogo/internal/infrastructure/api/validator"
	logger "github.com/oprimogus/cardapiogo/pkg/log"
)

var log = logger.NewLogger("AuthController")

// UserController struct
type AuthController struct {
	validator *validatorutils.Validator
	signIn    authentication.SignIn
}

func NewAuthController(
	validator *validatorutils.Validator,
	factory repository.Factory,
) *AuthController {
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
//	@Failure		400		{object}	errors.ErrorResponse
//	@Failure		500		{object}	errors.ErrorResponse
//	@Failure		502		{object}	errors.ErrorResponse
//	@Router			/v1/auth/login [post]
func (c *AuthController) SignIn(ctx *gin.Context) {
	var loginParams authentication.SignInParams
	err := ctx.BindJSON(&loginParams)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	jwt, err := c.signIn.Execute(ctx, loginParams.Email, loginParams.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, jwt)
}

// TestRoute godoc
//
//	@Summary		Route for test authentication middleware
//	@Description	Route for test authentication middleware
//	@Tags			Authentication
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	entity.JWT
//	@Failure		400	{object}	errors.ErrorResponse
//	@Failure		500	{object}	errors.ErrorResponse
//	@Failure		502	{object}	errors.ErrorResponse
//	@Security		Bearer Token
//	@Router			/v1/auth/test/authentication [get]
func (c *AuthController) ProtectedRoute(ctx *gin.Context) {
	userID := ctx.GetString("userID")
	log.Infof(userID)
	ctx.JSON(http.StatusOK, gin.H{"message": "Alright, you access the protected endpoint!"})
}