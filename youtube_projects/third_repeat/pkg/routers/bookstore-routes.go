package routers

import (
	"github.com/MirkomilSaitov2204/go-course/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegistratedRoutes = func(r *mux.Router) {

	r.HandleFunc("/books/", controllers.GetAllBooks).Methods("GET")
	r.HandleFunc("/books/{id}", controllers.GetBookById).Methods("GET")
	r.HandleFunc("/books/{id}", controllers.CreateBook).Methods("POST")
	r.HandleFunc("/books/{id}", controllers.UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", controllers.DeleteBook).Methods("DELETE")

}
