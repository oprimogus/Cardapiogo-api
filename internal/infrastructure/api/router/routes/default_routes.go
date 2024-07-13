package routes

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func DefaultRoutes(router *gin.Engine, reg *prometheus.Registry) {
	basePath := os.Getenv("API_BASE_PATH")

	v1 := router.Group(basePath + "/v1")

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
