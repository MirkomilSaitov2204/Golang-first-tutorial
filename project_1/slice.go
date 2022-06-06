package main

import "fmt"

func main() {
	ex_c1()
}

func ex_c1() {
	days := [...]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}

	fmt.Printf("days=%v\n", days)
	fmt.Printf("len(days)=%d\n", len(days))
	fmt.Printf("cap(days)=%d\n\n\n\n", cap(days))

	weekends := days[5:] //or 5:7

	fmt.Printf("weekends=%v\n", weekends)
	fmt.Printf("len(weekends)=%d\n", len(weekends))
	fmt.Printf("cap(weekends)=%d\n\n\n\n\n", cap(weekends))

	weekdays := days[:5] //or 0:5
	fmt.Printf("weekdays=%v\n", weekdays)
	fmt.Printf("len(weekdays)=%d\n", len(weekdays))
	fmt.Printf("cap(weekdays)=%d\n\n\n\n\n", cap(weekdays))

}
