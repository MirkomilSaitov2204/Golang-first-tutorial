package main

import (
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.tmpl")

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
