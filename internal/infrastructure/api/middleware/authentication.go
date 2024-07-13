package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/oprimogus/cardapiogo/internal/domain/entity"
	"github.com/oprimogus/cardapiogo/internal/domain/repository"
	"github.com/oprimogus/cardapiogo/internal/infrastructure/errors"
)

func AuthenticationMiddleware(repository repository.AuthenticationRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, xerrors.Unauthorized(""))
			return
		}
		token = strings.Replace(token, "Bearer ", "", -1)
		isValidToken, err := repository.IsValidToken(c, token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, xerrors.Unauthorized(err.Error()))
			return
		}

		if !isValidToken {
			c.AbortWithStatusJSON(http.StatusUnauthorized, xerrors.Unauthorized("Invalid access token"))
			return
		}

		claims, err := repository.DecodeAccessToken(c, token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, xerrors.New(http.StatusUnauthorized, err.Error()))
			return
		}

		realmAccess, ok := claims["realm_access"].(map[string]interface{})
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, xerrors.New(http.StatusUnauthorized, "Invalid token"))
			return
		}

		roles, ok := realmAccess["roles"].([]interface{})
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, xerrors.New(http.StatusUnauthorized, "Invalid token"))
			return
		}

		userRoles := []entity.UserRole{}
		for _, role := range roles {
			if roleStr, ok := role.(string); ok {
				if entity.IsValidUserRole(roleStr) {
					userRoles = append(userRoles, entity.UserRole(roleStr))
				}

			}
		}
		c.Set("userID", claims["sub"])
		c.Set("userRoles", userRoles)
		c.Next()
	}
}
