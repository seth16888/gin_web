package bootstrap

import "seth16888/api/gw/internal/config"

func InitConfig(configFile string) (conf *config.Conf, err error) {
	return config.ReadConfigFromFile(configFile), nil
}
