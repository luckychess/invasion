package world

import (
	"testing"

	"gotest.tools/assert"
)

var cities = []string{"Heidelberg", "Cologne", "Frankfurt", "Munich", "Berlin", "Strasbourg", "Nuremberg", "Regensburg", "Leipzig"}

func TestInitWorldMap(t *testing.T) {
	wm := InitWorldMap()
	// no cities, no aliens
	assert.Assert(t, len(wm.cities) == 0)
	assert.Assert(t, len(wm.aliens) == 0)
}

func TestSingleCityMap(t *testing.T) {
	wm := InitWorldMap()
	name := "Heidelberg"
	wm.AddCity(name, "", "", "", "")
	// only 1 city and no aliens so far
	assert.Assert(t, len(wm.cities) == 1)
	assert.Assert(t, len(wm.aliens) == 0)
	// still the same city without neighbours
	city := wm.cities[name]
	assert.Assert(t, city.name == name)
	assert.Assert(t, city.east == nil)
	assert.Assert(t, city.north == nil)
	assert.Assert(t, city.west == nil)
	assert.Assert(t, city.south == nil)
}

func TestMultipleCitiesMap(t *testing.T) {

	wm := createSimpleMap()
	// now let's check that all the cities are saved according to the scheme above
	// Berlin
	assert.Assert(t, wm.cities[cities[4]].east == nil)
	assert.Assert(t, wm.cities[cities[4]].north == nil)
	assert.Assert(t, wm.cities[cities[4]].west.name == cities[1])
	assert.Assert(t, wm.cities[cities[4]].south == nil)

	// Cologne
	assert.Assert(t, wm.cities[cities[1]].east.name == cities[4])
	assert.Assert(t, wm.cities[cities[1]].north == nil)
	assert.Assert(t, wm.cities[cities[1]].west == nil)
	assert.Assert(t, wm.cities[cities[1]].south.name == cities[2])

	// Frankfurt
	assert.Assert(t, wm.cities[cities[2]].east.name == cities[6])
	assert.Assert(t, wm.cities[cities[2]].north.name == cities[1])
	assert.Assert(t, wm.cities[cities[2]].west.name == cities[5])
	assert.Assert(t, wm.cities[cities[2]].south.name == cities[0])

	// Strasbourg
	assert.Assert(t, wm.cities[cities[5]].east.name == cities[2])
	assert.Assert(t, wm.cities[cities[5]].north == nil)
	assert.Assert(t, wm.cities[cities[5]].west == nil)
	assert.Assert(t, wm.cities[cities[5]].south == nil)

	// Nuremberg
	assert.Assert(t, wm.cities[cities[6]].east == nil)
	assert.Assert(t, wm.cities[cities[6]].north == nil)
	assert.Assert(t, wm.cities[cities[6]].west.name == cities[2])
	assert.Assert(t, wm.cities[cities[6]].south.name == cities[3])

	// Heidelberg
	assert.Assert(t, wm.cities[cities[0]].east.name == cities[3])
	assert.Assert(t, wm.cities[cities[0]].north.name == cities[2])
	assert.Assert(t, wm.cities[cities[0]].west == nil)
	assert.Assert(t, wm.cities[cities[0]].south == nil)

	// Munich
	assert.Assert(t, wm.cities[cities[3]].east == nil)
	assert.Assert(t, wm.cities[cities[3]].north.name == cities[6])
	assert.Assert(t, wm.cities[cities[3]].west.name == cities[0])
	assert.Assert(t, wm.cities[cities[3]].south == nil)

	// Regensburg
	assert.Assert(t, wm.cities[cities[7]].east == nil)
	assert.Assert(t, wm.cities[cities[7]].north.name == cities[8])
	assert.Assert(t, wm.cities[cities[7]].west == nil)
	assert.Assert(t, wm.cities[cities[7]].south == nil)

	// Leipzig
	assert.Assert(t, wm.cities[cities[8]].east == nil)
	assert.Assert(t, wm.cities[cities[8]].north == nil)
	assert.Assert(t, wm.cities[cities[8]].west == nil)
	assert.Assert(t, wm.cities[cities[8]].south.name == cities[7])
}

func TestAddAlien(t *testing.T) {
	wm := InitWorldMap()
	wm.AddCity("Zurich", "", "Frankfurt", "", "Milan")
	assert.Assert(t, wm.AddAlien(&Alien{name: "The Evil", city: "Zurich"}) == nil)
	// Alien should be added as expected
	assert.Assert(t, wm.aliens["The Evil"].city == "Zurich")
	assert.Assert(t, wm.aliens["The Evil"].name == "The Evil")

	// Expect non-zero error when attempt to invade non-existing city
	assert.Assert(t, wm.AddAlien(&Alien{name: "Not very clever", city: "Moscow"}) != nil)
}

func TestDestroyCity(t *testing.T) {
	wm := createSimpleMap()
	alien1 := &Alien{name: "Green dude", city: cities[2]}
	alien2 := &Alien{name: "Earth invader", city: cities[2]}
	wm.AddAlien(alien1)
	wm.AddAlien(alien2)
	wm.DestroyCity(alien1, alien2)
	// Frankfurt should be destroyed now and aliens should be dead
	assert.Assert(t, wm.aliens[alien1.name] == nil)
	assert.Assert(t, wm.aliens[alien2.name] == nil)
	assert.Assert(t, wm.cities[alien1.name] == nil)
	// Frankfurt connections are also destroyed now
	assert.Assert(t, wm.cities[cities[6]].west == nil)
	assert.Assert(t, wm.cities[cities[1]].south == nil)
	assert.Assert(t, wm.cities[cities[5]].east == nil)
	assert.Assert(t, wm.cities[cities[0]].north == nil)
}

func createSimpleMap() *WorldMap {
	// the map how it's supposed to look like (check first letters; *slightly* different to the real life)
	/*
			C - B   L
			|		|
		S - F - N - R
			|	|
			H - M
	*/
	wm := InitWorldMap()
	wm.AddCity(cities[4], "", "", cities[1], "")
	wm.AddCity(cities[1], cities[4], "", "", cities[2])
	wm.AddCity(cities[2], cities[6], cities[1], cities[5], cities[0])
	wm.AddCity(cities[5], cities[2], "", "", "")
	wm.AddCity(cities[6], "", "", cities[2], cities[3])
	wm.AddCity(cities[0], cities[3], cities[2], "", "")
	wm.AddCity(cities[3], "", cities[6], cities[0], "")
	wm.AddCity(cities[7], "", cities[8], "", "")

	return &wm
}