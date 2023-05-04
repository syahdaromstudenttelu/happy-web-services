package util

import (
	"happy-api-service/helper"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	HappyBaseUrlService     string `mapstructure:"HAPPY_BASE_URL_SERVICE"`
	HappyUserServiceUrl     string `mapstructure:"HAPPY_USER_SERVICE_URL"`
	HappyProductServiceUrl  string `mapstructure:"HAPPY_PRODUCT_SERVICE_URL"`
	HappyFeedbackServiceUrl string `mapstructure:"HAPPY_FEEDBACK_SERVICE_URL"`
	HappyOrderServiceUrl    string `mapstructure:"HAPPY_ORDER_SERVICE_URL"`
	JwtSecretKey            string `mapstructure:"JWT_SECRET_KEY"`
	AllowOrigins            string `mapstructure:"ALLOW_ORIGINS"`
	ServerAddr              string `mapstructure:"SERVER_ADDR"`
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
