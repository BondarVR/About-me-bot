package config

import (
	"github.com/caarlos0/env/v6"
	"os"
)

type Config struct {
	Token                     string      `env:"TELEGRAM_TOKEN"`
	DirectoryLog              string      `env:"DIRECTORY_LOG"`
	PathLogFile               string      `env:"PATH_LOG_FILE"`
	PermissionForLogFile      os.FileMode `env:"REMISSION_FOR_LOG_FILE"`
	PermissionForLogDirectory os.FileMode `env:"REMISSION_FOR_LOG_DIRECTORY"`
}

// NewConfig parses envs and constructs the config
func NewConfig() (*Config, error) {
	var config Config

	if err := env.Parse(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
