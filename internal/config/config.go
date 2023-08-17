package config

import (
	"fmt"
	"os"
)

type Config struct {
	PgUrl      string
	BucketName string
	CdnBaseUrl string
}

func New() *Config {
	pgUrl := dbConnectionUrl()
	bucketName := bucketName()
	cdnBaseUrl := cdnBaseUrl()

	return &Config{
		PgUrl:      pgUrl,
		BucketName: bucketName,
		CdnBaseUrl: cdnBaseUrl,
	}
}

func cdnBaseUrl() string {
	cdnBaseUrl := os.Getenv("CDN_BASE_URL")
	if cdnBaseUrl == "" {
		return "http://localhost:8080"
	}
	return cdnBaseUrl
}

func bucketName() string {
	bucketName := os.Getenv("BUCKET_NAME")
	if bucketName == "" {
		return "just-code-dev-bucket"
	}
	return bucketName
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
		"user=%s password=%s database=%s host=%s",
		user,
		password,
		name,
		connectionName,
	)
}
