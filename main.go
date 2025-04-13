package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"github.com/joho/godotenv"
)

// Create Movie struct to hold data from the OMDB API
type Movie struct {
	Title  string `json:"Title"`
	Year   string `json:"Year"`
	Poster string `json:"Poster"`
}

// Create Search struct which holds an array of Movie structs
type SearchResponse struct {
	Search []Movie `json:"Search"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// call the handleMovie function when /movie is visited
	http.HandleFunc("/movie", handleMovie)
  http.HandleFunc("/search", handleSearch)
	fmt.Println("Server running at http://localhost.8080")

	// start a web server at port 8080 and use the default handler
	http.ListenAndServe(":8080", nil)
}

func handleMovie(w http.ResponseWriter, r *http.Request){
	// Get "title" from the user's query
	movieTitle := r.URL.Query().Get("title")
	movieYear := r.URL.Query().Get("year")

	if movieTitle == "" {
		http.Error(w, "Missing title query parameters", http.StatusBadRequest)
		return
	}
	// Build OMDB API URL
	apiKey := os.Getenv("OMDB_API_KEY")
	endpoint := "https://www.omdbapi.com/"
	movieTitleQuery := url.QueryEscape(movieTitle)

	var fullURL string
	if movieYear != "" {
		fullURL = fmt.Sprintf("%s?apikey=%s&t=%s&y=%s", endpoint, apiKey, movieTitleQuery, movieYear)
	} else {
		fullURL = fmt.Sprintf("%s?apikey=%s&t=%s", endpoint, apiKey, movieTitleQuery)
	}

	// Make HTTP request to OMDB
	res, err := http.Get(fullURL)
	if err != nil {
		http.Error(w, "Failed to reach OMDB", http.StatusInternalServerError)
	}
	// For raw HTTP call using http.Get, use defer res.Body.Close() to avoid leaking resources.
	// In production-grade code, can a library that abstracts this away
	// Close the response body after handleMovie() is done running
	defer res.Body.Close()

	// Read the response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		http.Error(w, "Failed to read response from OMDB", http.StatusInternalServerError)
		return
	}

	// Deserialize into movie struct
	var movie Movie
	if err := json.Unmarshal(body, &movie); err != nil {
		http.Error(w, "Failed to parse movie data", http.StatusInternalServerError)
		return
	}

	// Return JSON to user
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movie)
}

// Uses s=keyword to find multiple possible matches (e.g., all movies with “batman”)
func handleSearch(w http.ResponseWriter, r *http.Request){
	searchKeyword := r.URL.Query().Get(("title"))
	if searchKeyword == "" {
		http.Error(w, "Missing search keyword!", http.StatusInternalServerError)
		return
	}

	apiKey := os.Getenv(("OMDB_API_KEY"))
	endpoint := "https://www.omdbapi.com/"
	movieTitleQuery := url.QueryEscape(searchKeyword)

	fullURL := fmt.Sprintf("%s?apikey=%s&s=%s", endpoint, apiKey, movieTitleQuery)

	res, err := http.Get(fullURL)
  if err != nil {
		http.Error(w, "Error fetching results", http.StatusInternalServerError)
		return
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		http.Error(w, "Failed to read response from OMDB", http.StatusInternalServerError)
		return
	}

	var result SearchResponse
	if err := json.Unmarshal(body, &result); err != nil {
		http.Error(w, "Error parsing", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
