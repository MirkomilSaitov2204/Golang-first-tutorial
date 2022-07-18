package main

import (
	"github.com/MirkomilSaitov2204/go-course/pkg/handlers"
	"net/http"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	//http.HandleFunc("/divide", Divide)

	_ = http.ListenAndServe(portNumber, nil)
}
