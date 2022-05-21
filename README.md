# Saga test task description: Alien Invasion

## Task description
Mad aliens are about to invade the earth and you are tasked with simulating the invasion.

You are given a map containing the names of cities in the non-existent world of X. The map is in a file, with one city per line. The city name is first, followed by 1-4 directions (north, south east, or west). Each one represents a road to another city that lies in that direction.

For example:
Foo north=Bar west=Baz south=Qu-ux
Bar south=Foo west=Bee

The city and each of the pairs are separated by a single space, and the directions are separated from their respective cities with an equals (=) sign.

You should create N aliens, where N is specified as a command-line argument.

These aliens start out at random places on the map, and wander around randomly, following links. Each iteration, the aliens can travel in any of the directions leading out of a city. In our example above, an alien that starts at Foo can go north to Bar, west to Baz, or south to Qu-ux.

When two aliens end up in the same place, they fight, and in the process kill each other and destroy the city. When a city is destroyed, it is removed from the map, and so are any roads that lead into or out of it. In our example above, if Bar were destroyed the map would now be something like:
Foo west=Baz south=Qu-ux

Once a city is destroyed, aliens can no longer travel to or through it. This may lead to aliens getting "trapped".

You should create a program that reads in the world map, creates N aliens, and unleashes them. The program should run until all the aliens have been destroyed, or each alien has moved at least 10,000 times. When two aliens fight, print out a message like:
Bar has been destroyed by alien 10 and alien 34!

(If you want to give them names, you may, but it is not required.) Once the program has finished, it should print out whatever is left of the world in the same format as the input file. Feel free to make assumptions (for example, that the city names will never contain numeric characters), but please add comments or assertions describing the assumptions you are making.

## Solution description
Project is fully written in Go and uses benefits of Go modules. Current version is 0.0.1 and it can be downloaded as a module with `go get github.com/luckychess/invasion@v0.0.1` command.

Package `main` contains program entry point, performs input data parsing and processing, starts simulation and prints simulation results. Package `world` contains representation of the world map. In `simulator` package you can find the simulation logic itself. Both `simulator` and `world` packages contain unit tests. Code from `main` package remains uncovered by tests which is one of possible project improvements.

Mocks for `battlefield.go` are generated with `GoMock`.

## How to build and run
To build this project you need Go 1.18 installed in your system. Use `make build` command to build the project. As an option you can build Docker image with `make docker`, then execute `make docker-run` and build project with `make build` inside running container.

In both host and docker environments you can use `make test` command.

There are 2 sample world maps provided with this project in `sample` directory. With `make run-simple` or `make run-big` you can test this project in 2 different configurations: 2 aliens and 5 cities map and 300 aliens and 128 cities one.

To manually run this solution you need to pass 2 arguments to the executable file. First argument sets amount of aliens and second sets path to the map file. E.g. `./invasion 100 sample/input_big.txt`.
