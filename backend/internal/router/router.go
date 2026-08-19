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
	dashboardHandler *handler.DashboardHandler,
	adminHandler *handler.AdminHandler,
	customerRepository repository.CustomerRepository,
	searchHandler *handler.SearchHandler,
) *gin.Engine {

	router := gin.Default()
	router.SetTrustedProxies(nil)
	router.Use(
		cors.New(cors.Config{

			AllowOrigins: []string{
				"http://localhost:4200",
			},

			AllowMethods: []string{
				"GET",
				"POST",
				"PUT",
				"DELETE",
				"OPTIONS",
			},

			AllowHeaders: []string{
				"Origin",
				"Content-Type",
				"Authorization",
			},

			ExposeHeaders: []string{

				"Content-Length",
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

	protected.GET(
		"/dashboard",
		dashboardHandler.Get,
	)

	protected.PUT(
		"/customers/me",
		middleware.RequireWriteAccess(),
		customerHandler.Update,
	)

	// ----------------------------
	// Vehicle Routes
	// ----------------------------
	vehicles := protected.Group("/vehicles")

	vehicles.POST(
		"",
		middleware.RequireWriteAccess(),
		vehicleHandler.Create,
	)

	vehicles.GET(
		"",
		vehicleHandler.List,
	)

	vehicles.DELETE(
		"/:id",
		middleware.RequireWriteAccess(),
		vehicleHandler.Delete,
	)

	vehicles.PATCH(
		"/:id/primary",
		middleware.RequireWriteAccess(),
		vehicleHandler.SetPrimary,
	)

	vehicles.PUT(
		"/:id",
		middleware.RequireWriteAccess(),
		vehicleHandler.Update,
	)

	// ----------------------------
	// Parts Search Routes
	// ----------------------------
	partsRoutes := protected.Group("/parts")
	{
		partsRoutes.GET(
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

	admin.GET(
		"/overview",
		adminHandler.Overview,
	)

	admin.PUT(
		"/users/:id/role",
		adminHandler.UpdateUserRole,
	)

	admin.PUT(
		"/vehicles/:id",
		adminHandler.UpdateVehicle,
	)

	admin.DELETE(
		"/vehicles/:id",
		adminHandler.DeleteVehicle,
	)

	admin.PUT(
		"/customers/:id",
		adminHandler.UpdateCustomer,
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
