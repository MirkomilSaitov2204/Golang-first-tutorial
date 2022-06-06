package main

import (
	"fmt"
	"os"
)

func main() {
	//array()
	//ex1()
	//ex2()
	//ex3()
	//ex4()
	ex5()
}

func array() {
	days := [...]string{"mon", "tue"}

	fmt.Println(days)

	nums := [...]int{10, 20, 30}
	nums2 := [3]int{10, 20, 30}
	var nums3 [3]int = [3]int{10, 20, 30} // redundant

	fmt.Println(nums)
	fmt.Println(nums2)
	fmt.Println(nums3)

}

func ex1() {
	var nums [3]int

	var sum1 int
	var sum2 int

	fmt.Printf("nums=%v type=%T length=%d\n", nums, nums, len(nums))

	for i := range nums {
		sum1 += i
		sum2 += nums[i]
	}
	fmt.Println(sum1, sum2)
}

func ex2() {
	nums := [...]int{10, 15, 30}
	var sum int
	nums[1] = 20

	for i := range nums {
		sum += nums[i]
	}

	fmt.Println(sum)
	fmt.Println(nums)
	fmt.Println(nums[0])
	fmt.Println(nums[2])
	fmt.Println(len(nums))

	for j := 0; j < len(nums); j++ {
		fmt.Print(nums[j], " ")
	}
	fmt.Println()

	var j int

	for j = 0; j < len(nums); j++ {
		fmt.Print(nums[j], " ")
	}
	fmt.Fprintf(os.Stdout, "\n%d", j)
}

func ex3() {
	x := [...]float32{2.1, 3.2, 4.7}
	fmt.Println(x)

	var total float32
	for _, val := range x {
		total += val
	}
	fmt.Println(total)
}

func ex4() {
	type Size int

	const (
		EX Size = iota
		LG
		MD
		SM
	)
	sz := [...]string{EX: "Extra Large", LG: "Large", MD: "Medium", SM: "Small"}
	fmt.Println(MD, sz[MD])
	fmt.Println(sz)

	x := [...]int{3: 10, 20}
	fmt.Println(x)
}

func ex5() {
	a := [2][3]int{{2, 4, 6}, {4, 6, 8}}
	fmt.Printf("a=%v type=%T len=%d\n", a, a, len(a))

	var b [2][3]int

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			b[i][j] = (j + j + 1) * 2
		}
	}

	fmt.Printf("a=%v type=%T len=%d\n", b, b, len(b))

	a[0][1] = 5
	a[1][1] = 10

	fmt.Printf("a=%v type=%T len=%d\n", a, a, len(a))

}
