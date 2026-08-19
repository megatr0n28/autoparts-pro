package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/megatr0n28/autoparts-pro/backend/internal/auth"
)

func RequireRole(
	role string,
) gin.HandlerFunc {

	return func(c *gin.Context) {

		userRole,
			exists :=
			c.Get("role")

		if !exists {

			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{
					"error": "missing role",
				},
			)

			return
		}

		if userRole != role {

			c.AbortWithStatusJSON(
				http.StatusForbidden,
				gin.H{
					"error": "insufficient permissions",
				},
			)

			return
		}

		c.Next()

	}

}

func RequireWriteAccess() gin.HandlerFunc {
	return func(c *gin.Context) {
		role := c.GetString("role")
		if !auth.CanModifyOwnData(role) {
			c.AbortWithStatusJSON(
				http.StatusForbidden,
				gin.H{"error": "viewer accounts are read-only"},
			)
			return
		}
		c.Next()
	}
}
