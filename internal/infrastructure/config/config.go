package config

import (
	"log/slog"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type AppConfig struct {
	Server ServerConfig `env-prefix:"SERVER_" yaml:"server"`
}

type ServerConfig struct {
	Port         int    `env:"PORT" env-default:"8080" yaml:"port"`
	CookieSecret string `env:"SECRET_COOKIE_STORE" env-default:"my-secret" yaml:"cookie_secret"`
}

func LoadConfig() (*AppConfig, error) {
	if godotenv.Load() != nil {
		slog.Warn("No .env file found")
	}

	cfg := new(AppConfig)

	if err := cleanenv.ReadEnv(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
