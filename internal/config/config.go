package config

import (
	"os"
	"time"
)

type ServerConfig struct {
	Host string
	Port string
	Env  string
}

type DatabaseConfig struct {
	Driver           string
	ConnectionString string
	MaxOpenConns     int
	MaxIdleConns     int
	MaxLifeTime      time.Duration
}

type Config struct {
	Database DatabaseConfig
	Server   ServerConfig
}

func getEnv(key string, fallback string) string {
	value := os.Getenv(key)

	if value == "" {
		return fallback
	}

	return value
}

func Load() Config {
	return Config{
		Server: ServerConfig{
			Host: getEnv("GO_HOST", "localhost"),
			Port: getEnv("GO_PORT", "8888"),
			Env:  getEnv("GO_ENV", "development"),
		},
		Database: DatabaseConfig{
			Driver:           getEnv("GO_DB_DRIVER", "sqlite3"),
			ConnectionString: getEnv("GO_DB_CONNECTION_STRING", "data/database.db"),
			MaxOpenConns:     1,
			MaxIdleConns:     1,
			MaxLifeTime:      time.Hour,
		},
	}
}
