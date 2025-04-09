package bootstrap

import (
	"seth16888/api/gw/pkg/logger"

	"go.uber.org/zap"
)

// InitLogger 初始化日志
func InitLogger(conf *logger.LogConfig) *zap.Logger {
	return logger.InitLogger(conf)
}
