package router

import (
	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {

	r :=
		gin.Default()

	api :=
		r.Group("/api/v1")

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
