package handlers

import (
	"github.com/MirkomilSaitov2204/go-course/pkg/config"
	"github.com/MirkomilSaitov2204/go-course/pkg/render"
	"net/http"
)

// Repo
var Repo *Repository

// Repository
type Repository struct {
	App *config.AppConfig
}

// NewRepo
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers
func NewHandlers(r *Repository) {
	Repo = r
}

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.gohtml")
}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.gohtml")
}

//func Divide(w http.ResponseWriter, r *http.Request) {
//	f, err := divideValues(100.0, 10.0)
//	if err != nil {
//		fmt.Fprintf(w, "Cannot divide by 0")
//	}
//	fmt.Fprintf(w, fmt.Sprintf("Answer is %f", f))
//}
//
//func divideValues(x, y float32) (float32, error) {
//	if y <= 0 {
//		err := errors.New("Cannot divide by zero")
//		return 0, err
//	}
//	return x / y, nil
//}
//
//func addValue(x, y int) int {
//	return x + y
//}
