package config

import (
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	JaegerHost string `mapstructure:"JAEGER_HOST"`
	ServerPort string `mapstructure:"SERVER_PORT"`
	AppName    string `mapstructure:"APP_NAME"`
}

func LoadConfig() (err error) {
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" {
		appEnv = "dev"
	}

	viper.SetConfigName(appEnv)
	viper.AddConfigPath(".")
	viper.SetConfigType("env")

	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	return
}

func GetConfig() (config Config, err error) {
	err = viper.Unmarshal(&config)
	return
}
