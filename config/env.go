package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type PostgreConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

var Envs = newPostgreConfig()

func newPostgreConfig() *PostgreConfig {
	godotenv.Load()
	port, err := strconv.Atoi(getEnv("PORT", "5432"))

	if err != nil {
		log.Fatal(err)
	}

	return &PostgreConfig{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     port,
		User:     getEnv("DB_USER", "postgres"),
		Password: getEnv("DB_PASSWORD", "1234"),
		DBName:   getEnv("DB_NAME", "postgres"),
		SSLMode:  "disable",
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
