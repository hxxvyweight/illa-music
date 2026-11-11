package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "modernc.org/sqlite"
)

var db *sql.DB

type Artist struct {
	ID   int    `json:"ID"`
	Name string `json:"Name"`
	Slug string `json:"Slug"`
}

type Genre struct {
	ID   int    `json:"ID"`
	Name string `json:"Name"`
	Slug string `json:"Slug"`
}

type Venue struct {
	ID       int    `json:"ID"`
	Name     string `json:"Name"`
	Location string `json:"Location"`
	Slug     string `json:"Slug"`
}
type Label struct {
	ID       int    `json:"ID"`
	Name     string `json:"Name"`
	Location string `json:"Location"`
	Slug     string `json:"Slug"`
}

func handleArtists(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	slug := r.URL.Query().Get("slug")
	if slug != "" {
		var artist Artist
		err := db.QueryRow("SELECT id, name, slug FROM artists WHERE slug = ?", slug).Scan(&artist.ID, &artist.Name, &artist.Slug)
		if err == sql.ErrNoRows {
			http.Error(w, "Artist not found", http.StatusNotFound)
			return
		} else if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(artist)
		return
	}

	rows, err := db.Query("SELECT id, name, slug FROM artists")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var artists []Artist
	for rows.Next() {
		var a Artist
		rows.Scan(&a.ID, &a.Name, &a.Slug)
		artists = append(artists, a)
	}
	json.NewEncoder(w).Encode(artists)
}

func handleGenres(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	slug := r.URL.Query().Get("slug")
	if slug != "" {
		var genre Genre
		err := db.QueryRow("SELECT id, name, slug FROM genres WHERE slug = ?", slug).Scan(&genre.ID, &genre.Name, &genre.Slug)
		if err == sql.ErrNoRows {
			http.Error(w, "Genre not found", http.StatusNotFound)
			return
		} else if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(genre)
		return
	}

	rows, err := db.Query("SELECT id, name, slug FROM genres")
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var genres []Genre
	for rows.Next() {
		var g Genre
		rows.Scan(&g.ID, &g.Name, &g.Slug)
		genres = append(genres, g)
	}
	json.NewEncoder(w).Encode(genres)
}

func handleVenues(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	slug := r.URL.Query().Get("slug")
	if slug != "" {
		var venue Venue
		err := db.QueryRow("SELECT id, name, location, slug FROM venues WHERE slug = ?", slug).Scan(&venue.ID, &venue.Name, &venue.Location, &venue.Slug)
		if err == sql.ErrNoRows {
			http.Error(w, "Venue not found", http.StatusNotFound)
			return
		} else if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(venue)
		return
	}

	rows, err := db.Query("SELECT id, name, location, slug FROM venues")
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var venues []Venue
	for rows.Next() {
		var v Venue
		rows.Scan(&v.ID, &v.Name, &v.Location, &v.Slug)
		venues = append(venues, v)
	}
	json.NewEncoder(w).Encode(venues)
}
func handleLabels(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	slug := r.URL.Query().Get("slug")
	if slug != "" {
		var label Label
		err := db.QueryRow("SELECT name, location, slug FROM labels WHERE slug = ?", slug).Scan(&label.Name, &label.Location, &label.Slug)
		if err == sql.ErrNoRows {
			http.Error(w, "Label not found", http.StatusNotFound)
			return
		} else if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(label)
		return
	}

	rows, err := db.Query("SELECT name, location, slug FROM labels")
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var labels []Label
	for rows.Next() {
		var l Label
		rows.Scan(&l.Name, &l.Location, &l.Slug)
		labels = append(labels, l)
	}
	json.NewEncoder(w).Encode(labels)
}

func main() {
	driverName := "sqlite"
	dbPath := "../database/illa.db"

	var err error
	db, err = sql.Open(driverName, dbPath)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	http.HandleFunc("/artists", handleArtists)
	http.HandleFunc("/genres", handleGenres)
	http.HandleFunc("/venues", handleVenues)
	http.HandleFunc("/labels", handleLabels)

	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
