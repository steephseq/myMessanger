package config

import (
	"github.com/caarlos0/env"
)

type Config struct {
	DB struct {
		Host     string `env:"DB_HOST" default:"localhost"`
		Port     string `env:"DB_PORT" default:"5432"`
		User     string `env:"DB_USER" required:"true"`
		Password string `env:"DB_PASSWORD" required:"true"`
		Name     string `env:"DB_NAME" required:"true"`
	}
	JWT struct {
		Secret string `env:"JWT_SECRET" required:"true"`
		Expire string `env:"JWT_EXPIRE" default:"1"`
	}
	Server struct {
		Port string `env:"SERVER_PORT" default:":50051"`
	}
	UserServiceAddr string `env:"USER_SERVICE_ADDR"`
}

func LoadConfig() (*Config, error) {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
