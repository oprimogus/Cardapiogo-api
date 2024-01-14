package controller

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"

	"github.com/gin-gonic/gin"

	validatorutils "github.com/oprimogus/cardapiogo/internal/api/validator"
	"github.com/oprimogus/cardapiogo/internal/domain/user"
	"github.com/oprimogus/cardapiogo/internal/errors"
	"github.com/oprimogus/cardapiogo/internal/services/oauth2"
	logger "github.com/oprimogus/cardapiogo/pkg/log"
)

// UserController struct
type Oauth2Controller struct {
	UserService *user.Service
	Validator   *validatorutils.Validator
}

func NewOauth2Controller(repository user.Repository, validator *validatorutils.Validator) *Oauth2Controller {
	return &Oauth2Controller{
		UserService: user.NewService(repository),
		Validator:   validator,
	}
}

func (C *Oauth2Controller) StartOAuthFlow(ctx *gin.Context) {
	conf := oauth2.NewGoogleOauthConf()

	url := conf.AuthCodeURL(generateRandomState())

	ctx.Redirect(http.StatusTemporaryRedirect, url)

}

func (C *Oauth2Controller) SignUpLoginOauthCallback(ctx *gin.Context) {
	code := ctx.Request.URL.Query().Get("code")
	conf := oauth2.NewGoogleOauthConf()
	token, err := conf.Exchange(ctx, code)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, errors.NewErrorResponse(http.StatusBadGateway, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, token)

}

func generateRandomState() string {
	log := logger.GetLoggerDefault("OAuth2Controller")
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		log.Errorf("Erro ao gerar o estado aleatório: %v", err)
	}
	return base64.URLEncoding.EncodeToString(b)
}
