package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Key string
}

func NewCongig() (*Config, error) {
	godotenv.Load()
	key := os.Getenv("KEY")
	if key == "" {
		return nil, errors.New("Не_могу_загрузить_переменную_окружения_KEY")
	}
	return &Config{
		Key: key,
	}, nil
}
