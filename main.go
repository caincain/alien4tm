package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// parseMap parses the map file given and fills a `cities` map
func parseMap(filePath string) {

	// attempt to open map file
	mapFile, err := os.Open(filePath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Couldn't open file")
		os.Exit(1)
	}
	defer mapFile.Close()

	scanner := bufio.NewScanner(mapFile)

	// read line by line
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text()) // [name north=city1 south=cit2]
		cityName := fields[0]                    // get name
		numLinks := len(fields[1:])              // get links
		if numLinks == 0 {
			fmt.Fprintln(os.Stderr, "One city has no link to other cities")
			os.Exit(1)
		}
		// parse the city's links
		cityToAdd := city{numLinks: numLinks}
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
}

// generateAliens creates aliens and place them in random cities
func generateAliens(numAliens int, rng *rand.Rand) {
	for ii := 0; ii < numAliens; ii++ {
		for {
			// find random city
			randCity := rng.Intn(len(listCities))
			atCity := listCities[randCity]
			// check if city has no alien
			if len(atCity.aliens) != 0 {
				continue
			}
			// add to list of aliens
			alienToAdd := alien{atCity: atCity}
			aliens = append(aliens, &alienToAdd)
			// add to the city as well
			atCity.aliens = append(atCity.aliens, &alienToAdd)
			//
			break
		}
	}
}

// iteration goes through one iteration of the game
func iteration(rng *rand.Rand) {
	for _, currentAlien := range aliens {
		// ignore dead aliens
		if currentAlien.dead {
			continue
		}
		atCity := currentAlien.atCity
		// make the alien move
		whichWayIdx := rng.Intn(atCity.numLinks)
	Loop:
		for {
			switch whichWayIdx {
			case 0:
				// if empty direction, go to the next
				if goTo := atCity.north; goTo == "" {
					whichWayIdx = (whichWayIdx + 1) % 4
				} else {
					atCity.aliens = atCity.aliens[:0]                               // remove alien from city
					currentAlien.atCity = cities[goTo]                              // set city of alien
					cities[goTo].aliens = append(cities[goTo].aliens, currentAlien) // add alien to city
					break Loop
				}

			case 1:
				if goTo := atCity.west; goTo == "" {
					whichWayIdx = (whichWayIdx + 1) % 4
				} else {
					atCity.aliens = atCity.aliens[:0]
					currentAlien.atCity = cities[goTo]
					cities[goTo].aliens = append(cities[goTo].aliens, currentAlien)
					break Loop
				}

			case 2:
				if goTo := atCity.east; goTo == "" {
					whichWayIdx = (whichWayIdx + 1) % 4
				} else {
					atCity.aliens = atCity.aliens[:0]
					currentAlien.atCity = cities[goTo]
					cities[goTo].aliens = append(cities[goTo].aliens, currentAlien)
					break Loop
				}

			case 3:
				if goTo := atCity.south; goTo == "" {
					whichWayIdx = (whichWayIdx + 1) % 4
				} else {
					atCity.aliens = atCity.aliens[:0]
					currentAlien.atCity = cities[goTo]
					cities[goTo].aliens = append(cities[goTo].aliens, currentAlien)
					break Loop
				}
			}
		} // endfor

		// is there more than one alien on the city? -> fight
		atCity = currentAlien.atCity
		if len(atCity.aliens) == 2 {
			// aliens kill each other
			atCity.aliens[0].dead = true
			atCity.aliens[1].dead = true
			// destroy city
			atCity.destroyed = true
			// destroy roads
			if atCity.north != "" {
				cities[atCity.north].south = ""
			}
			if atCity.west != "" {
				cities[atCity.west].east = ""
			}
			if atCity.east != "" {
				cities[atCity.east].west = ""
			}
			if atCity.south != "" {
				cities[atCity.south].north = ""
			}
		}

	} // end for aliens
}

//
// MAIN
//
func main() {

	// parsing arguments
	worldMap := flag.String("map", "", "the map file containing the cities")
	numAliens := flag.Int("aliens", 1, "number of aliens to create")
	if len(os.Args) == 1 {
		fmt.Println("you need to fill in arguments")
		flag.Usage()
		return
	}
	flag.Parse()
	if *numAliens <= 0 {
		fmt.Println("you need more aliens")
		flag.Usage()
		return
	}

	// parsing map file into a map[cityName]info
	cities = make(map[string]*city)
	parseMap(*worldMap)

	// create list of cities (as a helper)
	listCities = make([]*city, 0, len(cities))
	for _, c := range cities {
		listCities = append(listCities, c)
	}

	// get random source (no need for cryptographic randomness)
	randSource := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(randSource)

	// create aliens
	if *numAliens > len(listCities) {
		fmt.Println("you can't have more than one aliens per city")
		return
	}
	aliens = make([]*alien, 0, *numAliens)
	generateAliens(*numAliens, rng)

	//
}
