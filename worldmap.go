package main

// city is the structure to hold such parsed cities:
// Foo north=Bar west=Baz south=Qu-ux
type city struct {
	name string

	coordinates coordinates

	north string
	west  string
	east  string
	south string

	numLinks int // some directions might link to nothing

	aliens    []*alien
	destroyed bool
}

type coordinates struct {
	x int
	y int
}

// cities contains all of our cities :)
var cities map[string]*city
var listCities []*city
