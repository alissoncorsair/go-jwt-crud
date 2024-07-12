package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type PostgreConfig struct {
	Host                   string
	Port                   int
	User                   string
	Password               string
	DBName                 string
	SSLMode                string
	JWTExpirationInSeconds int64
	JWTSecret              string
}

var Envs = newPostgreConfig()

func newPostgreConfig() *PostgreConfig {
	godotenv.Load()
	port, err := strconv.Atoi(getEnv("PORT", "5432"))

	if err != nil {
		log.Fatal(err)
	}

	return &PostgreConfig{
		Host:                   getEnv("DB_HOST", "localhost"),
		Port:                   port,
		User:                   getEnv("DB_USER", "postgres"),
		Password:               getEnv("DB_PASSWORD", "1234"),
		DBName:                 getEnv("DB_NAME", "postgres"),
		SSLMode:                "disable",
		JWTExpirationInSeconds: getEnvAsInt64("JWT_EXPIRATION_IN_SECONDS", 3600),
		JWTSecret:              getEnv("JWT_SECRET", "secret"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func getEnvAsInt64(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}
		return i
	}

	return fallback
}
