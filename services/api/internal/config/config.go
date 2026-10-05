package config

import (
	"os"
)

type Config struct {
	Port        string
	Environment string
	DatabaseURL string
}

func Load() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	env := os.Getenv("ENV")
	if env == "" {
		env = "development"
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "postgres://postgres:postgres@localhost:5432/pos_saas?sslmode=disable"
	}

	return &Config{
		Port:        port,
		Environment: env,
		DatabaseURL: dbURL,
	}
}
