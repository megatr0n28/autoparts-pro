package logger

import (
	"go.uber.org/zap"
)

var Log *zap.Logger

func Initialize(
	level string,
) error {

	var err error

	Log, err =
		zap.NewProduction()

	if level == "debug" {

		Log, err =
			zap.NewDevelopment()

	}

	return err

}

func Sync() {

	if Log != nil {

		Log.Sync()

	}

}
