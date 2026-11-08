package main

import (
	"database/sql"
	"log"
	"os"

	_ "modernc.org/sqlite"
)

func createTable() {
	path := os.Getenv("DB_PATH")
	if path == "" {
		path = "database.db"
	}

	var err error

	db, err = sql.Open("sqlite", path)

	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`--sql
	create table if not exists clients(
		id         INTEGER PRIMARY KEY AUTOINCREMENT,
		name_enc   TEXT NOT NULL,
		phone_enc  TEXT NOT NULL,
		email_enc  TEXT NOT NULL,	
		job_enc    TEXT NOT NULL,
		clinic_enc TEXT NOT NULL
	)`)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("table have created")
}