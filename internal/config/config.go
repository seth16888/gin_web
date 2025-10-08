package config

import (
	"seth16888/api/gw/internal/server"
	"seth16888/api/gw/pkg/logger"
	"sync"

	"github.com/spf13/viper"
)

type Conf struct {
	Server *server.ServerConf `yaml:"server"`
	Log    *logger.LogConfig  `yaml:"log"`
}

var appConf *Conf
var mutex sync.Mutex

func GetConf() *Conf {
	mutex.Lock()
	defer mutex.Unlock()
	if appConf != nil {
		return appConf
	}
	return ReadConfigFromFile("")
}

func ReadConfigFromFile(file string) *Conf {
	defer func() {
		if r := recover(); r != nil {
			panic("Error reading config file: " + r.(error).Error() + "\n")
		}
	}()

	if file == "" {
		file = "configs/config.yaml"
	}

	viper.SetConfigFile(file)
	viper.AddConfigPath(".")
	viper.AddConfigPath("./configs")
	viper.AddConfigPath("../configs")
	viper.AddConfigPath("~")

	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	appConf = &Conf{}
	if err := viper.Unmarshal(appConf); err != nil {
		panic(err)
	}

	// watch
	viper.WatchConfig()

	return appConf
}
