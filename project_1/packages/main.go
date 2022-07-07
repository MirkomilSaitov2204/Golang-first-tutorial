package main

import (
	"fmt"
	"mirkomil/packages/shapes1"
	"mirkomil/packages/shapes2"
)

func main() {

}

func testCase1() {
	fmt.Println(shapes1.Area(shapes1.Length, shapes1.Width))
	fmt.Println(shapes1.SquareArea(2))
	fmt.Println(shapes2.Area(2, 3))
	fmt.Println(shapes2.SquareArea(2))

}
