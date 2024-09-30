package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/oprimogus/cardapiogo/internal/api/controller"
	"github.com/oprimogus/cardapiogo/internal/api/middleware"
	validatorutils "github.com/oprimogus/cardapiogo/internal/api/validator"
	"github.com/oprimogus/cardapiogo/internal/config"
	"github.com/oprimogus/cardapiogo/internal/core"
	"github.com/oprimogus/cardapiogo/internal/core/user"
)

func StoreRoutes(router *gin.Engine,
	validator *validatorutils.Validator,
	factory core.RepositoryFactory) {
	authRepository := factory.NewAuthenticationRepository()
	storeRepository := factory.NewStoreRepository()
	storeController := controller.NewStoreController(validator, storeRepository)

	basePath := config.GetInstance().Api.BasePath()

	v1 := router.Group(basePath + "/v1")
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
		v1.PUT("/store/profile-image",
			middleware.AuthenticationMiddleware(authRepository),
			middleware.AuthorizationMiddleware([]user.Role{user.RoleOwner}),
			storeController.SetProfileImage)
		v1.PUT("/store/header-image",
			middleware.AuthenticationMiddleware(authRepository),
			middleware.AuthorizationMiddleware([]user.Role{user.RoleOwner}),
			storeController.SetHeaderImage)
	}
}
