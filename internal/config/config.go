package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PORT        string
	DatabaseURL string
}

func LoadConfig() *Config {
	// Load env if exists
	_ = godotenv.Load()

	cfg := &Config{
		PORT:        getEnv("PORT", "8080"),
		DatabaseURL: getEnv("DATABASE_URL", "postgres://user:password@localhost:5432/dbname?sslmode=disable"),
	}

	return cfg
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)

	if value == "" {
		return fallback
	}

	return value
}
