package db

import (
	"database/sql"
	"log"
	"opentalk-service/config"
)

var DB *sql.DB

func ConnectDB() {
	var err error

	DB, err = sql.Open("mysql", config.DSN)
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatal("Failed to ping database: ", err)
	}
}
