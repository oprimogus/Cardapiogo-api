package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/oprimogus/cardapiogo/internal/api/controller"
	validatorutils "github.com/oprimogus/cardapiogo/internal/api/validator"
	"github.com/oprimogus/cardapiogo/internal/config"
	"github.com/oprimogus/cardapiogo/internal/core"
)

func AuthRoutes(router *gin.Engine, validator *validatorutils.Validator, factory core.RepositoryFactory) {
	authRepository := factory.NewAuthenticationRepository()
	authController := controller.NewAuthController(validator, authRepository)

	basePath := config.GetInstance().Api.BasePath()

	v1 := router.Group(basePath + "/v1")
	{
		v1.POST("/auth/sign-in", authController.SignIn)
		v1.POST("/auth/refresh", authController.RefreshUserToken)
	}
}
