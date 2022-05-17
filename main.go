package main

import (
	"log"
	"os"
	"strings"

	"github.com/luckychess/invasion/simulator"
	"github.com/luckychess/invasion/world"
)

// program entry point
func main() {
	totalAliens := 2
	lines := readFile("input.txt")
	worldMap := parseInput(lines)
	simulator := simulator.InitSimulation(&worldMap, uint32(totalAliens))
	simulator.Simulate()
	simulator.StopSimulation()
}

func readFile(fileName string) []string {
	// read all the file at once
	// it's probably more efficient to read line by line and
	// parse immediately but current approach seems to be simpler
	// expect lines are separated with \n newline separator
	inputBytes, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalf("Error happened when trying to read file %s: %s\n", fileName, err)
	}
	inputLines := strings.Split(string(inputBytes), "\n")
	for i, line := range inputLines {
		log.Printf("Line %d: %s\n", i, line)
	}
	return inputLines
}

func parseInput(lines []string) world.WorldMap {
	/*
		Sample file data:
		------------------
		Foo north=Bar west=Baz south=Qu-ux
		Bar south=Foo west=Bee
	*/
	worldMap := world.InitWorldMap()

	for _, line := range lines {
		// expect every line data is separated by spaces
		words := strings.Split(line, " ")
		// first word is always a city name (shouldn't contain spaces)
		newCity := words[0]
		var east, north, west, south string
		// expect city1=city2 pairs, up to 4 pairs, one pair for every direction {east, north, west, south}
		for i := 1; i < len(words); i++ {
			road := strings.Split(words[i], "=")
			if len(road) != 2 {
				log.Fatalf("Error parsing input data: expected city1=city2 format but got %s\n", words[i])
			}
			direction, city := road[0], road[1]
			switch direction {
			case "east":
				east = city
			case "north":
				north = city
			case "west":
				west = city
			case "south":
				south = city
			default:
				log.Fatalf("Error parsing input data: wrong direction %s\n", direction)
			}

		}
		worldMap.AddCity(newCity, east, north, west, south)
	}
	return worldMap
}
