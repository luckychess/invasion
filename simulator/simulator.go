package simulator

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/luckychess/invasion/world"
)

const (
	simulatorSteps = 10
)

type simulator struct {
	worldMap    world.WorldMap
	rand        *rand.Rand
	stepsCount  uint32
	aliensCount uint32
}

func InitSimulation(worldMap *world.WorldMap, aliens uint32) simulator {
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	return simulator{worldMap: *worldMap, rand: rand, stepsCount: simulatorSteps, aliensCount: aliens}
}

func (sim *simulator) Simulate() {
	sim.unleashAliens()
	for i := 0; i < int(sim.stepsCount); i++ {
		if len(sim.worldMap.Aliens) == 0 {
			log.Println("No more aliens to fight, stopping simulation")
			break
		}
		for _, alien := range sim.worldMap.Aliens {
			sim.worldMap.MoveAlien(alien, sim.rand)
		}
		sim.fightAliens()
	}
}

func (sim *simulator) StopSimulation() {
	log.Println("=== Simulation finished ===")
	for name, city := range sim.worldMap.Cities {
		cityOutput := fmt.Sprintf("%s ", name)
		directions := city.GetDirections()
		for _, dir := range directions {
			neighbour, err := city.GetNeighbour(dir)
			if err == nil {
				cityOutput += fmt.Sprintf("%s=%s ", dir, neighbour)
			} else {
				log.Println(err)
			}
		}
		log.Print(cityOutput)
	}
}

func (sim *simulator) fightAliens() {
	for city := range sim.worldMap.Cities {
		sim.worldMap.DestroyCity(city)
	}
}

func (sim *simulator) unleashAliens() {
	for i := 0; i < int(sim.aliensCount); i++ {
		name := sim.getRandomName()
		city := sim.getRandomCity()
		log.Printf("Unleashing alien %s into city %s", name, city)
		alien := world.Alien{Name: name, City: city}
		sim.worldMap.AddAlien(&alien)
	}
}

func (sim *simulator) getRandomName() string {
	const nameLength = 8
	name := ""
	for i := 0; i < nameLength; i++ {
		name += string(rune('a' + sim.rand.Intn(26)))
	}
	return name
}

func (sim *simulator) getRandomCity() string {
	// pretty inefficient O(n) implementation
	// unfortunately there is no easy way to get a random element
	// of map in Golang
	// O(1) is possible but more complicated
	// e.g. requires creating and maintaining
	// a separate slice with keys
	keys := make([]string, 0, len(sim.worldMap.Cities))
	for k := range sim.worldMap.Cities {
		keys = append(keys, k)
	}
	return keys[sim.rand.Intn(len(keys))]
}
