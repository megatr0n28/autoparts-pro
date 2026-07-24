package router

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/megatr0n28/autoparts-pro/backend/internal/auth"
	"github.com/megatr0n28/autoparts-pro/backend/internal/handler"
	"github.com/megatr0n28/autoparts-pro/backend/internal/middleware"
	"github.com/megatr0n28/autoparts-pro/backend/internal/repository"
)

func New(
	jwtManager *auth.JWTManager,
	userHandler *handler.UserHandler,
	authHandler *handler.AuthHandler,
	customerHandler *handler.CustomerHandler,
	vehicleHandler *handler.VehicleHandler,
	customerRepository repository.CustomerRepository,
	searchHandler *handler.SearchHandler,
) *gin.Engine {

	router := gin.Default()
	router.SetTrustedProxies(nil)
	router.Use(
		cors.New(cors.Config{

			AllowOrigins: []string{
				"http://localhost:5173",
			},

			AllowMethods: []string{
				"GET",
				"POST",
				"PUT",
				"PATCH",
				"DELETE",
			},

			AllowHeaders: []string{
				"Authorization",
				"Content-Type",
			},

			AllowCredentials: true,

			MaxAge: 12 * time.Hour,
		}),
	)

	api :=
		router.Group(
			"/api/v1",
		)

	//
	// Public authentication routes
	//
	authRoutes := api.Group("/auth")

	authRoutes.POST(
		"/register",
		authHandler.Register,
	)

	authRoutes.POST(
		"/login",
		authHandler.Login,
	)

	authRoutes.POST(
		"/refresh",
		authHandler.Refresh,
	)

	authRoutes.POST(
		"/logout",
		authHandler.Logout,
	)

	authRoutes.POST(
		"/logout-all",
		authHandler.LogoutAll,
	)

	//
	// Protected user routes
	//
	protected := api.Group("")
	protected.Use(
		middleware.JWTAuth(
			jwtManager,
			customerRepository,
		),
	)

	// ----------------------------
	// User Routes
	// -
	protected.GET(
		"/users/me",
		userHandler.Me,
	)

	// ----------------------------
	// Customer Routes
	// ----------------------------
	protected.GET(
		"/customers/me",
		customerHandler.Me,
	)

	protected.PUT(
		"/customers/me",
		customerHandler.Update,
	)

	// ----------------------------
	// Vehicle Routes
	// ----------------------------
	vehicles := protected.Group("/vehicles")

	vehicles.POST(
		"",
		vehicleHandler.Create,
	)

	vehicles.GET(
		"",
		vehicleHandler.List,
	)

	vehicles.DELETE(
		"/:id",
		vehicleHandler.Delete,
	)

	vehicles.PATCH(
		"/:id/primary",
		vehicleHandler.SetPrimary,
	)

	vehicles.PUT(
		"/:id",
		vehicleHandler.Update,
	)

	// ----------------------------
	// Parts Search Routes
	// ----------------------------
	parts := protected.Group("/parts")
	{
		parts.GET(
			"/search",
			searchHandler.SearchParts,
		)
	}

	//
	// Admin routes
	//
	admin :=
		api.Group("/admin")

	admin.Use(
		middleware.JWTAuth(
			jwtManager,
			customerRepository,
		),
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

	//
	// Health check route
	//
	api.GET(
		"/health",
		func(c *gin.Context) {

			c.JSON(
				200,
				gin.H{
					"status": "ok",
				},
			)

		},
	)

	return router
}
