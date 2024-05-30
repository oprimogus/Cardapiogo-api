package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	validatorutils "github.com/oprimogus/cardapiogo/internal/api/validator"
	"github.com/oprimogus/cardapiogo/internal/domain/user"
	"github.com/oprimogus/cardapiogo/internal/errors"
	"github.com/oprimogus/cardapiogo/internal/services/auth"
	"github.com/oprimogus/cardapiogo/internal/services/oauth2"
)

// UserController struct
type AuthController struct {
	UserService *user.Service
	Validator   *validatorutils.Validator
}

func NewAuthController(
	repository user.Repository,
	validator *validatorutils.Validator,
) *AuthController {
	return &AuthController{
		UserService: user.NewService(repository),
		Validator:   validator,
	}
}

// StartGoogleOAuthFlow godoc
//
//	@Summary		Start OAuth2 authentication
//	@Description	Start OAuth2 authentication with Google. The flow continues in the callback route.
//	@Tags			Authentication
//	@Accept			json
//	@Produce		json
//	@Success		307
//	@Failure		400	{object}	errors.ErrorResponse
//	@Failure		500	{object}	errors.ErrorResponse
//	@Failure		502	{object}	errors.ErrorResponse
//	@Router			/v1/auth/google [get]
func (c *AuthController) StartGoogleOAuthFlow(ctx *gin.Context) {
	conf := oauth2.NewGoogleOauthConf()

	jwt, err := auth.GenerateJWTForValidation()
	validateErrorResponse(ctx, err)

	url := conf.AuthCodeURL(jwt)

	ctx.Redirect(http.StatusTemporaryRedirect, url)
}

// SignUpLoginGoogleOauthCallback godoc
//
//	@Summary		OAuth2 login callback
//	@Description	Handle OAuth2 login callback and issue a JWT on successful authentication.
//	@Tags			Authentication
//	@Accept			json
//	@Produce		json
//	@Success		307
//	@Failure		400	{object}	errors.ErrorResponse
//	@Failure		500	{object}	errors.ErrorResponse
//	@Failure		502	{object}	errors.ErrorResponse
//	@Router			/v1/auth/google/callback [get]
func (c *AuthController) SignUpLoginGoogleOauthCallback(ctx *gin.Context) {
	stateToken := ctx.Query("state")
	valid, err := auth.ValidateStateToken(stateToken)
	if err != nil || !valid {
		er := errors.Unauthorized(err.Error())
		ctx.JSON(er.Status, er)
		return
	}

	code := ctx.Request.URL.Query().Get("code")
	conf := oauth2.NewGoogleOauthConf()
	userData, err := oauth2.GetGoogleUserData(ctx, conf, code)
	validateErrorResponse(ctx, err)

	jwt, err := auth.LoginWithOauth(ctx, c.UserService, userData)
	validateErrorResponse(ctx, err)

	ctx.JSON(
		http.StatusOK,
		gin.H{
			"token":    jwt,
			"redirect": "https://weather-app-angular-jeby7qw78-oprimogus.vercel.app/weather",
		},
	)
}

// Login godoc
//
//	@Summary		User login with email and password
//	@Description	Authenticate a user using email and password and issue a JWT on successful login.
//	@Tags			Authentication
//	@Accept			json
//	@Produce		json
//	@Param			request	body	user.Login	true	"Login credentials"
//	@Success		200
//	@Failure		400	{object}	errors.ErrorResponse
//	@Failure		500	{object}	errors.ErrorResponse
//	@Failure		502	{object}	errors.ErrorResponse
//	@Router			/v1/login [post]
func (c *AuthController) Login(ctx *gin.Context) {
	var user user.Login
	err := ctx.BindJSON(&user)
	validateErrorResponse(ctx, err)

	jwt, err := auth.Login(ctx, c.UserService, &user)
	if err != nil {
		validateErrorResponse(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": jwt})
}
