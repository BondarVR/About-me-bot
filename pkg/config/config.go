package config

import "github.com/caarlos0/env/v6"

type Config struct {
	Token string `env:"TELEGRAM_TOKEN" envDefault:"5450445453:AAHcCCBP-Gpg5yDq_ulecnqQl4tyAP_RDzQ"`
}

// NewConfig parses envs and constructs the config
func NewConfig() (*Config, error) {
	var config Config

	if err := env.Parse(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
