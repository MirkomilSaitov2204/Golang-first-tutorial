package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	//m_ex1()
	//m_ex2()
	//m_ex3()
	//m_ex4()
	m_ex5()
}

func m_ex1() {
	empSalary := map[string]float64{
		"blake":  123.00,
		"Parker": 1343444.00,
	}

	fmt.Println(empSalary)
}

func m_ex2() {
	values := []string{"ABC", "ACB", "BAC", "BCA", "CAB", "CBA"}

	factor := []int{100, 10, 1}

	hashKey := 0

	for v := range values {
		bytes := []byte(values[v])
		f := 0
		hashKey = 0

		for i := range bytes {
			fmt.Print(bytes[i], " ")
			hashKey += int(bytes[i]) * factor[f]
			f++
		}

		fmt.Printf("Hashkey: %d \n", hashKey)
	}
}

func m_ex3() {
	values := []string{"tea", "tea", "eat", "eta", "teal", "teatop", "teaparty", "tear"}

	for v := range values {
		hashKey := hashK(values[v])
		fmt.Printf("%-10s (hashKey: %10d) %v\n", values[v], hashKey, []byte(values[v]))
	}
}

func hashK(s string) int {
	result := 0
	idx := 0
	for i := len(s) - 1; i >= 0; i-- {
		result += int(math.Pow10(i)) * int(s[idx])
	}

	return result
}

func m_ex4() {
	//var days = make(map[string]int)
	var days map[string]int

	if days == nil {
		fmt.Println("days  is nill")
		days = make(map[string]int)
	}

	fmt.Println(days)

	days["sun"] = 3
	days["sun"] = -2

	fmt.Println(days)

	days["mon"] = 1
	days["mon"]++
	fmt.Println(days)

}

func m_ex5() {
	sal := map[string]float64{
		"Blake":  60000.00,
		"Parker": 120000.00,
		"Dakota": 93000.00,
	}

	names := make([]string, 0, len(sal))

	for n := range sal {
		names = append(names, n)
	}

	sort.Strings(names)
	fmt.Println(names)

	for _, n := range names {
		fmt.Printf("%s\t%.2f\n", n, sal[n])
	}

}
