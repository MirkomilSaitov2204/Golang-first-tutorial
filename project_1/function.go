package main

import "fmt"

func main() {
	//f_ex1()
	//f_ex2()
	//f_ex3()
	//f_ex4()
	f_ex5()
}

func f_ex1() {
	fmt.Println(add(2, 21))

	x1, x2 := swap1("one", "two")
	fmt.Println(x1, x2)

	y1, y2 := swap2("one", "two")
	fmt.Println(y1, y2)

}

func add(a, b int) int {
	return a + b
}

func swap1(a, b string) (string, string) {
	return b, a
}

func swap2(a, b string) (r1, r2 string) {
	r1 = b
	r2 = a
	return
}

func f_ex2() {
	stock1 := []float32{98, 20, 15}
	stock2 := []float32{92, 40, 16, 60, 79, 45}

	fmt.Println(avg(stock1))
	fmt.Println(avg(stock2))
}

func avg(stock []float32) float32 {
	var total float32
	for _, item := range stock {
		total += item
	}
	return total / float32(len(stock))
}

func f_ex3() {
	sum(1, 2)
	sum(3, 4, 5)

	nums := []int{1, 2, 3, 4, 5}

	sum(nums...)
	sum2(nums)
}

func sum(nums ...int) {
	var total int

	for _, i := range nums {
		total += i
	}
	println(total)
}

func sum2(nums []int) {
	var total int

	for _, i := range nums {
		total += i
	}
	println(total)
}

func f_ex4() {
	names := []string{"Amy", "SMT", "Helen"}
	echo(names...)
}

func echo(names ...string) {
	for _, s := range names {
		fmt.Println(s)
	}
}

func f_ex5() {
	data := []string{"Mike", "Hello", "", "Timmy"}

	fmt.Println(data)

	fmt.Println(trimSlice(data))

}

func trimSlice(names []string) []string {

	var newData []string
	for _, d := range names {
		if d != "" {
			newData = append(newData, d)
		}
	}
	return newData
}
