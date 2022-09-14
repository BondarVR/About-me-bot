package config

import "github.com/caarlos0/env/v6"

type Config struct {
	Token string `env:"TELEGRAM_TOKEN"`
}

// NewConfig parses envs and constructs the config
func NewConfig() (*Config, error) {
	var config Config

	if err := env.Parse(&config); err != nil {
		return nil, err
	}

	return &config, nil
}