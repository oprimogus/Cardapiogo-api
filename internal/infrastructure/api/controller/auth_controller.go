package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	validatorutils "github.com/oprimogus/cardapiogo/internal/infrastructure/api/validator"
	"github.com/oprimogus/cardapiogo/internal/infrastructure/services/authentication/keycloak"
)

// UserController struct
type AuthController struct {
	Validator       *validatorutils.Validator
	KeycloakService *keycloak.KeycloakService
}

func NewAuthController(validator *validatorutils.Validator) *AuthController {
	return &AuthController{
		Validator: validator,
	}
}

// SignInOrSignUpKeycloakOauth godoc
//
//	@Summary		Start OAuth2 authentication
//	@Description	Start OAuth2 authentication with Keycloak. The flow continues in the callback route.
//	@Tags			Authentication
//	@Accept			json
//	@Produce		json
//	@Success		302
//	@Failure		400	{object}	errors.ErrorResponse
//	@Failure		500	{object}	errors.ErrorResponse
//	@Failure		502	{object}	errors.ErrorResponse
//	@Router			/v1/auth/oauth [get]
func (c *AuthController) SignInOrSignUpKeycloakOauth(ctx *gin.Context) {
	// conf, err := keycloak.NewOauthConf(ctx)
	// if err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, err)
	// 	return
	// }
	jwt, err := keycloak.GenerateJWTForValidation()
	validateErrorResponse(ctx, err)

	url := c.KeycloakService.Config.AuthCodeURL(jwt)

	ctx.Redirect(http.StatusFound, url)
}

// SignInOrSignUpKeycloakOauthCallback godoc
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
//	@Router			/v1/auth/oauth/callback [get]
func (c *AuthController) SignInOrSignUpKeycloakOauthCallback(ctx *gin.Context) {
	stateToken := ctx.Query("state")
	valid, err := keycloak.ValidateStateToken(stateToken)
	if err != nil || !valid {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	code := ctx.Request.URL.Query().Get("code")
	conf := c.KeycloakService.Config
	oauth2Token, err := conf.Exchange(ctx, code)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, oauth2Token)
}

// Login2 godoc
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
//	@Router			/v1/auth/login2 [post]
func (c *AuthController) SignUp(ctx *gin.Context) {
	var loginParams struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	err := ctx.BindJSON(&loginParams)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	login, err := c.KeycloakService.SignUp(ctx, loginParams.Email, loginParams.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, login)

}
