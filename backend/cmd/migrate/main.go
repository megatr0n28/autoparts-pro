package main

import (
	"log"

	"github.com/megatr0n28/autoparts-pro/backend/internal/config"

	"github.com/megatr0n28/autoparts-pro/backend/internal/database"
)

func main() {

	cfg, err := config.Load()

	if err != nil {
		log.Fatal(err)
	}

	err = database.RunMigrations(
		cfg.Database,
	)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(
		"database migration completed",
	)

}
