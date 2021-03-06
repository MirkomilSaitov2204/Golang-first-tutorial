package main

import "fmt"

type artist struct {
	name string
	age  int
}

func (a *artist) showName() {
	fmt.Println(a.name)
}
func (a *artist) showAge() {
	fmt.Println(a.age)
}

type singer struct {
	field string
	artist
}

func (a *singer) showField() {
	fmt.Println(a.field)
}
func (a *singer) showAge() {
	fmt.Println(a.age)
}

func main() {
	r1 := artist{
		name: "Michael Jackson",
		age:  99,
	}

	fmt.Println(r1)
	r1.showName()
	r1.showAge()

	s1 := singer{
		field:  "singer",
		artist: r1,
	}

	fmt.Printf("\n%v\n", s1)
	s1.showName()
	s1.showAge()
	s1.showField()

	// ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~
	s2 := singer{
		artist: artist{
			name: "Frank Sinatra",
			age:  109,
		},
		field: "singer",
	}

	fmt.Printf("\n%v\n", s2)
	s2.showAge()
	s2.showName()
	s2.showField()
}
