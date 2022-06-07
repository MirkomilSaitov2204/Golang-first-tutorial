package main

import "fmt"

func main() {
	//ex_c1()
	//ex_c2()
	//ex_c3()
	//ex_c4()
	ex_c5()
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

func ex_c2() {
	s := "ABCDEFGHJK"

	slc := s[1:3]

	fmt.Printf("s=%v -------- type(s)=%T ------ slc=%v -------- type(slc)=%T", s, s, slc, slc)
}

func ex_c3() {
	var f []float32

	f = append(f, 1.2)
	f = append(f, 1.5, 2.1, 3.1, 3.5, 5.7)

	fmt.Println(f)
	f = append(f, f...)
	fmt.Println(f)

}

func ex_c4() {
	team1 := []string{"Messi", "Ronaldo", "Kaka", goalkeeper()}

	team := append([]string{"Nani", "Totti", "Evra"}, team1...)

	team = append(team, midfielders()...)

	for i, val := range team {
		fmt.Println(i, val)
	}
}

func goalkeeper() string {
	return "Buffon"
}

func midfielders() []string {
	return []string{"Ramos", "Pique", "Puyol"}
}

func ex_c5() {
	s1 := []int{1, 2, 3}
	s2 := make([]int, 2)

	copy(s2, s1)

	s3 := append(s2, 3, 4, 5)

	fmt.Println(s1, s2, s3)

	fmt.Println()

	for i := range s3 {
		fmt.Print(s3[i], " ")
	}
	fmt.Println()

	for _, val := range s3 {
		fmt.Print(val, " ")
	}
}
