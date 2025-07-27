package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	ServerPort  string `envconfig:"SERVER_PORT" default:"8081"`
	DatabaseURL string `envconfig:"DATABASE_URL" required:"true"`
}

func Load() (*Config, error) {
	var cfg Config

	if err := envconfig.Process("user_service", &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
