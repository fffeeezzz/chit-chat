package database

import (
	"database/sql"
	"log"
)

var DB *sql.DB

func init() {
	var err error
	DB, err = sql.Open("postgres", "dbname=chit_chat_db sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return
}
