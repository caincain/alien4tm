package main

import (
	"math/rand"
	"strconv"
	"strings"

	randomdata "github.com/Pallinder/go-randomdata"
)

func generateMap(numCities int, rng *rand.Rand) (resMap map[string]*city, resList []*city) {

	// list by coordinates
	worldMap := make(map[coordinates]*city, numCities)

	// create map and list
	resMap = make(map[string]*city, numCities)
	resList = make([]*city, 0, numCities)

	// create cities
	for i := 0; i < numCities; i++ {

		var cityToAdd city

		// generate a random name for the city
		name := strings.Replace(randomdata.City(), " ", "", -1) + strconv.Itoa(rng.Intn(50))
		for _, ok := resMap[name]; ok; { // make sure it doesn't already exist
			name = strings.Replace(randomdata.City(), " ", "", -1) + strconv.Itoa(rng.Intn(50))
		}
		cityToAdd.name = name

		// if it's the first city, just place it
		if i == 0 {
			worldMap[coordinates{0, 0}] = &cityToAdd
			resMap[cityToAdd.name] = &cityToAdd
			resList = append(resList, &cityToAdd)
			continue
		}

		// add next to a random city
		var newLocation coordinates
		for true {
			// find random city
			randCity := rng.Intn(len(resList))
			neighbour := resList[randCity]

			// find space next to random city
			newLocation = neighbour.coordinates
			randDirection := rng.Intn(4)
			switch randDirection {
			case 0:
				newLocation.x++
			case 1:
				newLocation.y++
			case 2:
				newLocation.x--
			case 3:
				newLocation.y--
			}

			// anyone there already?
			if worldMap[newLocation] == nil {
				worldMap[newLocation] = &cityToAdd
				cityToAdd.coordinates = newLocation
				break
			}
		}

		// update cities' links
		if n := worldMap[coordinates{x: newLocation.x, y: newLocation.y + 1}]; n != nil { // north
			cityToAdd.north = n.name // cityToAdd -> neighbour
			cityToAdd.numLinks++     // increment our link
			n.south = cityToAdd.name // neighbour -> cityToAdd
			n.numLinks++             // increment the neighbour number of links
		}
		if n := worldMap[coordinates{x: newLocation.x, y: newLocation.y - 1}]; n != nil { // south
			cityToAdd.south = n.name
			cityToAdd.numLinks++
			n.north = cityToAdd.name
			n.numLinks++
		}
		if n := worldMap[coordinates{x: newLocation.x + 1, y: newLocation.y}]; n != nil { // east
			cityToAdd.east = n.name
			cityToAdd.numLinks++
			n.west = cityToAdd.name
			n.numLinks++
		}
		if n := worldMap[coordinates{x: newLocation.x - 1, y: newLocation.y}]; n != nil { // west
			cityToAdd.west = n.name
			cityToAdd.numLinks++
			n.east = cityToAdd.name
			n.numLinks++
		}

		// add the new city
		resList = append(resList, &cityToAdd)
		resMap[cityToAdd.name] = &cityToAdd
	}

	return resMap, resList
}
