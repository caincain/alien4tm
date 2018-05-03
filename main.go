package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	randomdata "github.com/Pallinder/go-randomdata"
)

// generateAliens creates aliens and place them in random cities
func generateAliens(numAliens int, rng *rand.Rand) {
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
}

// iteration goes through one iteration of the game
func iteration(rng *rand.Rand) {
	for _, currentAlien := range aliens {
		// ignore dead aliens
		if currentAlien.dead {
			continue
		}
		atCity := currentAlien.atCity
		// ignore trapped aliens
		if atCity.numLinks == 0 {
			continue
		}
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
			fmt.Printf("alien: %s and %s died in a fight.\n", atCity.aliens[0].name, atCity.aliens[1].name)
			// aliens kill each other
			atCity.aliens[0].dead = true
			atCity.aliens[1].dead = true
			deadAliens += 2
			// destroy city
			atCity.destroyed = true
			fmt.Printf("city: %s has been destroyed.\n", atCity.name)
			// destroy roads
			if atCity.north != "" {
				cities[atCity.north].south = ""
				cities[atCity.north].numLinks--
			}
			if atCity.west != "" {
				cities[atCity.west].east = ""
				cities[atCity.west].numLinks--
			}
			if atCity.east != "" {
				cities[atCity.east].west = ""
				cities[atCity.east].numLinks--
			}
			if atCity.south != "" {
				cities[atCity.south].north = ""
				cities[atCity.south].numLinks--
			}
		}

	} // end for aliens
}

//
// MAIN
//
func main() {

	// get random source (no need for cryptographic randomness)
	randSource := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(randSource)

	// parsing arguments
	// Note that we assume that the map file is correct!
	worldMap := flag.String("map", "", "the map file containing the cities")
	numAliens := flag.Int("aliens", 1, "number of aliens to create")
	genMap := flag.Int("genMap", 0, "optional argument to generate a world map")

	if len(os.Args) == 1 {
		fmt.Println("you need to fill in arguments")
		flag.Usage()
		return
	}

	flag.Parse()

	// generating a map ?
	if numCities := *genMap; numCities > 0 {
		_, resList := generateMap(numCities, rng)
		printWorldForFile(resList, os.Stdout)
		return
	}

	// need at least 1 alien
	if *numAliens <= 0 {
		fmt.Println("you need more aliens")
		flag.Usage()
		return
	}

	// parsing map file into a map[cityName]info
	mapFile, err := os.Open(*worldMap)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Couldn't open file")
		os.Exit(1)
	}
	cities = parseMap(mapFile)
	mapFile.Close()

	// create list of cities (as a helper)
	listCities = make([]*city, 0, len(cities))
	for _, c := range cities {
		listCities = append(listCities, c)
	}

	// create aliens
	if *numAliens > len(listCities) {
		fmt.Println("you can't have more than one aliens per city")
		return
	}
	aliens = make([]*alien, 0, *numAliens)
	generateAliens(*numAliens, rng)

	// main game loop
	for i := 0; i < 10000; i++ {
		fmt.Println("step", i)
		// test if game has ended
		if deadAliens == len(aliens) {
			break
		}
		// game iteration
		iteration(rng)
	}

	//
	fmt.Println("Fin. Current state of city:")

	// print out the current world
	printWorldForFile(listCities, os.Stdout)
}
