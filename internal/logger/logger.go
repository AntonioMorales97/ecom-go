package logger

import (
	"go.uber.org/zap"
)

var (
	// Log variable is a globally accessible variable which will be initialized when the InitializeZapCustomLogger function is executed successfully.
	Log *zap.Logger
)

// InitializeZapCustomLogger Funtion initializes a logger using uber-go/zap package in the application.
func InitializeZapCustomLogger(env string) (err error) {

	if env == "dev" {
		Log, err = zap.NewDevelopment()
	} else {
		Log, err = zap.NewProduction()
	}

	defer Log.Sync()
	return
}
