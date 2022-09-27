package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"log"
)

type Config struct {
	Token       string `env:"TELEGRAM_TOKEN"`
	ServiceName string `env:"SERVICE_NAME"`
	LogServer   string `env:"LOG_SERVER"`
	LogLevel    string `env:"LOG_LEVEL"`
}

// NewConfig parses envs and constructs the config
func NewConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var config Config

	if err := env.Parse(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
