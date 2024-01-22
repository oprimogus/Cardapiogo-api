package middleware

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extrair o cookie
		cookie, err := c.Cookie("token")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Não autorizado"})
			return
		}

		// Validar o token
		token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
			// Certificar-se de que o algoritmo de assinatura esperado é utilizado
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Método de assinatura inesperado: %v", token.Header["alg"])
			}

			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Atribuir informações do usuário ao contexto, se necessário
			c.Set("userID", claims["sub"])
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Não autorizado"})
			return
		}

		c.Next()
	}
}
