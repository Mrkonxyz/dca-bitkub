package config

import (
	"github.com/spf13/viper"
)

var AppConfig Config

type Config struct {
	ApiKey    string `mapstructure:"API_KEY"`
	ApiSecret string `mapstructure:"API_SECRET"`
	BaseUrl   string `mapstructure:"BASE_URL"`
}

func LoadConfig(path string) (config Config) {
	viper.SetDefault("API_KEY", "")
	viper.SetDefault("API_SECRET", "")
	viper.SetDefault("BASE_URL", "")

	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	viper.ReadInConfig()
	viper.Unmarshal(&config)
	AppConfig = config
	return AppConfig
}
