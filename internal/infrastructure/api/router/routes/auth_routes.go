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

func AuthRoutes(router *gin.Engine, validator *validatorutils.Validator, authRepository repository.AuthenticationRepository) {
	authController := controller.NewAuthController(validator, authRepository)

	basePath := os.Getenv("API_BASE_PATH")

	v1 := router.Group(basePath + "/v1")
	{
		v1.POST("/auth/sign-in", authController.SignIn)
		v1.GET("/auth/test/authentication", middleware.AuthenticationMiddleware(authRepository), authController.ProtectedRoute)
		v1.GET("/auth/test/authorization", middleware.AuthenticationMiddleware(authRepository), middleware.AuthorizationMiddleware([]entity.UserRole{entity.UserRoleConsumer}, false), authController.ProtectedRouteForRoles)
	}
}
