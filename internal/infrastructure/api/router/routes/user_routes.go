package routes

import (
	"os"

	"github.com/gin-gonic/gin"

	"github.com/oprimogus/cardapiogo/internal/domain/entity"
	"github.com/oprimogus/cardapiogo/internal/domain/repository"
	"github.com/oprimogus/cardapiogo/internal/infrastructure/api/controller"
	"github.com/oprimogus/cardapiogo/internal/infrastructure/api/middleware"
	validatorutils "github.com/oprimogus/cardapiogo/internal/infrastructure/api/validator"
)

func UserRoutes(router *gin.Engine,
	validator *validatorutils.Validator,
	userRepository repository.UserRepository,
	authRepository repository.AuthenticationRepository) {
	userController := controller.NewUserController(validator, userRepository)

	basepath := os.Getenv("API_BASE_PATH")

	v1 := router.Group(basepath + "/v1")
	{
		v1.POST("/auth/sign-up", userController.CreateUser)
		v1.PUT("/user",
			middleware.AuthenticationMiddleware(authRepository),
			middleware.AuthorizationMiddleware([]entity.UserRole{}, true),
			userController.UpdateUser)
	}
}
