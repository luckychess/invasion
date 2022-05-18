package world

import (
	"math/rand"
	"testing"

	"gotest.tools/assert"
)

var cities = []string{"Heidelberg", "Cologne", "Frankfurt", "Munich", "Berlin", "Strasbourg", "Nuremberg", "Regensburg", "Leipzig"}

func TestInitWorldMap(t *testing.T) {
	wm := InitWorldMap()
	// no cities, no aliens
	assert.Assert(t, len(wm.GetCities()) == 0)
	assert.Assert(t, len(wm.GetAliens()) == 0)
}

func TestSingleCityMap(t *testing.T) {
	wm := InitWorldMap()
	name := "Heidelberg"
	wm.AddCity(name, "", "", "", "")
	// only 1 city and no aliens so far
	assert.Assert(t, len(wm.GetCities()) == 1)
	assert.Assert(t, len(wm.GetAliens()) == 0)
	// still the same city without neighbours
	city := wm.GetCities()[name]
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
	assert.Assert(t, wm.GetCities()[cities[4]].east == nil)
	assert.Assert(t, wm.GetCities()[cities[4]].north == nil)
	assert.Assert(t, wm.GetCities()[cities[4]].west.name == cities[1])
	assert.Assert(t, wm.GetCities()[cities[4]].south == nil)

	// Cologne
	assert.Assert(t, wm.GetCities()[cities[1]].east.name == cities[4])
	assert.Assert(t, wm.GetCities()[cities[1]].north == nil)
	assert.Assert(t, wm.GetCities()[cities[1]].west == nil)
	assert.Assert(t, wm.GetCities()[cities[1]].south.name == cities[2])

	// Frankfurt
	assert.Assert(t, wm.GetCities()[cities[2]].east.name == cities[6])
	assert.Assert(t, wm.GetCities()[cities[2]].north.name == cities[1])
	assert.Assert(t, wm.GetCities()[cities[2]].west.name == cities[5])
	assert.Assert(t, wm.GetCities()[cities[2]].south.name == cities[0])

	// Strasbourg
	assert.Assert(t, wm.GetCities()[cities[5]].east.name == cities[2])
	assert.Assert(t, wm.GetCities()[cities[5]].north == nil)
	assert.Assert(t, wm.GetCities()[cities[5]].west == nil)
	assert.Assert(t, wm.GetCities()[cities[5]].south == nil)

	// Nuremberg
	assert.Assert(t, wm.GetCities()[cities[6]].east == nil)
	assert.Assert(t, wm.GetCities()[cities[6]].north == nil)
	assert.Assert(t, wm.GetCities()[cities[6]].west.name == cities[2])
	assert.Assert(t, wm.GetCities()[cities[6]].south.name == cities[3])

	// Heidelberg
	assert.Assert(t, wm.GetCities()[cities[0]].east.name == cities[3])
	assert.Assert(t, wm.GetCities()[cities[0]].north.name == cities[2])
	assert.Assert(t, wm.GetCities()[cities[0]].west == nil)
	assert.Assert(t, wm.GetCities()[cities[0]].south == nil)

	// Munich
	assert.Assert(t, wm.GetCities()[cities[3]].east == nil)
	assert.Assert(t, wm.GetCities()[cities[3]].north.name == cities[6])
	assert.Assert(t, wm.GetCities()[cities[3]].west.name == cities[0])
	assert.Assert(t, wm.GetCities()[cities[3]].south == nil)

	// Regensburg
	assert.Assert(t, wm.GetCities()[cities[7]].east == nil)
	assert.Assert(t, wm.GetCities()[cities[7]].north.name == cities[8])
	assert.Assert(t, wm.GetCities()[cities[7]].west == nil)
	assert.Assert(t, wm.GetCities()[cities[7]].south == nil)

	// Leipzig
	assert.Assert(t, wm.GetCities()[cities[8]].east == nil)
	assert.Assert(t, wm.GetCities()[cities[8]].north == nil)
	assert.Assert(t, wm.GetCities()[cities[8]].west == nil)
	assert.Assert(t, wm.GetCities()[cities[8]].south.name == cities[7])
}

func TestAddAlien(t *testing.T) {
	wm := InitWorldMap()
	wm.AddCity("Zurich", "", "Frankfurt", "", "Milan")
	assert.Assert(t, wm.AddAlien(&Alien{Name: "The Evil", City: "Zurich"}) == nil)
	// Alien should be added as expected
	assert.Assert(t, wm.GetAliens()["The Evil"].City == "Zurich")
	assert.Assert(t, wm.GetAliens()["The Evil"].Name == "The Evil")

	// Expect non-zero error when attempt to invade non-existing city
	assert.Assert(t, wm.AddAlien(&Alien{Name: "Not very clever", City: "Moscow"}) != nil)
}

