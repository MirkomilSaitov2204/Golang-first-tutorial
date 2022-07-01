package athletes

import "strings"

type Info struct {
	Country   string
	HairColor string
}

type Player struct {
	Name  string
	Sport string
	Age   int
	Info
}

func (p *Player) ToLowerCase() *Player {
	p.Name = strings.ToLower(p.Name)
	p.Sport = strings.ToLower(p.Sport)
	p.Country = strings.ToLower(p.Country)
	p.HairColor = strings.ToLower(p.HairColor)
	return p
}
