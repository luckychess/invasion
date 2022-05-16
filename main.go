package main

import (
	"github.com/luckychess/invasion/simulator"
	"github.com/luckychess/invasion/world"
)

// program entry point
func main() {
	lines := readFile("input.txt")
	worldMap := parseInput(lines)
	simulator := simulator.InitSimulation(&worldMap)
	simulator.Simulate()
	simulator.StopSimulation()
}

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
