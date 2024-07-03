package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/oprimogus/cardapiogo/internal/domain/repository"
	"github.com/oprimogus/cardapiogo/internal/infrastructure/errors"
)

func AuthenticationMiddleware(repository repository.AuthenticationRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errors.Unauthorized(""))
			return
		}
		token = strings.Replace(token, "Bearer ", "", -1)
		isValidToken, err := repository.IsValidToken(c, token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errors.Unauthorized(err.Error()))
			return
		}

		if !isValidToken {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errors.Unauthorized("Invalid access token"))
			return
		}

		claims, err := repository.DecodeAccessToken(c, token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errors.New(http.StatusUnauthorized, err.Error()))
			return
		}
		c.Set("userID", claims["sub"])
		c.Next()
	}
}
