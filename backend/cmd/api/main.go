package main

import (
	"log"

	"github.com/megatr0n28/autoparts-pro/backend/internal/bootstrap"
)

func main() {

	app, err := bootstrap.New()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = app.Logger.Sync()
	}()

	app.Logger.Info(
		"AutoParts Pro API started",
	)
}
