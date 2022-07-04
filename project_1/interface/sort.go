package main

import (
	"fmt"
	"sort"
)

func main() {

	n := []int{4, 5, 2, 35, 73, 8, 9, 2}
	//fmt.Println(n)
	//sort.Ints(n)
	//fmt.Println(n)

	sort.Sort(sort.IntSlice(n))
	fmt.Println(n)
	sort.Sort(sort.Reverse(sort.IntSlice(n)))
	fmt.Println(n)

}
