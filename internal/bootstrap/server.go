package bootstrap

import (
	"seth16888/api/gw/internal/config"
	"seth16888/api/gw/internal/server"
	"seth16888/api/gw/internal/server/router"

	"go.uber.org/zap"
)

func InitServer(conf config.Conf, log *zap.Logger) *server.Server {
	r := router.NewRouter(log)

	return server.NewServer(conf.Server, log, r)
}
