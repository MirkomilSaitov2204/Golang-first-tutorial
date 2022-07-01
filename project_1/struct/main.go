package main

import (
	"fmt"
	athletes "mirkomil/struct/types"
	"strings"
)

func changeAthleteName1(p athletes.Player) {
	p.Name = "Anderson"
}
func changeAthleteName2(p *athletes.Player) {
	p.Name = "Anderson"
	p.Country = strings.ToUpper(p.Country)
}

func main() {
	player1 := athletes.Player{"ASSSS", "MMA", 34, athletes.Info{"Brasiz", "Black"}}

	fmt.Println(player1)

	changeAthleteName1(player1)
	fmt.Println(player1)

	changeAthleteName2(&player1)
	fmt.Println(player1)

	fmt.Println(*player1.ToLowerCase())

}
