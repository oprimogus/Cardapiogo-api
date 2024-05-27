package routes

import (
	"os"

	"github.com/MarceloPetrucio/go-scalar-api-reference"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	docs "github.com/oprimogus/cardapiogo/docs"
	"github.com/oprimogus/cardapiogo/internal/domain/factory"
)

func DefaultRoutes(
	router *gin.Engine,
	factory factory.RepositoryFactory,
	reg *prometheus.Registry,
) {
	basePath := os.Getenv("API_BASE_PATH")

	v1 := router.Group(basePath + "/v1")

	// SWAGGER
	docs.SwaggerInfo.BasePath = basePath
	v1.GET("/reference", func(c *gin.Context) {
		htmlContent, err := scalar.ApiReferenceHTML(&scalar.Options{
			SpecURL: "docs/swagger.json",
			CustomOptions: scalar.CustomOptions{
				PageTitle: "Cardapiogo-API",
			},
			DarkMode: true,
		})
		if err != nil {
			c.JSON(500, gin.H{
				"error": err,
			})
			return
		}
		c.Data(200, "text/html; charset=utf-8", []byte(htmlContent))
	})

	// Health
	v1.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})

	// Prometheus
	v1.GET("/metrics", func(c *gin.Context) {
		h := promhttp.HandlerFor(reg, promhttp.HandlerOpts{})
		h.ServeHTTP(c.Writer, c.Request)
	})
}
