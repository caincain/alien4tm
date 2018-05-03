package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
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
		//
		fmt.Println(cities)
	}

	// check for any errors
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Couldn't parse file")
		os.Exit(1)
	}
}

func main() {

	// parsing arguments
	worldMap := flag.String("map", "", "the map file containing the cities")
	numAliens := flag.Uint("aliens", 0, "number of aliens to create")
	if len(os.Args) == 1 {
		flag.Usage()
		return
	}
	flag.Parse()

	// parsing map file
	cities = make(map[string]city)
	parseMap(*worldMap)

	//
	fmt.Println(*numAliens)

}
