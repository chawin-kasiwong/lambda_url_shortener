package config

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	_ = godotenv.Load()
}

func Get(key, fallback string) string {
	if values, exists := os.LookupEnv(key); exists {
		return values
	}

	return fallback
}
