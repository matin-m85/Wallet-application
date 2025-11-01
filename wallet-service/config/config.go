package config

import (
	"fmt"
	"os"
)

type Config struct {
	DBHost string
	DBUser string
	DBPass string
	DBName string
	DBPort string
}

func LoadConfig() *Config {
	return &Config{
		DBHost: getEnv("DB_HOST", "localhost"),
		DBUser: getEnv("DB_USER", "wallet"),
		DBPass: getEnv("DB_PASSWORD", "wallet"),
		DBName: getEnv("DB_NAME", "wallet_db"),
		DBPort: getEnv("DB_PORT", "5432"),
	}
}

func (c *Config) GetDSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		c.DBHost, c.DBUser, c.DBPass, c.DBName, c.DBPort)
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
