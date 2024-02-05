package middleware

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"

	"github.com/oprimogus/cardapiogo/internal/errors"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("token")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errors.Unauthorized(""))
			return
		}
		if cookie == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errors.Unauthorized(""))
			return
		}
		token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
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

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("userID", claims["sub"])
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errors.New(http.StatusUnauthorized, err.Error()))
			return
		}

		c.Next()
	}
}
