package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
)

type Config struct {
	Secret       string `mapstructure:"SECRET"`
	ApiKey       string `mapstructure:"API_KEY"`
	ApiSecret    string `mapstructure:"API_SECRET"`
	BaseUrl      string `mapstructure:"BASE_URL"`
	DiscordHook  string `mapstructure:"DISCORD_HOOK"`
	MongoUrl     string `mapstructure:"MONGO_URL"`
	DatabaseName string `mapstructure:"DATABASE_NAME"`
	Port         string `mapstructure:"PORT"`
	DB           *mongo.Database
}

func LoadConfig(path string) Config {

	viper.SetDefault("API_KEY", "")
	viper.SetDefault("API_SECRET", "")
	viper.SetDefault("BASE_URL", "")
	viper.SetDefault("SECRET", "")
	viper.SetDefault("DISCORD_HOOK", "")
	viper.SetDefault("MONGO_URL", "")
	viper.SetDefault("DATABASE_NAME", "")
	viper.SetDefault("PORT", "8080")
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Warning: cannot read config file: %v", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatal("Unable to decode config:", err)
	}

	if err := cfg.Validate(); err != nil {
		log.Fatal("Config validation error:", err)
	}

	client, err := cfg.ConnectMongoDB()
	if err != nil {
		log.Fatal("Mongo connection error:", err)
	}
	cfg.DB = client.Database(cfg.DatabaseName)

	return cfg
}

func (cfg *Config) Validate() error {
	if cfg.ApiKey == "" {
		return fmt.Errorf("missing API_KEY environment variable")
	}
	if cfg.ApiSecret == "" {
		return fmt.Errorf("missing API_SECRET environment variable")
	}
	if cfg.BaseUrl == "" {
		return fmt.Errorf("missing BASE_URL environment variable")
	}
	if cfg.DiscordHook == "" {
		return fmt.Errorf("missing DISCORD_HOOK environment variable")
	}
	if cfg.Secret == "" {
		return fmt.Errorf("missing SECRET environment variable")
	}
	if cfg.MongoUrl == "" {
		return fmt.Errorf("missing MONGO_URL environment variable")
	}
	if cfg.DatabaseName == "" {
		return fmt.Errorf("missing DATABASE_NAME environment variable")
	}
	return nil
}
