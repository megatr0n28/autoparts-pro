package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	App AppConfig

	Log LogConfig

	Database DatabaseConfig

	Redis RedisConfig

	JWT JWTConfig
}

type AppConfig struct {
	Name string

	Environment string

	Port int
}

type LogConfig struct {
	Level string
}

type DatabaseConfig struct {
	Host string

	Port int

	Name string

	User string

	Password string

	SSLMode string

	MaxIdleConns int

	MaxOpenConns int

	MaxLifetime time.Duration
}

type RedisConfig struct {
	Host string

	Port int
}

type JWTConfig struct {
	Expiration time.Duration
}

func Load() (*Config, error) {

	v := viper.New()

	v.SetConfigName("development")

	v.SetConfigType("yaml")

	v.AddConfigPath("./configs")

	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {

		return nil, fmt.Errorf(
			"failed loading config: %w",
			err,
		)

	}

	expiration, err :=
		time.ParseDuration(
			v.GetString(
				"jwt.expiration",
			),
		)

	if err != nil {

		expiration = 24 * time.Hour

	}

	config := &Config{

		App: AppConfig{

			Name: v.GetString(
				"app.name",
			),

			Environment: v.GetString(
				"app.environment",
			),

			Port: v.GetInt(
				"app.port",
			),
		},

		Log: LogConfig{

			Level: v.GetString(
				"log.level",
			),
		},

		Database: DatabaseConfig{
			Host:         v.GetString("database.host"),
			Port:         v.GetInt("database.port"),
			Name:         v.GetString("database.name"),
			User:         v.GetString("database.user"),
			Password:     v.GetString("database.password"),
			SSLMode:      v.GetString("database.sslmode"),
			MaxIdleConns: v.GetInt("database.max_idle_conns"),
			MaxOpenConns: v.GetInt("database.max_open_conns"),
			MaxLifetime: time.Minute * time.Duration(
				v.GetInt("database.max_lifetime"),
			),
		},

		Redis: RedisConfig{

			Host: v.GetString(
				"redis.host",
			),

			Port: v.GetInt(
				"redis.port",
			),
		},

		JWT: JWTConfig{

			Expiration: expiration,
		},
	}

	return config, nil
}
