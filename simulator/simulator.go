package simulator

import "github.com/luckychess/invasion/world"

type simulator struct {
	worldMap world.WorldMap
}

func InitSimulation(fileName string) simulator {
	lines := readFile(fileName)
	worldMap := parseInput(lines)
	return simulator{worldMap: worldMap}
}

func (sim *simulator) Simulate() {}

func (sim *simulator) StopSimulation() {}

func readFile(fileName string) []string {
	return make([]string, 0)
}

func parseInput(lines []string) world.WorldMap {
	/*
		Foo north=Bar west=Baz south=Qu-ux
		Bar south=Foo west=Bee
	*/
	return world.InitWorldMap()
}
