package routes

import (
	"os"

	"github.com/gin-gonic/gin"

	"github.com/oprimogus/cardapiogo/internal/api/controller"
	validatorutils "github.com/oprimogus/cardapiogo/internal/api/validator"
	"github.com/oprimogus/cardapiogo/internal/domain/factory"
)

func AuthRoutes(router *gin.Engine, factory factory.RepositoryFactory, validator *validatorutils.Validator) {

	authController := controller.NewAuthController(factory.NewUserRepository(), validator)

	basePath := os.Getenv("API_BASE_PATH")

	v1 := router.Group(basePath + "/v1")
	{
		v1.POST("/login", authController.Login)
		v1.GET("/auth/google", authController.StartGoogleOAuthFlow)
		v1.GET("/auth/google/callback", authController.SignUpLoginGoogleOauthCallback)
	}

}
