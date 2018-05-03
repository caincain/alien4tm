package main

// city is the structure to hold such parsed cities:
// Foo north=Bar west=Baz south=Qu-ux
type city struct {
	north    string
	west     string
	east     string
	south    string
	numLinks int
	aliens   []*alien
}

// cities contains all of our cities :)
var cities map[string]*city
var listCities []*city
