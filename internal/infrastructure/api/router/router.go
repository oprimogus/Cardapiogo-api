package router

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/oprimogus/cardapiogo/internal/core"
	"github.com/oprimogus/cardapiogo/internal/infrastructure/api/middleware"
	"github.com/oprimogus/cardapiogo/internal/infrastructure/api/router/routes"
	validatorutils "github.com/oprimogus/cardapiogo/internal/infrastructure/api/validator"
	logger "github.com/oprimogus/cardapiogo/pkg/log"
)

// Initialize API
func Initialize(factory core.RepositoryFactory) {
	validator, err := validatorutils.NewValidator("pt")
	if err != nil && validator == nil {
		panic(err)
	}

	logger := logger.NewLogger("Gin Router")
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middleware.CorsMiddleware())
	router.Use(middleware.LoggerMiddleware(logger))

	metrics := middleware.NewPrometheusMetrics()
	router.Use(middleware.PrometheusMiddleware(metrics))

	routes.DefaultRoutes(router, metrics.Registry)
	routes.SwaggerRoutes(router)
	routes.AuthRoutes(router, validator, factory)
	routes.UserRoutes(router, validator, factory)
	routes.StoreRoutes(router, validator, factory)
	port := os.Getenv("API_PORT")
	if port == "" {
		port = "8080"
	}

	logger.Info("Docs available in http://localhost:8080/api/v1/reference/index.html")
	logger.Info("Docs available in http://localhost:8080/api/v2/reference")
	log.Printf("Listening and serving in HTTP :%v\n", port)
	err = router.Run(":" + port)
	if err != nil {
		panic(err)
	}
}
