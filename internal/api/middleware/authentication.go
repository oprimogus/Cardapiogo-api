package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/oprimogus/cardapiogo/internal/core/authentication"
	"github.com/oprimogus/cardapiogo/internal/core/user"
	"github.com/oprimogus/cardapiogo/internal/errors"
)

func AuthenticationMiddleware(repository authentication.Repository) gin.HandlerFunc {
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

		userRoles := []user.Role{}
		for _, role := range roles {
			if roleStr, ok := role.(string); ok {
				if user.IsValidRole(roleStr) {
					userRoles = append(userRoles, user.Role(roleStr))
				}

			}
		}
		c.Set("userID", claims["sub"])
		c.Set("userRoles", userRoles)
		c.Next()
	}
}
