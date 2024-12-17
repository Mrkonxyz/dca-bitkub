package config

import (
	"log"

	"github.com/spf13/viper"
)

var AppConfig Config

type Config struct {
	Secret      string `mapstructure:"SECRET"`
	ApiKey      string `mapstructure:"API_KEY"`
	ApiSecret   string `mapstructure:"API_SECRET"`
	BaseUrl     string `mapstructure:"BASE_URL"`
	DiscordHook string `mapstructure:"DISCORD_HOOK"`
}

func LoadConfig(path string) (config Config) {
	viper.SetDefault("API_KEY", "")
	viper.SetDefault("API_SECRET", "")
	viper.SetDefault("BASE_URL", "")
	viper.SetDefault("SECRET", "")
	viper.SetDefault("DISCORD_HOOK", "")

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
		log.Fatal("Error: Missing API_KEY environment variable")
	}

	if cfg.ApiSecret == "" {
		log.Fatal("Error: Missing API_SECRET environment variable")
	}

	if cfg.BaseUrl == "" {
		log.Fatal("Error: Missing BASE_URL environment variable")
	}

	if cfg.DiscordHook == "" {
		log.Fatal("Error: Missing DISCORD_HOOK environment variable")
	}

	if cfg.Secret == "" {
		log.Fatal("Error: Missing SECRET environment variable")
	}
}
