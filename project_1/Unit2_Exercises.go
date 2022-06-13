package main

import "fmt"

func main() {
	//fmt.Printf("%.2f", celsiusToFahrenheit(31))
	fmt.Println(sumDigits(889))
}

func celsiusToFahrenheit(degree float64) float64 {
	return 1.8*degree + 32
}

func sumDigits(digit int) int {
	last := digit % 10
	devide := digit / 10
	right := devide % 10
	left := devide / 10

	return last + right + left
}
