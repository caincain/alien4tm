package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"

	randomdata "github.com/Pallinder/go-randomdata"
)

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

// coordinates is used to generate a map
type coordinates struct {
	x int
	y int
}

// parseMap parses the map file given and fills a `cities` map
func parseMap(mapFile io.Reader) map[string]*city {

	scanner := bufio.NewScanner(mapFile)

	cities := make(map[string]*city)

	// read line by line
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line) // [name north=city1 south=cit2]
		cityName := fields[0]          // get name
		numLinks := len(fields[1:])    // get links
		if numLinks == 0 {
			fmt.Fprintln(os.Stderr, "One city has no link to other cities")
			os.Exit(1)
		}
		// parse the city's links
		cityToAdd := city{name: cityName, numLinks: numLinks}
		for _, direction := range fields[1:] {
			directionData := strings.Split(direction, "=")
			if len(directionData) != 2 {
				fmt.Fprintln(os.Stderr, "Couldn't parse file")
				os.Exit(1)
			}
			switch directionData[0] {
			case "north":
				cityToAdd.north = directionData[1]
			case "west":
				cityToAdd.west = directionData[1]
			case "east":
				cityToAdd.east = directionData[1]
			case "south":
				cityToAdd.south = directionData[1]
			}

		}
		// add empty alien list
		cityToAdd.aliens = make([]*alien, 0, 1)
		// add city to list of cities
		cities[cityName] = &cityToAdd
	}

	// check for any errors
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Couldn't parse file")
		os.Exit(1)
	}

	return cities
}

// printWorldForFile print the current state of the city
// following the same formating as the initial map input file
func printWorldForFile(listCities []*city, out io.Writer) {
	for _, c := range listCities {
		// ignore destroyed cities
		if c.destroyed {
			continue
		}
		fmt.Fprintf(out, "%s ", c.name)

		if c.north != "" {
			fmt.Fprintf(out, "north=%s ", c.north)
		}
		if c.south != "" {
			fmt.Fprintf(out, "south=%s ", c.south)
		}
		if c.west != "" {
			fmt.Fprintf(out, "west=%s ", c.west)
		}
		if c.east != "" {
			fmt.Fprintf(out, "east=%s ", c.east)
		}
		fmt.Fprintln(out)
	}
}

// generateMap generates a map of `numCities` cities
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
