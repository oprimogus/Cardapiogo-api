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

func StoreRoutes(router *gin.Engine,
	validator *validatorutils.Validator,
	storeRepository repository.StoreRepository,
	authRepository repository.AuthenticationRepository) {
	storeController := controller.NewStoreController(validator, storeRepository)

	basepath := os.Getenv("API_BASE_PATH")

	v1 := router.Group(basepath + "/v1")
	{
		v1.POST("/store",
			middleware.AuthenticationMiddleware(authRepository),
			middleware.AuthorizationMiddleware([]entity.UserRole{entity.UserRoleOwner}),
			storeController.Create)
		v1.PUT("/store",
			middleware.AuthenticationMiddleware(authRepository),
			middleware.AuthorizationMiddleware([]entity.UserRole{entity.UserRoleOwner}),
			storeController.Update)
		v1.PUT("/store/business-hour",
			middleware.AuthenticationMiddleware(authRepository),
			middleware.AuthorizationMiddleware([]entity.UserRole{entity.UserRoleOwner}),
			storeController.AddBusinessHours)
		v1.DELETE("/store/business-hour",
			middleware.AuthenticationMiddleware(authRepository),
			middleware.AuthorizationMiddleware([]entity.UserRole{entity.UserRoleOwner}),
			storeController.DeleteBusinessHours)
		v1.GET("/store/:id", storeController.GetStoreByID)
	}
}
