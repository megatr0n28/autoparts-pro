package bootstrap

import (
	"github.com/megatr0n28/autoparts-pro/backend/internal/config"
	"github.com/megatr0n28/autoparts-pro/backend/internal/database"
	"github.com/megatr0n28/autoparts-pro/backend/internal/logger"
	"go.uber.org/zap"
)

type Application struct {
	Config       *config.Config
	Logger       *zap.Logger
	Repositories *Repositories
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
	repositories := NewRepositories(db)
	err = database.RunMigrations(
		cfg.Database,
	)

	if err != nil {
		return nil, err
	}

	app := &Application{
		Config:       cfg,
		Logger:       log,
		Repositories: repositories,
	}

	return app, nil
}
