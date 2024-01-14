package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func TransactionIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		transactionId := uuid.New().String()
		c.Set("transactionId", transactionId)
		c.Next()
	}
}
