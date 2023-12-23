package router

import (
	"github.com/gin-gonic/gin"
	"github.com/oprimogus/cardapiogo/internal/domain/factory"
	"github.com/oprimogus/cardapiogo/internal/domain/user"
	logger "github.com/oprimogus/cardapiogo/pkg/log"
	"github.com/oprimogus/cardapiogo/pkg/validator"
)

var log = logger.GetLogger("router")

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
