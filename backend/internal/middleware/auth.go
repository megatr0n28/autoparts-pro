package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/megatr0n28/autoparts-pro/backend/internal/auth"
)

func JWTAuth(
	manager *auth.JWTManager,
) gin.HandlerFunc {

	return func(c *gin.Context) {

		header :=
			c.GetHeader(
				"Authorization",
			)

		if !strings.HasPrefix(
			header,
			"Bearer ",
		) {

			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{
					"error": "missing token",
				},
			)

			return
		}

		token :=
			strings.TrimPrefix(
				header,
				"Bearer ",
			)

		claims, err :=
			manager.ValidateToken(
				token,
			)

		if err != nil {

			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{
					"error": "invalid token",
				},
			)

			return
		}

		c.Set(
			"user_id",
			claims.UserID,
		)

		c.Set(
			"role",
			claims.Role,
		)

		c.Next()

	}

}
