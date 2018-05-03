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

// Parse the map file given
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
		cityName := fields[0]
		// parse direction of a city
		var cityToAdd city
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
		// add city to list of cities
		cities[cityName] = cityToAdd
	}

	// check for any errors
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Couldn't parse file")
		os.Exit(1)
	}
}

func generateAliens(numAliens uint, rng *rand.Rand) {
	for ii := uint(0); ii < numAliens; ii++ {
		randCity := rng.Intn(len(listCities))
		aliens = append(aliens, alien{atCity: listCities[randCity]})
	}
}

func main() {

	// parsing arguments
	worldMap := flag.String("map", "", "the map file containing the cities")
	numAliens := flag.Uint("aliens", 1, "number of aliens to create")
	if len(os.Args) == 1 {
		fmt.Println("you need to fill in arguments")
		flag.Usage()
		return
	}
	flag.Parse()
	if *numAliens == 0 {
		fmt.Println("you need more aliens")
		flag.Usage()
		return
	}

	// parsing map file into a map[cityName]info
	cities = make(map[string]city)
	parseMap(*worldMap)

	// create list of cities (as a helper)
	listCities = make([]string, 0, len(cities))
	for cityName, _ := range cities {
		listCities = append(listCities, cityName)
	}

	// get random source (no need for cryptographic randomness)
	randSource := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(randSource)

	// create aliens
	aliens = make([]alien, 0, *numAliens)
	generateAliens(*numAliens, rng)

}
