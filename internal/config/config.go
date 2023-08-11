package config

import (
	"fmt"
	"os"
)

type Config struct {
	PgUrl string
}

func New() *Config {
	pgUrl := dbConnectionUrl()
	return &Config{
		PgUrl: pgUrl,
	}
}

func dbConnectionUrl() string {
	connectionName := os.Getenv("DB_CONNECTION_NAME")
	if connectionName == "" {
		return "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
	}

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")

	return fmt.Sprintf(
		"postgres://%s:%s@unix(%s)/%s?parseTime=true",
		user,
		password,
		connectionName,
		name,
	)
}
