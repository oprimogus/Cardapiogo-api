package middleware

import (
	"time"

	"github.com/gin-gonic/gin"

	logger "github.com/oprimogus/cardapiogo/pkg/log"
)

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
		transactionId, exists := c.Get("transactionId")
		if !exists {
			transactionId = ""
		}

		entry := logger.GetEntry()

		// Loga os detalhes do request como campos estruturados
		entry.With(
			"transactionid", transactionId,
			"client_ip", clientIP,
			"method", method,
			"path", path,
			"status", status,
			"latency", latency,
		).Info("Request Handled")
	}
}
