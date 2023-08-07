package config

import "os"

type Config struct {
	PgUrl string
}

func New() *Config {
	pgUrl := os.Getenv("PG_URL")
	if pgUrl == "" {
		pgUrl = "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
	}
	return &Config{
		PgUrl: pgUrl,
	}
}
