package http

import (
	"github.com/gin-gonic/gin"

	"github.com/megatr0n28/autoparts-pro/backend/internal/interfaces/http/handlers"
)

func NewRouter() *gin.Engine {

	router := gin.New()

	router.Use(
		gin.Logger(),
		gin.Recovery(),
	)

	healthHandler :=
		handlers.NewHealthHandler()

	api := router.Group("/api")

	{

		v1 := api.Group("/v1")

		{

			v1.GET(
				"/health",
				healthHandler.Health,
			)

			v1.GET(
				"/live",
				healthHandler.Live,
			)

			v1.GET(
				"/ready",
				healthHandler.Ready,
			)

		}

	}

	return router
}
