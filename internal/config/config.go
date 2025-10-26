package config

import (
	"fmt"
	"os"
	"strings"
)

// Config holds all configuration values
type Config struct {
	DBDriver   string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPass     string
	DBName     string
	RedisAddr  string
	UseRedis   bool
	ServerPort string
}

// LoadConfig loads configuration from environment variables
func LoadConfig() *Config {
	driver := getEnv("DB_DRIVER", "sqlite") // default local development

	useRedis := strings.ToLower(getEnv("USE_REDIS", "false")) == "true"

	return &Config{
		DBDriver:   driver,
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPass:     getEnv("DB_PASS", "postgres"),
		DBName:     getEnv("DB_NAME", "shortener.db"), // sqlite file or postgres DB
		RedisAddr:  getEnv("REDIS_ADDR", "localhost:6379"),
		UseRedis:   useRedis,
		ServerPort: getEnv("PORT", "8080"),
	}
}

// DSN returns the database connection string
func (c *Config) DSN() string {
	if c.DBDriver == "sqlite" {
		return c.DBName
	}
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.DBHost, c.DBPort, c.DBUser, c.DBPass, c.DBName,
	)
}

// getEnv returns the value of the environment variable or fallback if empty
func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
