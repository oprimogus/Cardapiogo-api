package router

import (
	"os"
	"strings"
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

	log := logger.GetLoggerDefault("Router")

	router := gin.New()
	setGinMode()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{os.Getenv("WEB_URL")},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	router.Use(middleware.LoggerMiddleware(logger.GetLoggerDefault("GIN")))

	metrics := middleware.NewPrometheusMetrics()
	router.Use(middleware.PrometheusMiddleware(metrics))

	routes.DefaultRoutes(router, factory, metrics.Registry)
	routes.AuthRoutes(router, factory, validator)
	routes.UserRoutes(router, factory, validator)
	routes.ProfileRoutes(router, factory, validator)

	router.Use(gin.Recovery())

	const host = "0.0.0.0"
	port := os.Getenv("API_PORT")
	if port == "" {
		port = "8080"
	}

	log.Infof("Running server in %s:%s", host, port)
	router.Run(host + ":" + port)
}

func setGinMode() {
	env := strings.ToLower(os.Getenv("API_ENVIRONMENT"))
	if env == "local" || env == "staging" {
		gin.SetMode(gin.DebugMode)
	}
	if env == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}
}
