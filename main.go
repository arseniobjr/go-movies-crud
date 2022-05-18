package main

import (
	"encoding/json"
	"fmt"
	"log"

	//"math/rand"
	"net/http"
	//"strconv"
	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/jason")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
}

func main() {
	fmt.Println("Hello CRUD API with Golang")

	r := mux.NewRouter()

	movies = append(movies, Movie{
		ID:    "1",
		isbn:  "438227",
		Title: "Movie One",
		Director: &Director{
			Firstname: "John",
			Lastname:  "Doe"},
	},
	)

	movies = append(movies, Movie{
		ID:    "2",
		isbn:  "45455",
		Title: "Movie Two",
		Director: &Director{
			Firstname: "Steve",
			Lastname:  "Smith"},
	},
	)

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server port at 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
