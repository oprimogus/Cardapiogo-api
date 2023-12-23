package router

import (
	"github.com/gin-gonic/gin"
	"github.com/oprimogus/cardapiogo/internal/api/controller"
	"github.com/oprimogus/cardapiogo/internal/domain/factory"
)

// InitializeRoutes initialize all routes of application
func InitializeRoutes(router *gin.Engine, factory factory.RepositoryFactory) {

	userController := controller.NewUserController(factory.NewUserRepository())

	basePath := "/api/v1"
	v1 := router.Group(basePath)
	v1.POST("user", userController.CreateUserHandler)

}
