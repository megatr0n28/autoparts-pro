package router

import (
	"github.com/gin-gonic/gin"
	"github.com/megatr0n28/autoparts-pro/backend/internal/auth"
	"github.com/megatr0n28/autoparts-pro/backend/internal/handler"
	"github.com/megatr0n28/autoparts-pro/backend/internal/middleware"
)

func New(
	jwtManager *auth.JWTManager,
	userHandler *handler.UserHandler,
	authHandler *handler.AuthHandler,
	customerHandler *handler.CustomerHandler,
	vehicleHandler *handler.VehicleHandler,
) *gin.Engine {

	gin.SetMode(gin.ReleaseMode)
	r :=
		gin.Default()

	api :=
		r.Group("/api/v1")

	//
	// Public authentication routes
	//
	authRoutes := api.Group("/auth")

	authRoutes.POST(
		"/login",
		authHandler.Login,
	)

	authRoutes.POST(
		"/register",
		authHandler.Register,
	)

	authRoutes.POST(
		"/refresh",
		authHandler.Refresh,
	)

	authRoutes.POST(
		"/logout",
		authHandler.Logout,
	)

	//
	// Protected user routes
	//
	protected := api.Group("")
	protected.Use(
		middleware.JWTAuth(jwtManager),
	)

	protected.GET(
		"/users/me",
		userHandler.Me,
	)

	vehicles := protected.Group("/vehicles")
	vehicles.GET(
		"",
		vehicleHandler.List,
	)

	vehicles.POST(
		"",
		vehicleHandler.Create,
	)

	protected.GET(
		"/customers/me",
		customerHandler.Me,
	)

	protected.PUT(
		"/customers/me",
		customerHandler.Update,
	)

	protected.POST(
		"/auth/logout-all",
		authHandler.LogoutAll,
	)

	//
	// Admin routes
	//
	admin :=
		api.Group("/admin")

	admin.Use(
		middleware.JWTAuth(jwtManager),
	)

	admin.Use(
		middleware.RequireRole(
			"admin",
		),
	)

	admin.GET(
		"/health",
		func(c *gin.Context) {

			c.JSON(
				200,
				gin.H{
					"status": "admin access",
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
