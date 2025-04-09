package di

import (
	"seth16888/api/gw/internal/config"
	"seth16888/api/gw/internal/server"

	"go.uber.org/zap"
)

var (
	di *Container
)

func init() {
	di = new(Container)
}

func Get() *Container {
	if di == nil {
		panic("di not initialized")
	}
	return di
}

type Container struct {
	Conf   *config.Conf
	Log    *zap.Logger
	Server *server.Server
}
