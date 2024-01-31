package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Host Host
}

type Host struct {
	Uri  string
	Port int
}

func LoadConfig(path, env string) (config Config, err error) {
	viper.SetConfigName(fmt.Sprintf("appsettings.%s", env))
	viper.AddConfigPath(path)
	viper.SetConfigType("json")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return
	}

	return
}
