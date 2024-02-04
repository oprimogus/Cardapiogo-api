package routes

import (
	"os"

	"github.com/gin-gonic/gin"

	"github.com/oprimogus/cardapiogo/internal/api/controller"
	"github.com/oprimogus/cardapiogo/internal/api/middleware"
	validatorutils "github.com/oprimogus/cardapiogo/internal/api/validator"
	"github.com/oprimogus/cardapiogo/internal/domain/factory"
)

func UserRoutes(router *gin.Engine, factory factory.RepositoryFactory, validator *validatorutils.Validator) {

	userController := controller.NewUserController(factory.NewUserRepository(), validator)

	basePath := os.Getenv("API_BASE_PATH")

	v1 := router.Group(basePath)
	{
		v1.POST("user", userController.CreateUserHandler)
		v1.GET("user/:id", middleware.AuthMiddleware(), userController.GetUserHandler)
		v1.GET("user", middleware.AuthMiddleware(), userController.GetUsersListHandler)
		v1.PUT("user/change-password", middleware.AuthMiddleware(), userController.UpdateUserPasswordHandler)
		v1.PUT("user", middleware.AuthMiddleware(), userController.UpdateUserHandler)
	}

}
