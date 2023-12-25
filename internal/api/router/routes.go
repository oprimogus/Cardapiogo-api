package router

import (
	"github.com/gin-gonic/gin"
	"github.com/oprimogus/cardapiogo/internal/domain/factory"
	"github.com/oprimogus/cardapiogo/internal/domain/user"
	"github.com/oprimogus/cardapiogo/internal/api/validator"
)

// InitializeRoutes initialize all routes of application
func InitializeRoutes(router *gin.Engine, factory factory.RepositoryFactory) {
	validator, err := validatorutils.NewValidator("pt")
	if err != nil && validator == nil {
		panic(err)
	}
	userController := user.NewUserController(factory.NewUserRepository(), validator)

	basePath := "/api/v1"
	v1 := router.Group(basePath)
	v1.POST("user", userController.CreateUserHandler)

}
