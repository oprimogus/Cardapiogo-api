package router

import (
	"github.com/gin-gonic/gin"
	"github.com/oprimogus/cardapiogo/internal/api/controller"
	"github.com/oprimogus/cardapiogo/internal/api/validator"
	"github.com/oprimogus/cardapiogo/internal/domain/factory"
)

// InitializeRoutes initialize all routes of application
func InitializeRoutes(router *gin.Engine, factory factory.RepositoryFactory) {
	validator, err := validatorutils.NewValidator("pt")
	if err != nil && validator == nil {
		panic(err)
	}
	userController := controller.NewUserController(factory.NewUserRepository(), validator)

	basePath := "/api/v1"
	v1 := router.Group(basePath)
	v1.POST("user", userController.CreateUserHandler)
	v1.GET("user/:id", userController.GetUserHandler)
	v1.GET("user", userController.GetUsersListHandler)

}
