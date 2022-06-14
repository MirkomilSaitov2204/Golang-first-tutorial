package main

import "fmt"

func main() {
	p_ex1()
}

func p_ex1() {
	x := 1
	p := &x
	fmt.Printf("x=%T &x=%T p=%T *p=%T &p=%T \n\n", x, &x, p, *p, &p)

	fmt.Printf("x=%d &x=%d p=%d *p=%d &p=%d \n\n", x, &x, p, *p, &p)
	//*p = 2 // dereferencing
}
