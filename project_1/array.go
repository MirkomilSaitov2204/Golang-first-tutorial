package main

import (
	"fmt"
	"os"
)

func main() {
	//array()
	//ex1()
	ex2()
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
