package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	validatorutils "github.com/oprimogus/cardapiogo/internal/api/validator"
	logger "github.com/oprimogus/cardapiogo/pkg/log"
	// "github.com/oprimogus/cardapiogo/internal/domain/auth"
	// "github.com/oprimogus/cardapiogo/internal/domain/types"
	"github.com/oprimogus/cardapiogo/internal/domain/auth"
	"github.com/oprimogus/cardapiogo/internal/domain/types"
	"github.com/oprimogus/cardapiogo/internal/domain/user"
	"github.com/oprimogus/cardapiogo/internal/errors"
	"github.com/oprimogus/cardapiogo/internal/services/oauth2"
)

var log = logger.GetLoggerDefault("OAuthController")

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

func (c *AuthController) StartOAuthFlow(ctx *gin.Context) {
	conf := oauth2.NewGoogleOauthConf()

	jwt, err := auth.GenerateJWTOAuthWithClaims(1, string(types.AccountProviderGoogle))
	if err != nil {
		er := errors.InternalServerError(err.Error())
		ctx.JSON(er.Status, er)
		return
	}

	url := conf.AuthCodeURL(jwt)

	ctx.Redirect(http.StatusTemporaryRedirect, url)

}

// func (c *AuthController) SignUpLoginOauthCallback(ctx *gin.Context) {

// 	stateToken := ctx.Query("state")
// 	valid, err := auth.ValidateStateToken(stateToken)
// 	if err != nil || !valid {
// 		er := errors.Unauthorized("")
// 		ctx.JSON(er.Status, er)
// 		return
// 	}

// 	code := ctx.Request.URL.Query().Get("code")
// 	conf := oauth2.NewGoogleOauthConf()
// 	userData, err := oauth2.GetUserData(ctx, conf, code)
// 	if err != nil {
// 		errorResponse, ok := err.(*errors.ErrorResponse)
// 		if !ok {
// 			ctx.JSON(http.StatusInternalServerError, errors.InternalServerError(err.Error()))
// 			return
// 		}
// 		ctx.JSON(errorResponse.Status, err.Error())
// 	}
// 	existUser, err := c.UserService.GetUserByEmail(ctx, userData.Email)
// 	if err != nil {
// 		dbErr, ok := err.(*errors.ErrorResponse)
// 		if !ok {
// 			ctx.JSON(http.StatusInternalServerError, errors.InternalServerError(err.Error()))
// 			return
// 		}
// 		if dbErr.ErrorMessage == errors.NOT_FOUND_RECORD {
// 			err = c.UserService.CreateUserWithOAuth(ctx, user.CreateUserWithOAuthParams{
// 				Email:           userData.Email,
// 				Role:            types.UserRoleConsumer,
// 				AccountProvider: types.AccountProviderGoogle,
// 			})
// 			if err != nil {
// 				dbErr, ok := err.(*errors.ErrorResponse)
// 				if !ok {
// 					ctx.JSON(http.StatusInternalServerError, errors.InternalServerError(err.Error()))
// 					return
// 				}
// 				ctx.JSON(dbErr.Status, dbErr)
// 				return
// 			}
// 			err = c.UserService.Get
// 		}
// 	}

// 	ctx.JSON(http.StatusOK, userData)

// }

func (c *AuthController) Login(ctx *gin.Context) {
	var user user.Login
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	jwt, err := auth.Login(ctx, c.UserService, &user)
	if err != nil {
		dbErr, ok := err.(*errors.ErrorResponse)
		if !ok {
			ctx.JSON(http.StatusInternalServerError, err.Error())
		}
		ctx.JSON(dbErr.Status, dbErr.ErrorMessage)
	}

	httpOnlyCookie := http.Cookie{
		Name:     "token",
		Value:    jwt,
		Expires:  time.Now().Add(time.Hour * time.Duration(auth.TimeExpireInHour)), // Definir um tempo de expiração
		HttpOnly: false,                                                            // Importante para prevenir acesso via JavaScript
		Secure:   true,                                                             // Recomendado para uso apenas em conexões HTTPS
		Path:     "/",                                                              // O cookie é válido para todo o site
	}

	http.SetCookie(ctx.Writer, &httpOnlyCookie)
	ctx.Status(http.StatusOK)
}
