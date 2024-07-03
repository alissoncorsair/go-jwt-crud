package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/alissoncorsair/goapi/config"
	_ "github.com/lib/pq"
)

func NewPostgreSQLStorage(config config.PostgreConfig) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode,
	)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}
