package database

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/megatr0n28/autoparts-pro/backend/internal/config"
)

func RunMigrations(
	cfg config.DatabaseConfig,
) error {

	dsn := fmt.Sprintf(

		"postgres://%s:%s@%s:%d/%s?sslmode=%s",

		cfg.User,

		cfg.Password,

		cfg.Host,

		cfg.Port,

		cfg.Name,

		cfg.SSLMode,
	)

	m, err := migrate.New(

		"file://migrations",

		dsn,
	)

	if err != nil {
		return err
	}

	err = m.Up()

	if err != nil &&
		err != migrate.ErrNoChange {

		return err
	}

	return nil
}
