package controller

import (
	"net/http"
	"time"

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

func NewAuthController(repository user.Repository, validator *validatorutils.Validator) *AuthController {
	return &AuthController{
		UserService: user.NewService(repository),
		Validator:   validator,
	}
}

// StartGoogleOAuthFlow godoc
// @Summary Inicia fluxo de OAuth2
// @Description Inicia fluxo de OAuth2
// @Tags Auth
// @Accept  json
// @Produce  json
// @Success 307
// @Failure 400  {object} errors.ErrorResponse
// @Failure 500  {object} errors.ErrorResponse
// @Failure 502  {object} errors.ErrorResponse
// @Router /auth/google [get]
func (c *AuthController) StartGoogleOAuthFlow(ctx *gin.Context) {
	conf := oauth2.NewGoogleOauthConf()

	jwt, err := auth.GenerateJWTForValidation()
	validateErrorResponse(ctx, err)

	url := conf.AuthCodeURL(jwt)

	ctx.Redirect(http.StatusTemporaryRedirect, url)

}

// SignUpLoginGoogleOauthCallback godoc
// @Summary Callback de login via OAuth2
// @Description Callback de login via OAuth2
// @Tags Auth
// @Accept  json
// @Produce  json
// @Success 307
// @Failure 400  {object} errors.ErrorResponse
// @Failure 500  {object} errors.ErrorResponse
// @Failure 502  {object} errors.ErrorResponse
// @Router /auth/google/callback [get]
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

	httpOnlyCookie := http.Cookie{
		Name:     "token",
		Value:    jwt,
		Expires:  time.Now().Add(time.Hour * time.Duration(auth.TimeExpireInHour)),
		HttpOnly: false,
		Secure:   true,
		Path:     "/",
		SameSite: http.SameSiteStrictMode,
	}
	http.SetCookie(ctx.Writer, &httpOnlyCookie)
	ctx.Redirect(http.StatusMovedPermanently, "https://weather-app-angular-jeby7qw78-oprimogus.vercel.app/weather")
}

// Login godoc
// @Summary Login de usuário com email e senha
// @Description Login de usuário com email e senha
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param   request body user.Login true "Login"
// @Success 200
// @Failure 400  {object} errors.ErrorResponse
// @Failure 500  {object} errors.ErrorResponse
// @Failure 502  {object} errors.ErrorResponse
// @Router /login [post]
func (c *AuthController) Login(ctx *gin.Context) {
	var user user.Login
	err := ctx.BindJSON(&user)
	validateErrorResponse(ctx, err)

	jwt, err := auth.Login(ctx, c.UserService, &user)
	validateErrorResponse(ctx, err)

	httpOnlyCookie := http.Cookie{
		Name:     "token",
		Value:    jwt,
		Expires:  time.Now().Add(time.Hour * time.Duration(auth.TimeExpireInHour)),
		HttpOnly: false,
		Secure:   true,
		Path:     "/",
		SameSite: http.SameSiteStrictMode,
	}

	http.SetCookie(ctx.Writer, &httpOnlyCookie)
	ctx.Status(http.StatusOK)
}
