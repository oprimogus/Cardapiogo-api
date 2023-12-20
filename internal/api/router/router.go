package router

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/oprimogus/cardapiogo/internal/domain/factory"
)

func Initialize(factory factory.RepositoryFactory) {
	router := gin.Default()

	InitializeRoutes(router, factory)

	port := os.Getenv("API_PORT")
	if port == "" {
		port = "8080"
	}

	router.Run("0.0.0.0:" + port)
	

}