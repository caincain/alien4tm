package main

import "testing"

func TestGameShouldAlwaysFinish(t *testing.T) {
	// This test is going to be a tad slow
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	// generate a map
	mapList, _ := generateMap(10, rng)

	state := newGame(mapList, 3, rng)
	state.run(10000)

	// just finishing this test is enough
}

func TestSimpleBridge(t *testing.T) {
	// 2 cities with A -> B
	cityA := &city{name: "A", north: "B", numLinks: 1}
	cityB := &city{name: "B"}
	cities := map[string]*city{"A": cityA, "B": cityB}
	listCities := []*city{cityA, cityB}
	// alien in A
	alien1 := &alien{name: "1", atCity: cityA}
	cityA.aliens = []*alien{alien1}
	// create state
	state := gameState{cities: cities, listCities: listCities, aliens: []*alien{alien1}, rng: rng}
	state.run(1)

	if len(cityA.aliens) != 0 || len(cityB.aliens) != 1 || alien1.atCity != cityB {
		t.Fatalf("Alien hasn't moved correctly")
	}
}

func TestStuck(t *testing.T) {
	// 2 cities with A <-!-B
	cityA := &city{name: "A", north: "B", numLinks: 1}
	cityB := &city{name: "B"}
	cities := map[string]*city{"A": cityA, "B": cityB}
	listCities := []*city{cityA, cityB}
	// alien in B
	alien1 := &alien{name: "1", atCity: cityB}
	cityB.aliens = []*alien{alien1}
	aliens := []*alien{alien1}
	// create state
	state := gameState{cities: cities, listCities: listCities, aliens: aliens, rng: rng}
	state.run(1)

	if len(cityA.aliens) != 0 || len(cityB.aliens) != 1 || alien1.atCity != cityB {
		t.Fatalf("Alien has moved when it shouldn't have")
	}
}

func TestTwoAliens(t *testing.T) {
	// 2 cities with A <-> B
	cityA := &city{name: "A", north: "B", numLinks: 1}
	cityB := &city{name: "B", south: "A", numLinks: 1}
	cities := map[string]*city{"A": cityA, "B": cityB}
	listCities := []*city{cityA, cityB}
	// aliens in both A and B
	alien1 := &alien{name: "1", atCity: cityA}
	alien2 := &alien{name: "2", atCity: cityB}
	cityA.aliens = []*alien{alien1}
	cityB.aliens = []*alien{alien2}
	aliens := []*alien{alien1, alien2}
	// create state
	state := gameState{cities: cities, listCities: listCities, aliens: aliens, rng: rng}
	state.run(1)

	allGood := true
	allGood = allGood || cityA.destroyed
	allGood = allGood || cityB.destroyed
	allGood = allGood || state.deadAliens == 2

	if !allGood {
		t.Fatalf("simple state has failed")
	}
}
