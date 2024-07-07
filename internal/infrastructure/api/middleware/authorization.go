package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/oprimogus/cardapiogo/internal/domain/entity"
	xerrors "github.com/oprimogus/cardapiogo/internal/infrastructure/errors"
)

func AuthorizationMiddleware(allowedRoles []entity.UserRole, isUser bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !isUser {
			c.AbortWithStatusJSON(http.StatusForbidden, xerrors.Forbidden("You aren't resource owner"))
			return
		}

		userRolesContext, exists := c.Get("userRoles")
		if !exists {
			c.AbortWithStatusJSON(http.StatusForbidden, xerrors.Forbidden(""))
			return
		}

		userRoles, ok := userRolesContext.([]entity.UserRole)
		if !ok {
			c.AbortWithStatusJSON(http.StatusInternalServerError, xerrors.Forbidden(""))
			return
		}
		if len(userRoles) == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, xerrors.Forbidden(""))
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
		if !isAllowed && len(allowedRoles) != 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, xerrors.Forbidden(""))
			return
		}
		c.Next()

	}
}
