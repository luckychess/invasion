package simulator

import "github.com/luckychess/invasion/world"

type simulator struct {
	worldMap world.WorldMap
}

func InitSimulation(worldMap *world.WorldMap) simulator {
	return simulator{worldMap: *worldMap}
}

func (sim *simulator) Simulate() {}

func (sim *simulator) StopSimulation() {}
