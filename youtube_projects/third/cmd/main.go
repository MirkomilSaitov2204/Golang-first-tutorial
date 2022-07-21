package main

import (
	"github.com/MirkomilSaitov2204/go-course/pkg/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8887", r))
}
