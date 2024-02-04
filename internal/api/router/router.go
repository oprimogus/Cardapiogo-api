package router

import (
	"os"

	"github.com/gin-gonic/gin"

	"github.com/oprimogus/cardapiogo/internal/api/middleware"
	"github.com/oprimogus/cardapiogo/internal/api/router/routes"
	validatorutils "github.com/oprimogus/cardapiogo/internal/api/validator"
	"github.com/oprimogus/cardapiogo/internal/domain/factory"
	logger "github.com/oprimogus/cardapiogo/pkg/log"
)

// Initialize API
func Initialize(factory factory.RepositoryFactory) {

	validator, err := validatorutils.NewValidator("pt")
	if err != nil && validator == nil {
		panic(err)
	}

	log := logger.GetLoggerDefault("Router")

	router := gin.New()
	router.Use(middleware.CorsMiddleware())
	router.Use(middleware.TransactionIDMiddleware())
	router.Use(middleware.LoggerMiddleware(logger.GetLoggerDefault("GIN")))

	routes.DefaultRoutes(router, factory)
	routes.UserRoutes(router, factory, validator)
	routes.AuthRoutes(router, factory, validator)

	router.Use(gin.Recovery())

	const host = "0.0.0.0"
	port := os.Getenv("API_PORT")
	if port == "" {
		port = "8080"
	}

	log.Infof("Running server in %s:%s", host, port)
	router.Run(host + ":" + port)

}
