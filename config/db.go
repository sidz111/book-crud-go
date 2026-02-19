package config

import (
	"database/sql"
	"log"
)

func ConnectDB() *sql.DB {
	db_url := "root:root@tcp(localhost:3303)/books"

	db, err := sql.Open("mysql", db_url)

	if err != nil {
		log.Fatal("DB Not Found")
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Failed to connect DB")
	}
	return db
}