func TestMoveAlien(t *testing.T) {
	// Initialize rng with seed=0 to make tests more predictable
	rng := rand.New(rand.NewSource(0))
	firstCity := "Prague"
	secondCity := "Amsterdam"
	wm := InitWorldMap()
	wm.AddCity(firstCity, "", "", "", "")
	alien := &Alien{Name: "Lazy cat", City: firstCity}
	wm.AddAlien(alien)
	// Trying to move alien but there are no cities to move to
	wm.MoveAlien(alien, rng)
	assert.Assert(t, alien.City == firstCity)
	// Now add another city connected with Prague
	wm.AddCity(secondCity, firstCity, "", "", "")
	// Now alien should be in Amsterdam
	wm.MoveAlien(alien, rng)
}

func TestDestroyCity(t *testing.T) {
	wm := createSimpleMap()
	alien1 := &Alien{Name: "Green dude", City: cities[2]}
	alien2 := &Alien{Name: "Earth invader", City: cities[2]}
	wm.AddAlien(alien1)
	wm.AddAlien(alien2)
	wm.DestroyCity(cities[2])
	// Frankfurt should be destroyed now and aliens should be dead
	assert.Assert(t, wm.GetAliens()[alien1.Name] == nil)
	assert.Assert(t, wm.GetAliens()[alien2.Name] == nil)
	assert.Assert(t, wm.GetCities()[alien1.Name] == nil)
	// Frankfurt connections are also destroyed now
	assert.Assert(t, wm.GetCities()[cities[6]].west == nil)
	assert.Assert(t, wm.GetCities()[cities[1]].south == nil)
	assert.Assert(t, wm.GetCities()[cities[5]].east == nil)
	assert.Assert(t, wm.GetCities()[cities[0]].north == nil)
}

func TestGetDirections(t *testing.T) {
	// it's not necessary to create a WorldMap instance here but it's easier to test this way
	wm := InitWorldMap()
	wm.AddCity("Hannover", "Berlin", "Hamburg", "Cologne", "Mainz")
	hannoverDirections := wm.GetCities()["Hannover"].GetDirections()
	assert.Assert(t, len(hannoverDirections) == 4)
	assert.Assert(t, hannoverDirections[0] == "east")
	assert.Assert(t, hannoverDirections[1] == "north")
	assert.Assert(t, hannoverDirections[2] == "west")
	assert.Assert(t, hannoverDirections[3] == "south")
	hamburgDirections := wm.GetCities()["Hamburg"].GetDirections()
	assert.Assert(t, len(hamburgDirections) == 1)
	assert.Assert(t, hamburgDirections[0] == "south")
}

func TestGetNeighbour(t *testing.T) {
	city := City{name: "Geneva", aliens: make(map[string]bool)}
	// east direction
	_, err := city.GetNeighbour("east")
	assert.Error(t, err, "no cities in east direction")
	eastCity := &City{name: "Milan", east: nil, north: nil, west: &city, south: nil, aliens: make(map[string]bool)}
	city.east = eastCity
	milan, _ := city.GetNeighbour("east")
	assert.Assert(t, milan == eastCity.name)
	// north direction
	_, err = city.GetNeighbour("north")
	assert.Error(t, err, "no cities in north direction")
	northCity := &City{name: "Bern", east: nil, north: nil, west: &city, south: nil, aliens: make(map[string]bool)}
	city.north = northCity
	bern, _ := city.GetNeighbour("north")
	assert.Assert(t, bern == northCity.name)
	// west direction
	_, err = city.GetNeighbour("west")
	assert.Error(t, err, "no cities in west direction")
	westCity := &City{name: "Lyon", east: nil, north: nil, west: &city, south: nil, aliens: make(map[string]bool)}
	city.west = westCity
	lyon, _ := city.GetNeighbour("west")
	assert.Assert(t, lyon == westCity.name)
	// south direction
	_, err = city.GetNeighbour("south")
	assert.Error(t, err, "no cities in south direction")
	southCity := &City{name: "Marseille", east: nil, north: nil, west: &city, south: nil, aliens: make(map[string]bool)}
	city.south = southCity
	marseille, _ := city.GetNeighbour("south")
	assert.Assert(t, marseille == southCity.name)

	_, err = city.GetNeighbour("wrong")
	assert.Error(t, err, "wrong direction wrong")
}

func createSimpleMap() WorldMap {
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

	return wm
}
