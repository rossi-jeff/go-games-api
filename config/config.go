package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DbConfig struct {
	DbUser string
	DbPass string
	DbHost string
	DbName string
	DbPort string
}

type Config struct {
	Database DbConfig
}

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func New() *Config {
	return &Config{
		Database: DbConfig{
			DbUser: getEnv("DB_USER", ""),
			DbPass: getEnv("DB_PASS", ""),
			DbHost: getEnv("DB_HOST", ""),
			DbName: getEnv("DB_NAME", ""),
			DbPort: getEnv("DB_PORT", ""),
		},
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
