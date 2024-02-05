package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	logger "github.com/oprimogus/cardapiogo/pkg/log"
)

func LoggerMiddleware(logger *logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Início do request
		t := time.Now()
		transactionId := uuid.New().String()
		c.Set("transactionId", transactionId)

		// Processa a próxima função no encadeamento
		c.Next()

		// Após o request ser processado
		latency := time.Since(t).String()
		c.Set("latency", latency)
		status := c.Writer.Status()
		clientIP := c.ClientIP()
		method := c.Request.Method
		path := c.Request.URL.Path
		userID, exists := c.Get("userID")
		if !exists {
			userID = ""
		}

		entry := logger.GetEntry()

		// Loga os detalhes do request como campos estruturados
		entry.With(
			"transactionid", transactionId,
			"userid", userID,
			"client_ip", clientIP,
			"method", method,
			"path", path,
			"status", status,
			"latency", latency,
		).Info("Request Handled")
	}
}
