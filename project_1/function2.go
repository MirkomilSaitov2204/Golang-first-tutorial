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
	//h1 := sayGreetings("ESP")
	//fmt.Println(h1())
	//h1 = sayGreetings("GER")
	//fmt.Println(sayGreetings("ENG")())
	//fmt.Println(h1())

	//num := getPositiveInt()
	//fmt.Println(num()) //1
	//fmt.Println(num()) //2
	//fmt.Println(num()) //3
	//
	//num2 := getPositiveInt()
	//fmt.Println(num2()) //1

	addCounter, multCounter := addBy(), multBy()

	fmt.Println(addCounter(2))
	fmt.Println(addCounter(3))
	fmt.Println(multCounter(3))
	fmt.Println(multCounter(3))

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

//func getPositiveInt() func() int {
//	i := 0
//	return func() int {
//		i++
//		return i
//	}
//}

func addBy() func(int) int {
	total := 0
	return func(i int) int {
		total += i
		return total
	}
}

func multBy() func(int) int {
	total := 1
	return func(i int) int {
		total *= i
		return total
	}
}
