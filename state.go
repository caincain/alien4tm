package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

var slow bool

type gameState struct {
	step       int              // stepping into the game
	deadAliens int              // number of dead aliens
	cities     map[string]*city // map of cities
	listCities []*city          // list of cities (helper)
	aliens     []*alien         // list of aliens

	rng *rand.Rand
}

func newGame(cities map[string]*city, numAliens int, rng *rand.Rand) *gameState {

	// init
	gs := gameState{rng: rng, cities: cities}

	// create list of cities (as a helper)
	gs.listCities = make([]*city, 0, len(cities))
	for _, c := range cities {
		gs.listCities = append(gs.listCities, c)
	}

	// create aliens
	gs.aliens = generateAliens(numAliens, gs.listCities, rng)

	// return game state
	return &gs
}

func (gs *gameState) run(iteration int) {
	color.Green("----- Game has started! -----")
	fmt.Println()

	for i := 0; i < iteration; i++ {
		// test if game has ended
		if gs.deadAliens == len(gs.aliens) {
			color.Yellow(" All aliens have died!")
			break
		}
		// game iteration
		gs.iteration()
	}
}

// iteration goes through one iteration of the game
func (gs *gameState) iteration() {
	// stepping
	fmt.Printf("-> turn %d\n", gs.step)
	gs.step++
	if slow {
		time.Sleep(1 * time.Second)
	}

	for _, currentAlien := range gs.aliens {
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
		whichWayIdx := gs.rng.Intn(atCity.numLinks)
	Loop:
		for {
			switch whichWayIdx {
			case 0:
				// if empty direction, go to the next
				if goTo := atCity.north; goTo == "" {
					whichWayIdx = (whichWayIdx + 1) % 4
				} else {
					color.Blue(" alien %s moved from %s to %s", currentAlien.name, atCity.name, goTo)
					atCity.aliens = atCity.aliens[:0]                                     // remove alien from city
					currentAlien.atCity = gs.cities[goTo]                                 // set city of alien
					gs.cities[goTo].aliens = append(gs.cities[goTo].aliens, currentAlien) // add alien to city
					break Loop
				}

			case 1:
				if goTo := atCity.west; goTo == "" {
					whichWayIdx = (whichWayIdx + 1) % 4
				} else {
					color.Blue(" alien %s moved from %s to %s", currentAlien.name, atCity.name, goTo)
					atCity.aliens = atCity.aliens[:0]
					currentAlien.atCity = gs.cities[goTo]
					gs.cities[goTo].aliens = append(gs.cities[goTo].aliens, currentAlien)
					break Loop
				}

			case 2:
				if goTo := atCity.east; goTo == "" {
					whichWayIdx = (whichWayIdx + 1) % 4
				} else {
					color.Blue(" alien %s moved from %s to %s", currentAlien.name, atCity.name, goTo)
					atCity.aliens = atCity.aliens[:0]
					currentAlien.atCity = gs.cities[goTo]
					gs.cities[goTo].aliens = append(gs.cities[goTo].aliens, currentAlien)
					break Loop
				}

			case 3:
				if goTo := atCity.south; goTo == "" {
					whichWayIdx = (whichWayIdx + 1) % 4
				} else {
					color.Blue(" alien %s moved from %s to %s", currentAlien.name, atCity.name, goTo)
					atCity.aliens = atCity.aliens[:0]
					currentAlien.atCity = gs.cities[goTo]
					gs.cities[goTo].aliens = append(gs.cities[goTo].aliens, currentAlien)
					break Loop
				}
			}
		} // endfor

		// is there more than one alien on the city? -> fight
		atCity = currentAlien.atCity
		if len(atCity.aliens) == 2 {
			color.Red(" alien: %s and %s died in a fight.\n", atCity.aliens[0].name, atCity.aliens[1].name)
			// aliens kill each other
			atCity.aliens[0].dead = true
			atCity.aliens[1].dead = true
			gs.deadAliens += 2
			// destroy city
			atCity.destroyed = true
			color.Magenta(" city: %s has been destroyed.\n", atCity.name)
			// destroy roads
			if atCity.north != "" {
				gs.cities[atCity.north].south = ""
				gs.cities[atCity.north].numLinks--
			}
			if atCity.west != "" {
				gs.cities[atCity.west].east = ""
				gs.cities[atCity.west].numLinks--
			}
			if atCity.east != "" {
				gs.cities[atCity.east].west = ""
				gs.cities[atCity.east].numLinks--
			}
			if atCity.south != "" {
				gs.cities[atCity.south].north = ""
				gs.cities[atCity.south].numLinks--
			}
		}

	} // end for aliens
}
