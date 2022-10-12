package utils

import (
	"github.com/darkcat013/pr-client/config"
	"go.uber.org/zap"
)

var Log *zap.Logger

func InitializeLogger() {
	if config.LOGS_ENABLED {
		Log, _ = zap.NewDevelopment()
	} else {
		Log = zap.NewNop()
	}
}
