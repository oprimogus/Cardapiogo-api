package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/oprimogus/cardapiogo/internal/domain/entity"
	"github.com/oprimogus/cardapiogo/internal/infrastructure/errors"
)

func AuthorizationMiddleware(allowedRoles []entity.UserRole) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRolesContext, exists := c.Get("userRoles")
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errors.Forbidden(""))
			return
		}

		userRoles, ok := userRolesContext.([]entity.UserRole)
		if !ok {
			c.AbortWithStatusJSON(http.StatusInternalServerError, errors.Forbidden(""))
			return
		}
		if len(userRoles) == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errors.Forbidden(""))
			return
		}

		isAllowed := false
		for _, userRole := range userRoles {
			for _, allowedRole := range allowedRoles {
				if userRole == allowedRole {
					isAllowed = true
				}
			}
		}
		if !isAllowed {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errors.Forbidden(""))
			return
		}
		c.Next()

	}
}
