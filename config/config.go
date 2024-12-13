package config

import (
	"log"

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

func (cfg *Config) Validate() {
	if cfg.ApiKey == "" {
		log.Fatal("Error: Missing ApiKey environment variable")
	}

	if cfg.ApiSecret == "" {
		log.Fatal("Error: Missing ApiSecret environment variable")
	}

	if cfg.BaseUrl == "" {
		log.Fatal("Error: Missing BaseUrl environment variable")
	}
}
