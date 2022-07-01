package main

import "fmt"

type movie struct {
	name  string
	actor string
}

type imdb struct {
	movie
	name  string
	actor string
}

func (m movie) fullInfo() string {
	return m.name + " " + m.actor
}

func (m imdb) fullInfo() string {
	return m.name + " " + m.actor + " " + m.movie.name + " " + m.movie.actor
}

func main() {
	m := movie{"Mike", "AAAAAAAAAAAAAAA"}
	i1 := imdb{m, "444444", "gggggggggg"}

	i2 := imdb{
		movie: movie{
			name:  "The goodfather",
			actor: "Hancock",
		},
		name:  "Marshal",
		actor: "Hello world",
	}

	fmt.Println(m.fullInfo())
	fmt.Println(i1.fullInfo())
	fmt.Println(i2.fullInfo())
}
