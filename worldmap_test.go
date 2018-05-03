package main

import (
	"bytes"
	"math/rand"
	"testing"
)

var rng *rand.Rand

// I hate writing tests
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

func TestGameShouldAlwaysFinish(t *testing.T) {

}
