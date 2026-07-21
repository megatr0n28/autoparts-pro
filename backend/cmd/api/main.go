package main

import (
	"log"

	"github.com/megatr0n28/autoparts-pro/backend/internal/config"

	"github.com/megatr0n28/autoparts-pro/backend/internal/logger"
)

func main() {

	cfg, err :=
		config.Load()

	if err != nil {

		log.Fatal(err)

	}

	err =
		logger.Initialize(
			cfg.Log.Level,
		)

	if err != nil {

		log.Fatal(err)

	}

	defer logger.Sync()

	logger.Log.Info(
		"AutoParts Pro starting",
	)

	logger.Log.Info(
		"configuration loaded",
	)

}
