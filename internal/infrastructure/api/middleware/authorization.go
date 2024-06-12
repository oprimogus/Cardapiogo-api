package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/oprimogus/cardapiogo/internal/domain/entity"
	"github.com/oprimogus/cardapiogo/internal/infrastructure/errors"
)

func AuthorizationMiddleware(allowedRoles []entity.UserRole) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRoleInterface, exists := c.Get("userRole")
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errors.Unauthorized(""))
			return
		}

		userRoleString, ok := userRoleInterface.(string)
		if !ok {
			c.AbortWithStatusJSON(http.StatusInternalServerError, errors.InternalServerError("Error in process userRole"))
			return
		}

		userRole := entity.UserRole(userRoleString)

		isAllowed := false
		for _, role := range allowedRoles {
			if role == userRole {
				isAllowed = true
				break
			}
		}

		if !isAllowed {
			c.AbortWithStatusJSON(http.StatusForbidden, errors.Forbidden(""))
			return
		}

		c.Next()
	}
}
