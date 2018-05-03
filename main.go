package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	termbox "github.com/nsf/termbox-go"
)

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
	slowFlag := flag.Bool("slow", false, "slow the game")
	fullscreen := flag.Bool("fullscreen", false, "display the game on fullscreen")

	if len(os.Args) == 1 {
		fmt.Println("you need to fill in arguments")
		flag.Usage()
		return
	}

	flag.Parse()

	if *fullscreen {
		err := termbox.Init()
		if err != nil {
			panic(err)
		}
		defer termbox.Close()
	}

	slow = *slowFlag

	// just want to generate a map ?
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
	cities := parseMap(mapFile)
	mapFile.Close()

	if *numAliens > len(cities) {
		fmt.Println("you can't have more than one aliens per city")
		return
	}

	// create game + run
	state := newGame(cities, *numAliens, rng)
	state.run(10000)

	//
	fmt.Println()
	fmt.Println("----------- Fin. ------------")
	fmt.Println("----Current state of city----")
	fmt.Println()

	// print out the current world
	printWorldForFile(state.listCities, os.Stdout)

	fmt.Println()
	fmt.Println("press a key to exit _")
	fmt.Scanf("%s")
}
