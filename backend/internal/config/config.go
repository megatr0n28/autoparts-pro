package config

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	App      AppConfig      `mapstructure:"app"`
	Log      LogConfig      `mapstructure:"log"`
	Database DatabaseConfig `mapstructure:"database"`
	Redis    RedisConfig    `mapstructure:"redis"`
	JWT      JWTConfig      `mapstructure:"jwt"`
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
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	Name         string `mapstructure:"name"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	SSLMode      string `mapstructure:"sslmode"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxLifetime  time.Duration
}

type RedisConfig struct {
	Host string `mapstructure:"host"`

	Port int `mapstructure:"port"`
}

type JWTConfig struct {
	Secret            string `mapstructure:"secret"`
	Expiration        time.Duration
	RefreshExpiration time.Duration `mapstructure:"refresh_expiration"`
}

func Load() (*Config, error) {

	// Load .env if present
	_ = godotenv.Load()

	v := viper.New()

	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	v.SetConfigName(env)
	v.SetConfigType("yaml")

	v.AddConfigPath("./configs")
	v.AddConfigPath("../configs")

	// Environment variables override YAML
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	err := v.ReadInConfig()
	if err != nil {

		return nil, fmt.Errorf(
			"failed loading config: %w",
			err,
		)

	}

	var cfg Config

	if err := v.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	// Parse JWT duration because Unmarshal doesn't convert "24h" into time.Duration automatically.
	cfg.JWT.Expiration, err = time.ParseDuration(v.GetString("jwt.expiration"))
	if err != nil {
		cfg.JWT.Expiration = 24 * time.Hour
	}

	fmt.Printf("Loaded configuration: %s\n", env)

	return &cfg, nil
}
