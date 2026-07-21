package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
