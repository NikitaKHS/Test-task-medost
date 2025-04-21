package config

import (
	"errors"
	"os"
)

type Config struct {
	DatabaseURL string
	JWTSecret   string
}

func Load() (*Config, error) {
	dbURL := os.Getenv("DATABASE_URL")
	secret := os.Getenv("JWT_SECRET")
	if dbURL == "" || secret == "" {
		return nil, errors.New("Ошибка: Нет DATABASE_URL или JWT_SECRET")
	}
	return &Config{DatabaseURL: dbURL, JWTSecret: secret}, nil
}
