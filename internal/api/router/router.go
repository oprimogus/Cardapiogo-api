package router

import (
	"os"
	"github.com/gin-gonic/gin"
	"github.com/oprimogus/cardapiogo/internal/domain/factory"
	logger "github.com/oprimogus/cardapiogo/pkg/log"
)

// Initialize API
func Initialize(factory factory.RepositoryFactory) {
	log := logger.GetLogger("Main")
	router :=  gin.Default()
	InitializeRoutes(router, factory)

	const host = "0.0.0.0"
	port := os.Getenv("API_PORT")
	if port == "" {
		port = "8080"
	}

	router.Run(host + ":" + port)
	log.Infof("Running server in %s:%s", host, port)

}
