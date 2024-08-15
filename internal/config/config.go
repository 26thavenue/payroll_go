package config

import (
	"os"
)

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}


var (
	JWT_SECRET   = getEnv("JWT_SECRET", "MY_SECRET_JWT_KEY")
	DB_HOST      = getEnv("DB_HOST", "localhost")
	DB_USER      = getEnv("DB_USER", "user")
	DB_PASSWORD  = getEnv("DB_PASSWORD", "password")
	DB_NAME      = getEnv("DB_NAME", "database")
	DB_PORT      = getEnv("DB_PORT", "5432")
)



