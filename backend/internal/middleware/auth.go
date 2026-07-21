package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTAuth(
	secret string,
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

		c.Next()
	}

}
