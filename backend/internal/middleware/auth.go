package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/megatr0n28/autoparts-pro/backend/internal/auth"
	"github.com/megatr0n28/autoparts-pro/backend/internal/repository"
)

func JWTAuth(
	manager *auth.JWTManager,
	customerRepo repository.CustomerRepository,
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

		userID, err :=
			uuid.Parse(
				claims.UserID,
			)

		if err != nil {

			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{
					"error": "invalid user id",
				},
			)

			return
		}

		c.Set(
			"user_id",
			userID,
		)
		customer, err :=
			customerRepo.FindByUserID(
				c,
				userID,
			)

		if err != nil {

			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{
					"error": "customer profile not found",
				},
			)

			return
		}

		c.Set(
			"customer_id",
			customer.ID.String(),
		)

		c.Set(
			"role",
			claims.Role,
		)

		c.Next()

	}

}
