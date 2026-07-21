package bootstrap

import (
	"github.com/megatr0n28/autoparts-pro/backend/internal/config"
	"github.com/megatr0n28/autoparts-pro/backend/internal/database"
	"github.com/megatr0n28/autoparts-pro/backend/internal/domain/customer"
	"github.com/megatr0n28/autoparts-pro/backend/internal/logger"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Application struct {
	Config *config.Config
	Logger *zap.Logger
	DB     *gorm.DB
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

	err = database.AutoMigrate(
		db,
		&customer.Customer{},
	)

	if err != nil {
		return nil, err
	}

	app := &Application{
		Config: cfg,
		Logger: log,
		DB:     db,
	}

	return app, nil
}
