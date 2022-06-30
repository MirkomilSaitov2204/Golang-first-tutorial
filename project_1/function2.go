package main

import "fmt"

func main() {
	//printMsg("Calling a function")
	//
	//showMsg := func(msg string) {
	//	fmt.Println(msg)
	//}
	//showMsg("Second function")
	//
	//func(msg string) {
	//	fmt.Println(msg)
	//}("quick message")
	//
	//factor := 2
	//
	//mult := func(i int, j int) int {
	//	return i * j * factor
	//}
	//fmt.Println(mult(3, 4))

	////////////////////////
	h1 := sayGreetings("ESP")
	fmt.Println(h1())
	h1 = sayGreetings("GER")
	fmt.Println(sayGreetings("ENG")())
	fmt.Println(h1())
}

//func printMsg(msg string) {
//	fmt.Println(msg)
//}

//func sayGreetings(lang string) func() string {
//	h1 := func() string {
//		return "Hello"
//	}
//
//	if lang == "ESP" {
//		h1 = func() string {
//			return "Hola"
//		}
//	} else if lang == "GER" {
//		h1 = func() string {
//			return "Hallo"
//		}
//	}
//
//	return h1
//}
