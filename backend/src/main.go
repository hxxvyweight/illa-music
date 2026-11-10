package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "modernc.org/sqlite"
)

type Artist struct {
	ID   int
	Name string
}

var db *sql.DB

const dbPath = "../database/illa.db"
const driverName = "sqlite"

func handleArtists(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	rows, err := db.Query("SELECT id, name FROM artists")
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var artists []Artist

	for rows.Next() {

		var artist Artist

		err := rows.Scan(&artist.ID, &artist.Name)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		artists = append(artists, artist)

	}

	if err = rows.Err(); err != nil {
		log.Printf("ERROR: Failed to fetch artist: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(artists)

}

func main() {
	var err error

	db, err = sql.Open(driverName, dbPath)
	if err != nil {
		log.Fatalln("Unable to open database", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println("Database OK")

	http.HandleFunc("/artists", handleArtists)

	http.ListenAndServe(":8080", nil)

}
