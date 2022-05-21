package main

import (
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/luckychess/invasion/simulator"
	"github.com/luckychess/invasion/world"
)

// program entry point
func main() {
	log.SetFlags(0)
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s <N> where N is amount of aliens", os.Args[0])
	}
	totalAliens, err := strconv.ParseUint(os.Args[1], 10, 32)
	if err != nil {
		log.Fatalf("Command line argument expected to be a non-negative number: %s", err)
	}
	lines := readFile("input.txt")
	worldMap := parseInput(lines)
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	simulator := simulator.InitSimulation(worldMap, rng, uint32(totalAliens))
	simulator.Simulate()
	simulationResult := simulator.StopSimulation()
	log.Print(simulationResult)
}

func readFile(fileName string) []string {
	// read all the file at once
	// it's probably more efficient to read line by line and
	// parse immediately but current approach seems to be simpler
	// expect lines are separated with \n newline separator
	inputBytes, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalf("Error happened when trying to read file %s: %s", fileName, err)
	}
	inputLines := strings.Split(string(inputBytes), "\n")
	for i, line := range inputLines {
		log.Printf("Line %d: %s", i, line)
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
				log.Fatalf("Error parsing input data: expected city1=city2 format but got %s", words[i])
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
				log.Fatalf("Error parsing input data: wrong direction %s", direction)
			}

		}
		worldMap.AddCity(newCity, east, north, west, south)
	}
	return worldMap
}
