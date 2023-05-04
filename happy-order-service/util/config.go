package util

import (
	"happy-order-service/helper"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	DBDriver   string `mapstructure:"DB_DRIVER"`
	DBSource   string `mapstructure:"DB_SOURCE"`
	ServerAddr string `mapstructure:"SERVER_ADDR"`
}

func setupConfigFile(path, configName string) error {
	viper.AddConfigPath(path)
	viper.SetConfigName(configName)
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	return err
}

func LoadConfig(path string) (config Config, err error) {
	appEnv := os.Getenv("APP_ENV")

	if appEnv == "prod" {
		err = setupConfigFile(path, "prod")
		helper.DoPanicIfError(err)
	}
	if appEnv == "dev" {
		err = setupConfigFile(path, "dev")
		helper.DoPanicIfError(err)
	}

	err = viper.Unmarshal(&config)
	return
}
