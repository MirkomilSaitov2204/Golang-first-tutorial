package main

import (
	"fmt"
	"time"
)

//type player struct {
//	name, sport string
//	age         int
//}

type employee struct {
	name, job    string
	lastLoggedIn string
	DOB          date.Date
}

func main() {
	var emp employee

	emp.name = "Mike"
	emp.job = "GO prog"
	emp.lastLoggedIn = time.Now().Format(time.RFC850)
	emp.DOB = date.Today()
	fmt.Println(emp)

	//type myType float64
	//var total myType
	//
	//total = 44
	//fmt.Printf("%.2f %T \n", total, total)
	//
	//var total2 float64
	//
	//total2 = float64(total)
	//fmt.Printf("%.2f %T \n", total2, total2)

	//player1 := player{"Messi", "Scoorer", 24}
	//fmt.Println(player1)
	//player2 := player{
	//	name:  "2",
	//	sport: "2",
	//	age:   2,
	//}
	//fmt.Println(player2)
	//
	//var player3 player
	//fmt.Println(player3)
	//player3.name = "Mike"
	//player3.sport = "boxing"
	//player3.age = 2
	//fmt.Println(player3)
	//
	//player4 := player{
	//	name:  "SMTH",
	//	sport: "smht",
	//}
	//fmt.Println(player4)
	//
	//player5 := new(player)
	//player5.name = "UUU"
	//fmt.Println(*player5)

	//player1 := struct {
	//	name, sport string
	//	age         int
	//}{"Mike", "1212312", 3}
	//
	//fmt.Println(player1)

}
