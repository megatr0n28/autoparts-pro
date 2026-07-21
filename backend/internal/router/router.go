package router

import (
	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {

	r :=
		gin.Default()

	api :=
		r.Group("/api/v1")

	authRoutes := api.Group("/auth")

	authRoutes.POST(
		"/login",
		func(c *gin.Context) {

			c.JSON(
				200,
				gin.H{
					"message": "login endpoint",
				},
			)

		},
	)

	authRoutes.POST(
		"/register",
		func(c *gin.Context) {

			c.JSON(
				201,
				gin.H{
					"message": "register endpoint",
				},
			)

		},
	)

	api.GET(
		"health",
		func(c *gin.Context) {

			c.JSON(
				200,
				gin.H{
					"status": "ok",
				},
			)

		},
	)

	return r
}
