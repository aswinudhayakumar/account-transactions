package logger

import (
	"go.uber.org/zap"
)

var Log *zap.Logger

// InitLogger initialises the global logger
func InitLogger() error {
	var err error
	cfg := zap.NewProductionConfig()

	cfg.DisableStacktrace = true

	Log, err = cfg.Build()
	if err != nil {
		return err
	}
	return nil
}

func SyncLogger() {
	if Log != nil {
		_ = Log.Sync()
	}
}
