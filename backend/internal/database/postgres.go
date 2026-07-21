package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/megatr0n28/autoparts-pro/backend/internal/config"
)

func Connect(cfg config.DatabaseConfig) (*gorm.DB, error) {

	dsn := fmt.Sprintf(

		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",

		cfg.Host,

		cfg.Port,

		cfg.User,

		cfg.Password,

		cfg.Name,

		cfg.SSLMode,
	)

	db, err := gorm.Open(

		postgres.Open(dsn),

		&gorm.Config{},
	)

	if err != nil {

		return nil, err
	}

	sqlDB, err := db.DB()

	if err != nil {

		return nil, err
	}

	sqlDB.SetMaxIdleConns(
		cfg.MaxIdleConns,
	)

	sqlDB.SetMaxOpenConns(
		cfg.MaxOpenConns,
	)

	sqlDB.SetConnMaxLifetime(
		cfg.MaxLifetime,
	)

	if err := sqlDB.Ping(); err != nil {

		return nil, err
	}

	return db, nil

}

func AutoMigrate(
	db *gorm.DB,
	models ...interface{},
) error {

	return db.AutoMigrate(models...)

}
