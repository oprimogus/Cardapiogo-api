package routes

import (
	"os"

	"github.com/gin-gonic/gin"

	"github.com/oprimogus/cardapiogo/internal/infrastructure/api/controller"
	validatorutils "github.com/oprimogus/cardapiogo/internal/infrastructure/api/validator"
)

func AuthRoutes(router *gin.Engine, validator *validatorutils.Validator) {

	authController := controller.NewAuthController(validator)

	basePath := os.Getenv("API_BASE_PATH")

	v1 := router.Group(basePath + "/v1")
	{
		v1.POST("/auth/login", authController.SignUp)
		v1.GET("/auth/oauth", authController.SignInOrSignUpKeycloakOauth)
		v1.GET("/auth/oauth/callback", authController.SignInOrSignUpKeycloakOauthCallback)
	}

}
