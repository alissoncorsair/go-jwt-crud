package main

import (
	"database/sql"
	"log"

	"github.com/alissoncorsair/goapi/cmd/api"
	"github.com/alissoncorsair/goapi/config"
	"github.com/alissoncorsair/goapi/db"
)

func main() {
	dbCfg := config.Envs
	db, err := db.NewPostgreSQLStorage(*dbCfg)

	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)

	server := api.NewAPIServer(":8080", db)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("connected to db")
}
