package simulator

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/luckychess/invasion/world"
)

const (
	simulatorSteps = 10000
)

type simulator struct {
	worldMap    world.WorldMap
	rng         *rand.Rand
	stepsCount  uint32
	aliensCount uint32
}

// InitSimulation creates an empty world map from given parameters.
func InitSimulation(worldMap world.WorldMap, rng *rand.Rand, aliens uint32) simulator {
	return simulator{worldMap: worldMap, rng: rng, stepsCount: simulatorSteps, aliensCount: aliens}
}

// Simulate performs the invasion simulation. At the beginning it creates and randomly spreads
// aliens along the world map. This follows by a fight check: if there are 2 or more aliens
// in the same city, this city is destroyed together with all the aliens in it.
// Then for the given amount of simulation steps (currently hardcoded to 10000)
// Simulate moves each alien in a random direction and performs a new fight check
// AFTER all aliens have moved. This means that during the simulation step it's possible to
// exist more than one alien in the same city without the fight if at the end of the simulation step
// less than 2 aliens remain in the city.
func (sim *simulator) Simulate() {
	sim.unleashAliens()
	for i := 0; i < int(sim.stepsCount); i++ {
		if len(sim.worldMap.GetAliens()) == 0 {
			log.Println("No more aliens to fight, stopping simulation")
			break
		}
		for _, alien := range sim.worldMap.GetAliens() {
			sim.worldMap.MoveAlien(alien, sim.rng)
		}
		sim.fightAliens()
	}
}

// StopSimulation returns status of the world in the same format as input data.
func (sim *simulator) StopSimulation() string {
	result := ""
	result += "=== Simulation finished ===\n"
	for name, city := range sim.worldMap.GetCities() {
		cityOutput := fmt.Sprintf("%s ", name)
		directions := city.GetDirections()
		for _, dir := range directions {
			neighbour, err := city.GetNeighbour(dir)
			if err == nil {
				cityOutput += fmt.Sprintf("%s=%s ", dir, neighbour)
			} else {
				result += err.Error() + "\n"
			}
		}
		result += cityOutput + "\n"
	}
	return result
}

func (sim *simulator) fightAliens() {
	for city := range sim.worldMap.GetCities() {
		sim.worldMap.DestroyCity(city)
	}
}

func (sim *simulator) unleashAliens() {
	for i := 0; i < int(sim.aliensCount); i++ {
		name := sim.getRandomName()
		city, err := sim.getRandomCity()
		if err == nil {
			log.Printf("Unleashing alien %s into city %s", name, city)
			alien := world.Alien{Name: name, City: city}
			sim.worldMap.AddAlien(&alien)
		} else {
			log.Println(err)
		}
	}
	sim.fightAliens()
}

func (sim *simulator) getRandomName() string {
	const nameLength = 8
	name := ""
	for i := 0; i < nameLength; i++ {
		name += string(rune('a' + sim.rng.Intn(26)))
	}
	return name
}

func (sim *simulator) getRandomCity() (string, error) {
	// pretty inefficient O(n) implementation
	// unfortunately there is no easy way to get a random element
	// of map in Golang
	// O(1) is possible but more complicated
	// e.g. requires creating and maintaining
	// a separate slice with keys
	if len(sim.worldMap.GetCities()) == 0 {
		return "", fmt.Errorf("there are no cities in the world")
	}
	keys := make([]string, 0, len(sim.worldMap.GetCities()))
	for k := range sim.worldMap.GetCities() {
		keys = append(keys, k)
	}
	return keys[sim.rng.Intn(len(keys))], nil
}
