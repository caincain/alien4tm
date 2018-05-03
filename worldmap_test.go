package main

import (
	"bytes"
	"math/rand"
	"testing"
)

var rng *rand.Rand

func init() {
	randSource := rand.NewSource(42)
	rng = rand.New(randSource)
}

func TestMapGenerationAndParsing(t *testing.T) {
	for i := 2; i < 50; i++ {
		// generate a map
		_, resList := generateMap(i, rng)

		// create a map file from it
		var b bytes.Buffer
		printWorldForFile(resList, &b)

		// parse it
		cities := parseMap(&b)

		if len(cities) != i {
			t.Fatalf("map generation and parsing did not work for size %d", i)
		}
	}
}

func TestFindWorldCoordinates(t *testing.T) {

	// generate a map
	resMap, _ := generateMap(20, rng)
	_, _, _, _, err := findWorldCoordinates(resMap)

	if err != nil {
		t.Fatalf("finding coordinates didn't work")
	}

}
