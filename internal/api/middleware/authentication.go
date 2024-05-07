package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"

	"github.com/oprimogus/cardapiogo/internal/errors"
)

func AuthenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errors.Unauthorized(""))
			return
		}
		token = strings.Replace(token, "Bearer ", "", -1)
		validatedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("signature method not expected: %v", token.Header["alg"])
			}
			if err := token.Claims.Valid(); err != nil {
				return nil, fmt.Errorf("invalid token: %v", err.Error())
			}

			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errors.New(http.StatusUnauthorized, err.Error()))
			return
		}

		claims, ok := validatedToken.Claims.(jwt.MapClaims)
		if !ok && !validatedToken.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errors.New(http.StatusUnauthorized, err.Error()))
			return
		}
		c.Set("userRole", claims["role"].(string))
		c.Set("userID", claims["sub"].(string))
		c.Next()
	}
}
