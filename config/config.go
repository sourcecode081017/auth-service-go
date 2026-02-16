package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	App   AppConfig
	Mongo MongoConfig
}

type AppConfig struct {
	HostName string
	Port     int
}

type MongoConfig struct {
	URI      string
	Database string
}

func Load() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("read config: %w", err)
	}
	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unmarshal config: %w", err)
	}
	return &cfg, nil
}
