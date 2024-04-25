package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func ConnectDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./assets/cryptocurrencies_tracker.sqlite3")
	if err != nil {
		log.Fatalf("Can't connect to database with err: %s", err)
		panic(err)
	}

	// defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	log.Println("Success Connect To Database")
	return db

}
