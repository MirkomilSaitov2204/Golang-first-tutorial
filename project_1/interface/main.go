package main

import "fmt"

type rectangle struct {
	w, l int
}

func (r *rectangle) area() int {
	return r.w * r.l
}

type square struct {
	s int
}

func (r *square) area() int {
	return r.s * r.s
}

type shape interface {
	area() int
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	r1 := rectangle{2, 3}
	fmt.Println(r1.area())
	info(&r1)

	s1 := square{3}
	fmt.Println(s1.area())
	info(&s1)

	var shapes [2]shape
	shapes[0] = &r1
	shapes[1] = &s1
	info(shapes[0])
	info(shapes[1])

}
