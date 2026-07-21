package main

import (
	"log"

	"github.com/megatr0n28/autoparts-pro/backend/internal/config"
	"github.com/megatr0n28/autoparts-pro/backend/internal/database"
	"go.uber.org/zap"

	"github.com/megatr0n28/autoparts-pro/backend/internal/logger"
)

func main() {

	cfg, err :=
		config.Load()

	if err != nil {

		log.Fatal(err)

	}

	err = logger.Initialize(cfg.Log.Level)
	if err != nil {
		log.Fatal(err)
	}

	defer logger.Sync()

	logger.Log.Info(
		"AutoParts Pro starting",
	)

	logger.Log.Info(
		"Configuration loaded",
		zap.String("environment", cfg.App.Environment),
	)

	err = database.Connect(
		cfg.Database,
	)

	if err != nil {

		logger.Log.Fatal(
			"database connection failed",
		)
	}

	logger.Log.Info(
		"database connected",
	)

}
