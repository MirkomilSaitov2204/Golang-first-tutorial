package main

import "fmt"

func main() {
	//first()
	//second()
	//third()
	fourth()
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

func ()  {
	
}
