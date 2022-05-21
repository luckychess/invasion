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
	assert.Assert(t, city.Name == name)
	assert.Assert(t, city.East == nil)
	assert.Assert(t, city.North == nil)
	assert.Assert(t, city.West == nil)
	assert.Assert(t, city.South == nil)
}

func TestMultipleCitiesMap(t *testing.T) {

	wm := createSimpleMap()
	// now let's check that all the cities are saved according to the scheme above
	// Berlin
	assert.Assert(t, wm.GetCities()[cities[4]].East == nil)
	assert.Assert(t, wm.GetCities()[cities[4]].North == nil)
	assert.Assert(t, wm.GetCities()[cities[4]].West.Name == cities[1])
	assert.Assert(t, wm.GetCities()[cities[4]].South == nil)

	// Cologne
	assert.Assert(t, wm.GetCities()[cities[1]].East.Name == cities[4])
	assert.Assert(t, wm.GetCities()[cities[1]].North == nil)
	assert.Assert(t, wm.GetCities()[cities[1]].West == nil)
	assert.Assert(t, wm.GetCities()[cities[1]].South.Name == cities[2])

	// Frankfurt
	assert.Assert(t, wm.GetCities()[cities[2]].East.Name == cities[6])
	assert.Assert(t, wm.GetCities()[cities[2]].North.Name == cities[1])
	assert.Assert(t, wm.GetCities()[cities[2]].West.Name == cities[5])
	assert.Assert(t, wm.GetCities()[cities[2]].South.Name == cities[0])

	// Strasbourg
	assert.Assert(t, wm.GetCities()[cities[5]].East.Name == cities[2])
	assert.Assert(t, wm.GetCities()[cities[5]].North == nil)
	assert.Assert(t, wm.GetCities()[cities[5]].West == nil)
	assert.Assert(t, wm.GetCities()[cities[5]].South == nil)

	// Nuremberg
	assert.Assert(t, wm.GetCities()[cities[6]].East == nil)
	assert.Assert(t, wm.GetCities()[cities[6]].North == nil)
	assert.Assert(t, wm.GetCities()[cities[6]].West.Name == cities[2])
	assert.Assert(t, wm.GetCities()[cities[6]].South.Name == cities[3])

	// Heidelberg
	assert.Assert(t, wm.GetCities()[cities[0]].East.Name == cities[3])
	assert.Assert(t, wm.GetCities()[cities[0]].North.Name == cities[2])
	assert.Assert(t, wm.GetCities()[cities[0]].West == nil)
	assert.Assert(t, wm.GetCities()[cities[0]].South == nil)

	// Munich
	assert.Assert(t, wm.GetCities()[cities[3]].East == nil)
	assert.Assert(t, wm.GetCities()[cities[3]].North.Name == cities[6])
	assert.Assert(t, wm.GetCities()[cities[3]].West.Name == cities[0])
	assert.Assert(t, wm.GetCities()[cities[3]].South == nil)

	// Regensburg
	assert.Assert(t, wm.GetCities()[cities[7]].East == nil)
	assert.Assert(t, wm.GetCities()[cities[7]].North.Name == cities[8])
	assert.Assert(t, wm.GetCities()[cities[7]].West == nil)
	assert.Assert(t, wm.GetCities()[cities[7]].South == nil)

	// Leipzig
	assert.Assert(t, wm.GetCities()[cities[8]].East == nil)
	assert.Assert(t, wm.GetCities()[cities[8]].North == nil)
	assert.Assert(t, wm.GetCities()[cities[8]].West == nil)
	assert.Assert(t, wm.GetCities()[cities[8]].South.Name == cities[7])
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
	assert.Assert(t, wm.GetCities()[cities[6]].West == nil)
	assert.Assert(t, wm.GetCities()[cities[1]].South == nil)
	assert.Assert(t, wm.GetCities()[cities[5]].East == nil)
	assert.Assert(t, wm.GetCities()[cities[0]].North == nil)
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
	city := City{Name: "Geneva", Aliens: make(map[string]bool)}
	// east direction
	_, err := city.GetNeighbour("east")
	assert.Error(t, err, "no cities in east direction")
	eastCity := &City{Name: "Milan", East: nil, North: nil, West: &city, South: nil, Aliens: make(map[string]bool)}
	city.East = eastCity
	milan, _ := city.GetNeighbour("east")
	assert.Assert(t, milan == eastCity.Name)
	// north direction
	_, err = city.GetNeighbour("north")
	assert.Error(t, err, "no cities in north direction")
	northCity := &City{Name: "Bern", East: nil, North: nil, West: &city, South: nil, Aliens: make(map[string]bool)}
	city.North = northCity
	bern, _ := city.GetNeighbour("north")
	assert.Assert(t, bern == northCity.Name)
	// west direction
	_, err = city.GetNeighbour("west")
	assert.Error(t, err, "no cities in west direction")
	westCity := &City{Name: "Lyon", East: nil, North: nil, West: &city, South: nil, Aliens: make(map[string]bool)}
	city.West = westCity
	lyon, _ := city.GetNeighbour("west")
	assert.Assert(t, lyon == westCity.Name)
	// south direction
	_, err = city.GetNeighbour("south")
	assert.Error(t, err, "no cities in south direction")
	southCity := &City{Name: "Marseille", East: nil, North: nil, West: &city, South: nil, Aliens: make(map[string]bool)}
	city.South = southCity
	marseille, _ := city.GetNeighbour("south")
	assert.Assert(t, marseille == southCity.Name)

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
