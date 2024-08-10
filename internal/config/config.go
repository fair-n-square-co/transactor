package config

import "os"

type Config struct {
	Database DatabaseConfig
}

type DatabaseConfig struct {
	DSN string
}

// TODO: refactor using some config management library
func GetDatabaseConfig() DatabaseConfig {
	// read from env vars or set default values
	databaseUrl := "postgres://postgres:postgres@localhost:5432/transactions?sslmode=disable"
	envDatabaseUrl := os.Getenv("DATABASE_URL")
	if envDatabaseUrl != "" {
		databaseUrl = envDatabaseUrl
	}
	return DatabaseConfig{
		DSN: databaseUrl,
	}
}

func NewConfig() *Config {
	return &Config{
		Database: GetDatabaseConfig(),
	}
}
