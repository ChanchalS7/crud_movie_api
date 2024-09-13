package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Movie represents a movie with an ID, ISBN, title, and director
type Movie struct {
	ID       string    `json:"id"`       // Unique identifier for the movie
	Isbn     string    `json:"isbn"`     // ISBN of the movie
	Title    string    `json:"title"`    // Title of the movie
	Director *Director `json:"director"` // Director information
}

// Director represents the director of a movie
type Director struct {
	Firstname string `json:"firstname"` // Director's first name
	Lastname  string `json:"lastname"`  // Director's last name
}

// In-memory movie store
var movies []Movie

// getMovies handles GET requests and returns a list of movies
func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // Set response header to JSON
	json.NewEncoder(w).Encode(movies)                  // Encode movies slice as JSON and write to response
}

// deleteMovie handles DELETE requests and removes a movie by ID
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // Set response header to JSON
	params := mux.Vars(r)                              // Extract variables from the request URL

	// Loop through movies slice to find the movie with the given ID
	for index, item := range movies {
		if item.ID == params["id"] {
			// Remove the movie from the slice
			movies = append(movies[:index], movies[index+1:]...)
			json.NewEncoder(w).Encode(movies) // Return updated list of movies as JSON
			return
		}
	}
}

// getMovie handles GET requests and returns a single movie by ID
func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // Set response header to JSON
	params := mux.Vars(r)                              // Extract variables from the request URL

	// Loop through movies slice to find the movie with the given ID
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item) // Return the found movie as JSON
			return
		}
	}
}

// createMovie handles POST requests and creates a new movie
func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // Set response header to JSON
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)     // Decode the request body into a Movie object
	movie.ID = strconv.Itoa(rand.Intn(1000000000)) // Generate a random ID for the new movie
	movies = append(movies, movie)                 // Add the new movie to the slice
	json.NewEncoder(w).Encode(movie)               // Return the newly created movie as JSON
}

// updateMovie handles PUT requests and updates an existing movie by ID
func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // Set response header to JSON
	params := mux.Vars(r)                              // Extract variables from the request URL

	// Loop through movies slice to find the movie with the given ID
	for index, item := range movies {
		if item.ID == params["id"] {
			// Remove the existing movie from the slice
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie) // Decode the request body into a Movie object
			movie.ID = params["id"]                    // Set the ID of the movie to the ID from the URL
			movies = append(movies, movie)             // Add the updated movie to the slice
			json.NewEncoder(w).Encode(movie)           // Return the updated movie as JSON
			return
		}
	}
}

// main sets up the router and starts the server
func main() {
	r := mux.NewRouter() // Create a new router

	// Initialize some sample movies
	movies = append(movies, Movie{ID: "1", Isbn: "43725", Title: "Movie one", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "43729", Title: "Movie two", Director: &Director{Firstname: "John", Lastname: "Smith"}})

	// Define routes and associate them with handler functions
	r.HandleFunc("/api/movies", getMovies).Methods("GET")
	r.HandleFunc("/api/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/api/movies", createMovie).Methods("POST")
	r.HandleFunc("/api/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/api/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8080\n") // Print a message indicating server start
	log.Fatal(http.ListenAndServe(":8080", r))   // Start the server and log any errors
}
