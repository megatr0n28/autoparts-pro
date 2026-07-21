package logger

import (
	"go.uber.org/zap"
)

func New(level string) (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()

	if level == "debug" {
		cfg = zap.NewDevelopmentConfig()
	}

	return cfg.Build()
}
