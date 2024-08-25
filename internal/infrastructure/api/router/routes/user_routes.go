package routes

import (
	"os"

	"github.com/gin-gonic/gin"

	"github.com/oprimogus/cardapiogo/internal/core"
	"github.com/oprimogus/cardapiogo/internal/core/user"
	"github.com/oprimogus/cardapiogo/internal/infrastructure/api/controller"
	"github.com/oprimogus/cardapiogo/internal/infrastructure/api/middleware"
	validatorutils "github.com/oprimogus/cardapiogo/internal/infrastructure/api/validator"
)

func UserRoutes(router *gin.Engine,
	validator *validatorutils.Validator,
	factory core.RepositoryFactory) {
	userRepository := factory.NewUserRepository()
	authRepository := factory.NewAuthenticationRepository()
	userController := controller.NewUserController(validator, userRepository)

	basepath := os.Getenv("API_BASE_PATH")

	v1 := router.Group(basepath + "/v1")
	{
		v1.POST("/auth/sign-up", userController.CreateUser)
		v1.PUT("/user",
			middleware.AuthenticationMiddleware(authRepository),
			middleware.AuthorizationMiddleware([]user.Role{}),
			userController.UpdateUser)
		v1.POST("/user/roles",
			middleware.AuthenticationMiddleware(authRepository),
			middleware.AuthorizationMiddleware([]user.Role{}),
			userController.AddRolesToUser)
	}
}
