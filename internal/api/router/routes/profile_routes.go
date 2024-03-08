package routes

import (
	"os"

	"github.com/gin-gonic/gin"

	"github.com/oprimogus/cardapiogo/internal/api/controller"
	"github.com/oprimogus/cardapiogo/internal/api/middleware"
	validatorutils "github.com/oprimogus/cardapiogo/internal/api/validator"
	"github.com/oprimogus/cardapiogo/internal/domain/factory"
)

func ProfileRoutes(router *gin.Engine, factory factory.RepositoryFactory, validator *validatorutils.Validator) {

	profileController := controller.NewProfileController(factory.NewProfileRepository(), factory.NewUserRepository(), validator)

	basePath := os.Getenv("API_BASE_PATH")

	v1 := router.Group(basePath + "/v1")
	{
		v1.POST("/profile", middleware.AuthenticationMiddleware(), profileController.CreateProfileHandler)
		v1.GET("/profile", middleware.AuthenticationMiddleware(), profileController.GetProfileByUserIDHandler)
		v1.PUT("/profile", middleware.AuthenticationMiddleware(), profileController.UpdateProfileHandler)
	}

}
