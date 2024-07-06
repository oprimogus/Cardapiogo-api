package routes

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/oprimogus/cardapiogo/internal/domain/repository"
	"github.com/oprimogus/cardapiogo/internal/infrastructure/api/controller"
	validatorutils "github.com/oprimogus/cardapiogo/internal/infrastructure/api/validator"
)

func UserRoutes(router *gin.Engine, validator *validatorutils.Validator, factory repository.Factory) {
	userController := controller.NewUserController(validator, factory)

	basepath := os.Getenv("API_BASE_PATH")

	v1 := router.Group(basepath + "/v1")
	{
		v1.POST("/auth/sign-up", userController.CreateUser)
	}
}
