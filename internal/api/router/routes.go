package router

import (
	"github.com/gin-gonic/gin"
	"github.com/oprimogus/cardapiogo/internal/domain/factory"
)

func InitializeRoutes(router *gin.Engine, factory factory.RepositoryFactory) {

	basePath := "/api/v1"
	v1 := router.Group(basePath)

	v1.POST("user",)

	
	
}