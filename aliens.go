package main

import (
	"math/rand"

	randomdata "github.com/Pallinder/go-randomdata"
)

type alien struct {
	name   string
	atCity *city
	dead   bool
}

// generateAliens creates aliens and place them in random cities
func generateAliens(numAliens int, listCities []*city, rng *rand.Rand) []*alien {

	aliens := make([]*alien, 0, numAliens)

	for i := 0; i < numAliens; i++ {
		// give alien a name
		alienToAdd := alien{name: randomdata.SillyName()}

		for {
			// find random city
			randCity := rng.Intn(len(listCities))
			atCity := listCities[randCity]
			// check if city has no alien
			if len(atCity.aliens) != 0 {
				continue
			}
			alienToAdd.atCity = atCity
			// add to list of aliens
			aliens = append(aliens, &alienToAdd)
			// add to the city as well
			atCity.aliens = append(atCity.aliens, &alienToAdd)
			//
			break
		}
	}

	return aliens
}

func (a *alien) moveAlien(from, to *city) {
	from.aliens = from.aliens[:0] // there can only be one alien in a city
	to.aliens = append(to.aliens, a)
	a.atCity = to
}
