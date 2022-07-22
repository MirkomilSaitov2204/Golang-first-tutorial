package cmd

import (
	"github.com/MirkomilSaitov2204/go-course/pkg/routers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routers.RegistratedRoutes(r)

	log.Fatal(http.ListenAndServe(":8070", r))

}
