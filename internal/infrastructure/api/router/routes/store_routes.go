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

func StoreRoutes(router *gin.Engine,
	validator *validatorutils.Validator,
	factory core.RepositoryFactory) {
	authRepository := factory.NewAuthenticationRepository()
	storeRepository := factory.NewStoreRepository()
	storeController := controller.NewStoreController(validator, storeRepository)

	basepath := os.Getenv("API_BASE_PATH")

	v1 := router.Group(basepath + "/v1")
	{
		v1.POST("/store",
			middleware.AuthenticationMiddleware(authRepository),
			middleware.AuthorizationMiddleware([]user.Role{user.RoleOwner}),
			storeController.Create)
		v1.PUT("/store",
			middleware.AuthenticationMiddleware(authRepository),
			middleware.AuthorizationMiddleware([]user.Role{user.RoleOwner}),
			storeController.Update)
		v1.PUT("/store/business-hour",
			middleware.AuthenticationMiddleware(authRepository),
			middleware.AuthorizationMiddleware([]user.Role{user.RoleOwner}),
			storeController.AddBusinessHours)
		v1.DELETE("/store/business-hour",
			middleware.AuthenticationMiddleware(authRepository),
			middleware.AuthorizationMiddleware([]user.Role{user.RoleOwner}),
			storeController.DeleteBusinessHours)
		v1.GET("/store/:id", storeController.GetStoreByID)
		v1.GET("/store", storeController.GetStoreByFilter)
	}
}
