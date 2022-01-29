package config

import (
	"fmt"

	"github.com/caarlos0/env/v6"
	"github.com/labstack/gommon/log"
)

const DefaultPort = 8080

type ServerConfig struct {
	Port     int     `env:"PORT"`
	LogLevel log.Lvl `env:"SERVER_LOG_LEVEL"`
}

type Config struct {
	Server ServerConfig
}

func New() *Config {
	return &Config{
		Server: ServerConfig{
			Port: DefaultPort,
		},
	}
}

func FromEnv() (*Config, error) {
	cfg := New()
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("env parsing: %w", err)
	}
	return cfg, nil
}
