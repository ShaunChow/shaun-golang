package infrastructure

import (
	"go.uber.org/zap"
)

var Logger *zap.Logger

func InitLogger() {
	logger, _ := zap.NewProduction()

	Logger = logger

	defer logger.Sync() // flushes buffer, if any
}
