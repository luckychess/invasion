package simulator

type simulator struct {
	worldMap WorldMap
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

func parseInput(lines []string) WorldMap {
	/*
		Foo north=Bar west=Baz south=Qu-ux
		Bar south=Foo west=Bee
	*/
	return initWorldMap()
}
