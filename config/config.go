package config

import (
	"fmt"

	"github.com/caarlos0/env/v6"
	"github.com/labstack/gommon/log"
)

const (
	DefaultPort = 8080
	LocalDSN    = "postgres://postgres:@localhost:5432/postgres?sslmode=disable"
)

type DB struct {
	DSN string `env:"DB_DSN"`
}

type Server struct {
	Port     int     `env:"PORT"`
	LogLevel log.Lvl `env:"SERVER_LOG_LEVEL"`
}

type Config struct {
	Server Server
	DB     DB
}

func New() *Config {
	return &Config{
		Server: Server{
			Port: DefaultPort,
		},
		DB: DB{
			DSN: LocalDSN,
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
