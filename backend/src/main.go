package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Struct definitions matching your updated schema with slugs
type Artist struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type Genre struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type Venue struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
	Slug     string `json:"slug"`
}

// 1. Handle Artists (Supports fetching all OR a single artist by slug ?slug=...)
func handleArtists(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Check if a specific slug was requested
	slug := r.URL.Query().Get("slug")
	if slug != "" {
		var artist Artist
		query := "SELECT id, name, slug FROM artists WHERE slug = ?"
		err := db.QueryRow(query, slug).Scan(&artist.ID, &artist.Name, &artist.Slug)

		if err == sql.ErrNoRows {
			http.Error(w, "Artist not found", http.StatusNotFound)
			return
		} else if err != nil {
			log.Printf("ERROR: Failed to fetch artist by slug: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(artist)
		return
	}

	// Otherwise, fetch all artists
	rows, err := db.Query("SELECT id, name, slug FROM artists")
	if err != nil {
		log.Printf("ERROR: Failed to query artists: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var artists []Genre // or []Artist depending on your slice type
	// Wait, let's make sure it's the right slice type:
	var artistList []Artist
	for rows.Next() {
		var a Artist
		if err := rows.Scan(&a.ID, &a.Name, &a.Slug); err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		artistList = append(artistList, a)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(artistList)
}

// 2. Handle Genres (Includes slugs)
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
		w.Header().Set("Content-Type", "application/json")
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
		var genre Genre
		err := rows.Scan(&genre.ID, &genre.Name, &genre.Slug)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		genres = append(genres, genre)
	}

	if err = rows.Err(); err != nil {
		log.Printf("ERROR: Failed to iterate genres: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(genres)
}

// 3. Handle Venues (Includes location and slugs)
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
		w.Header().Set("Content-Type", "application/json")
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
		var venue Venue
		err := rows.Scan(&venue.ID, &venue.Name, &venue.Location, &venue.Slug)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		venues = append(venues, venue)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(venues)
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
	http.HandleFunc("/genres", handleGenres)
	http.HandleFunc("/venues", handleVenues)

	http.ListenAndServe(":8080", nil)

}
