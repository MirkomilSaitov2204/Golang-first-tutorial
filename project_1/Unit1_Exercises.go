package main

import "fmt"

func main() {
	//first()
	//second()
	//third()
	//fourth()
	//fifth()
	findArea()
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

func second() {
	fmt.Println("Welcome to Go Programming")
	fmt.Println("Programming is fun ")
}

func third() {
	times := 5

	for i := 1; i <= times; i++ {
		fmt.Printf("Welcome to Java %v\n\n", i)
	}
}

func fourth() {
	//1.4
	arr := [4][3]int{{1, 1, 1}, {2, 4, 8}, {3, 9, 27}, {4, 16, 64}}

	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("    %v", arr[i][j])
		}
		fmt.Println()
	}
}

func fifth() {
	var val float32

	val = (9.5*4.5 - 2.5*3) / (45.5 - 3.5)
	fmt.Println(val)
}

func findArea() {
	var radius float32
	var pi float32
	pi = 3.14

	fmt.Println("Please enter radius")

	fmt.Scan(&radius)

	perimetr := 2 * radius * pi
	area := radius * radius * pi

	fmt.Println(perimetr)
	fmt.Println(area)

}
