package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

type Artist struct {
	id   int
	name string
}

var db *sql.DB

const dbPath = "../database/illa.db"
const driverName = "sqlite3"

func handleGet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

}

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
