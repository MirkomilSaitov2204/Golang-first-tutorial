package main

import (
	"github.com/MirkomilSaitov2204/go-course/pkg/config"
	"github.com/MirkomilSaitov2204/go-course/pkg/handlers"
	"github.com/MirkomilSaitov2204/go-course/pkg/render"
	"log"
	"net/http"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	app.TemplateCache = tc

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	//http.HandleFunc("/divide", Divide)

	_ = http.ListenAndServe(portNumber, nil)
}
