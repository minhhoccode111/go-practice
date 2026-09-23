package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Addr   string
	DBPath string
}

func Load() (*Config, error) {
	_ = godotenv.Load() // optional file; read env vars win

	cfg := &Config{
		Addr:   getenv("ADDR", ":8080"),
		DBPath: getenv("DB_PATH", "data/session-auth.db"),
	}
	return cfg, nil
}

func getenv(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}
