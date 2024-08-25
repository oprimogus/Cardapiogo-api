package routes

import (
	"os"

	"github.com/gin-gonic/gin"

	"github.com/oprimogus/cardapiogo/internal/core"
	"github.com/oprimogus/cardapiogo/internal/infrastructure/api/controller"
	validatorutils "github.com/oprimogus/cardapiogo/internal/infrastructure/api/validator"
)

func AuthRoutes(router *gin.Engine, validator *validatorutils.Validator, factory core.RepositoryFactory) {
	authRepository := factory.NewAuthenticationRepository()
	authController := controller.NewAuthController(validator, authRepository)

	basePath := os.Getenv("API_BASE_PATH")

	v1 := router.Group(basePath + "/v1")
	{
		v1.POST("/auth/sign-in", authController.SignIn)
		v1.POST("/auth/refresh", authController.RefreshUserToken)
	}
}
