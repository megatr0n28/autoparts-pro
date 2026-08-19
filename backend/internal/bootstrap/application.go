package bootstrap

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/megatr0n28/autoparts-pro/backend/internal/auth"
	"github.com/megatr0n28/autoparts-pro/backend/internal/authz"
	"github.com/megatr0n28/autoparts-pro/backend/internal/config"
	"github.com/megatr0n28/autoparts-pro/backend/internal/database"
	"github.com/megatr0n28/autoparts-pro/backend/internal/handler"
	"github.com/megatr0n28/autoparts-pro/backend/internal/logger"
	"github.com/megatr0n28/autoparts-pro/backend/internal/provider/mock"
	"github.com/megatr0n28/autoparts-pro/backend/internal/router"
	"github.com/megatr0n28/autoparts-pro/backend/internal/service"
	"github.com/megatr0n28/autoparts-pro/backend/internal/service/search"
	"github.com/megatr0n28/autoparts-pro/backend/internal/service/vehicle"
	"go.uber.org/zap"
)

type Application struct {
	Config         *config.Config
	Logger         *zap.Logger
	Repositories   *Repositories
	Router         *gin.Engine
	vehicleHandler *handler.VehicleHandler
	searchHandler  *handler.SearchHandler
	authorizer     *authz.Authorizer
}

func New() (*Application, error) {

	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}

	log, err := logger.New(cfg.Log.Level)
	if err != nil {
		return nil, err
	}

	db, err := database.Connect(cfg.Database)
	if err != nil {
		return nil, err
	}

	if err := database.RunMigrations(cfg.Database); err != nil {
		return nil, err
	}

	repositories := NewRepositories(db)

	customerRepository := repositories.Customer

	// ----------------------------
	// Authentication
	// ----------------------------
	jwtManager := auth.NewJWTManager(
		cfg.JWT.Secret,
		cfg.JWT.Expiration,
	)

	refreshService :=
		auth.NewRefreshTokenService(
			repositories.RefreshTokens,
			jwtManager,
			cfg.JWT.RefreshExpiration,
		)

	openFGAClient, err := authz.New(
		authz.Config{
			APIURL:               cfg.OpenFGA.APIURL,
			StoreID:              cfg.OpenFGA.StoreID,
			AuthorizationModelID: cfg.OpenFGA.AuthorizationModelID,
		},
	)

	if err != nil {
		return nil, fmt.Errorf(
			"initialize OpenFGA: %w",
			err,
		)
	}
	if err := openFGAClient.CheckConnection(
		context.Background(),
	); err != nil {

		return nil, fmt.Errorf(
			"connect to OpenFGA: %w",
			err,
		)
	}

	authorizer :=
		authz.NewAuthorizer(
			openFGAClient,
		)

	authService :=
		auth.NewService(
			repositories.User,
			jwtManager,
			refreshService,
			repositories.Customer,
			authorizer,
		)

	authHandler := handler.NewAuthHandler(
		authService,
	)

	// ----------------------------
	// Customer Service
	// ----------------------------
	customerService :=
		service.NewCustomerService(
			repositories.Customer,
		)

	customerHandler :=
		handler.NewCustomerHandler(
			customerService,
		)

	// ----------------------------
	// Vehicles
	// ----------------------------
	vehicleService :=
		vehicle.NewService(
			repositories.Vehicle,
			authorizer,
		)

	vehicleHandler :=
		handler.NewVehicleHandler(
			vehicleService,
		)

	dashboardService :=
		service.NewDashboardService(
			customerService,
			vehicleService,
			authorizer,
		)

	dashboardHandler :=
		handler.NewDashboardHandler(
			dashboardService,
		)

	adminService :=
		service.NewAdminService(
			repositories.User,
			repositories.Customer,
			repositories.Vehicle,
		)

	adminHandler :=
		handler.NewAdminHandler(
			adminService,
		)

	// ----------------------------
	// Search Providers
	// ----------------------------
	mockProvider :=
		mock.New()

	searchSvc :=
		search.New(
			log,
			mockProvider,
		)

	searchHandler :=
		handler.NewSearchHandler(
			searchSvc,
		)

	// ----------------------------
	// Users
	// ----------------------------
	userHandler := handler.NewUserHandler()

	// ----------------------------
	// Router
	// ----------------------------
	appRouter := router.New(
		jwtManager,
		userHandler,
		authHandler,
		customerHandler,
		vehicleHandler,
		dashboardHandler,
		adminHandler,
		customerRepository,
		searchHandler,
	)

	app := &Application{
		Config:         cfg,
		Logger:         log,
		Repositories:   repositories,
		Router:         appRouter,
		vehicleHandler: vehicleHandler,
		searchHandler:  searchHandler,
		authorizer:     authorizer,
	}

	return app, nil
}
