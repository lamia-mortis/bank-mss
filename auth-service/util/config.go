package util

import (
	"github.com/spf13/viper"
)

type Config struct {
	Environment   string `mapstructure:"ENVIRONMENT"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
	AuthDbUrl     string `mapstructure:"AUTH_DB_URL"`
	DbDriver      string `mapstructure:"DB_DRIVER"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
