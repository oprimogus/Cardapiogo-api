package router

import (
	"github.com/gin-gonic/gin"

	"github.com/oprimogus/cardapiogo/internal/api/controller"
	"github.com/oprimogus/cardapiogo/internal/api/middleware"
	validatorutils "github.com/oprimogus/cardapiogo/internal/api/validator"
	"github.com/oprimogus/cardapiogo/internal/domain/factory"
)

// InitializeRoutes initialize all routes of application
func InitializeRoutes(router *gin.Engine, factory factory.RepositoryFactory) {
	validator, err := validatorutils.NewValidator("pt")
	if err != nil && validator == nil {
		panic(err)
	}
	userController := controller.NewUserController(factory.NewUserRepository(), validator)
	authController := controller.NewAuthController(factory.NewUserRepository(), validator)
	basePath := "/api/v1"

	v1 := router.Group(basePath)

	publicV1 := v1.Group("")
	{
		publicV1.POST("user", userController.CreateUserHandler)
		publicV1.GET("auth", authController.StartOAuthFlow)
		publicV1.POST("login", authController.Login)
		// publicV1.GET("auth/callback", oauthController.SignUpLoginOauthCallback)
	}

	authMiddleware := middleware.AuthMiddleware()

	protectedV1 := v1.Group("")
	protectedV1.Use(authMiddleware)
	{
		protectedV1.GET("user/:id", userController.GetUserHandler)
		protectedV1.GET("user", userController.GetUsersListHandler)
		protectedV1.PUT("user/change-password", userController.UpdateUserPasswordHandler)
		protectedV1.PUT("user", userController.UpdateUserHandler)
	}
}
