package router

import (
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
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

	logger := logger.GetLoggerDefault("GIN Router")
	router := gin.New()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Accept",
			"Authorization",
			"X-Requested-With",
		},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	router.Use(middleware.LoggerMiddleware(logger))

	metrics := middleware.NewPrometheusMetrics()
	router.Use(middleware.PrometheusMiddleware(metrics))

	routes.DefaultRoutes(router, factory, metrics.Registry)
	routes.SwaggerRoutes(router)
	routes.AuthRoutes(router, factory, validator)
	routes.UserRoutes(router, factory, validator)
	routes.ProfileRoutes(router, factory, validator)

	router.Use(gin.Recovery())

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
