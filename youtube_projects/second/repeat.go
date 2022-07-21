package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type Movie struct {
	ID           string `json:"id"`
	Title        string `json:"title"`
	Published_at string `json:"published_at"`
	Director     *Director
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

func showMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
		}

	}
	//json.NewEncoder(w).Encode(movies)
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode("Sueccessfully deleted")
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}

	}
}

func main() {

	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Published_at: "22.04.1997", Title: "Spidermen", Director: &Director{Firstname: "Mirkomil", Lastname: "Saitov"}})
	movies = append(movies, Movie{ID: "2", Published_at: "22.04.1998", Title: "Spidermen 2", Director: &Director{Firstname: "Mirkomil 2", Lastname: "Saitov 2"}})

	r.HandleFunc("/", getMovies).Methods("GET")
	r.HandleFunc("/{id}", showMovie).Methods("GET")
	r.HandleFunc("/", createMovie).Methods("POST")
	r.HandleFunc("/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/{id}", deleteMovie).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8099", r))
}
