package main

import (
	"database/sql"
	"fmt"
	"log"
)

var db *sql.DB

const dbPath = "../database/illa.db"
const driverName = "sqlite3"

func main() {

	db, err := sql.Open(driverName, dbPath)
	if err != nil {
		log.Fatalln("Unable to open database")
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println("Database OK")

}
