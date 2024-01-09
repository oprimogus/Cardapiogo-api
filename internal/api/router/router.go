package router

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/oprimogus/cardapiogo/internal/domain/factory"
	logger "github.com/oprimogus/cardapiogo/pkg/log"
)

// Initialize API
func Initialize(factory factory.RepositoryFactory) {
	log := logger.GetLogger("Router")

	router := gin.New()
	router.Use(LoggerMiddleware(logger.GetLogger("GIN")))

	InitializeRoutes(router, factory)

	const host = "0.0.0.0"
	port := os.Getenv("API_PORT")
	if port == "" {
		port = "8080"
	}

	log.Infof("Running server in %s:%s", host, port)
	router.Run(host + ":" + port)

}

func LoggerMiddleware(logger *logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Início do request
		t := time.Now()

		// Processa a próxima função no encadeamento
		c.Next()

		// Após o request ser processado
		latency := time.Since(t).String()
		status := c.Writer.Status()
		clientIP := c.ClientIP()
		method := c.Request.Method
		path := c.Request.URL.Path

		// Use GetEntry para acessar logrus.Entry
		entry := logger.GetEntry()

		// Loga os detalhes do request como campos estruturados
		entry.WithFields(logrus.Fields{
			"client_ip": clientIP,
			"method":    method,
			"path":      path,
			"status":    status,
			"latency":   latency,
		}).Info("Request handled")
	}
}
