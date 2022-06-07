package main

import "fmt"

func main() {
	first()
}

func first() {
	//  pi = 4 * (1-1/3+1/5-1/7+1/9.....+1/n)

	fmt.Println("Enter the first name")

	var first int

	fmt.Scanln(&first)

	var result = 1
	for i := 1; i <= first; i++ {
		result += (2*i - 1)
	}
	fmt.Println(result)

}
