package world

import (
	"fmt"
	"log"
	"math/rand"
)

type Alien struct {
	Name string
	City string
}

type City struct {
	name   string
	east   *City
	north  *City
	west   *City
	south  *City
	aliens map[string]bool
}

func (c *City) GetDirections() []string {
	directions := make([]string, 0)
	if c.east != nil {
		directions = append(directions, "east")
	}
	if c.north != nil {
		directions = append(directions, "north")
	}
	if c.west != nil {
		directions = append(directions, "west")
	}
	if c.south != nil {
		directions = append(directions, "south")
	}
	return directions
}

func (c *City) GetNeighbour(direction string) (string, error) {
	switch direction {
	case "east":
		if c.east != nil {
			return c.east.name, nil
		}
	case "north":
		if c.north != nil {
			return c.north.name, nil
		}
	case "west":
		if c.west != nil {
			return c.west.name, nil
		}
	case "south":
		if c.south != nil {
			return c.south.name, nil
		}
	default:
		return "", fmt.Errorf("wrong direction %s", direction)
	}
	return "", fmt.Errorf("no cities in %s direction", direction)
}

type WorldMap struct {
	Cities map[string]*City
	Aliens map[string]*Alien
}

func InitWorldMap() WorldMap {
	worldMap := WorldMap{}
	worldMap.Cities = make(map[string]*City)
	worldMap.Aliens = make(map[string]*Alien)
	return worldMap
}

func (m *WorldMap) AddCity(name string, east string, north string, west string, south string) {
	city := City{name: name, aliens: make(map[string]bool)}
	if east != "" {
		eastCity := m.Cities[east]
		if eastCity == nil {
			eastCity = &City{name: east, east: nil, north: nil, west: &city, south: nil, aliens: make(map[string]bool)}
		}
		city.east = eastCity
		eastCity.west = &city
		m.Cities[east] = eastCity
	}
	if north != "" {
		northCity := m.Cities[north]
		if northCity == nil {
			northCity = &City{name: north, east: nil, north: nil, west: nil, south: &city, aliens: make(map[string]bool)}
		}
		city.north = northCity
		northCity.south = &city
		m.Cities[north] = northCity
	}
	if west != "" {
		westCity := m.Cities[west]
		if westCity == nil {
			westCity = &City{name: west, east: &city, north: nil, west: nil, south: nil, aliens: make(map[string]bool)}
		}
		city.west = westCity
		westCity.east = &city
		m.Cities[west] = westCity
	}
	if south != "" {
		southCity := m.Cities[south]
		if southCity == nil {
			southCity = &City{name: south, east: nil, north: &city, west: nil, south: nil, aliens: make(map[string]bool)}
		}
		city.south = southCity
		southCity.north = &city
		m.Cities[south] = southCity
	}

	m.Cities[name] = &city
}

func (m *WorldMap) AddAlien(alien *Alien) error {
	if m.Cities[alien.City] == nil {
		return (fmt.Errorf("trying to unleash an alien %s into non-existing city %s", alien.Name, alien.City))
	}
	m.Aliens[alien.Name] = alien
	m.Cities[alien.City].aliens[alien.Name] = true
	return nil
}

func (m *WorldMap) MoveAlien(alien *Alien, rng *rand.Rand) {
	city := m.Cities[alien.City]
	directions := city.GetDirections()
	if len(directions) > 0 {
		direction := directions[rng.Intn(len(directions))]
		newCity, err := city.GetNeighbour(direction)
		if err == nil {
			alien.City = newCity
			delete(city.aliens, alien.Name)
			m.Cities[alien.City].aliens[alien.Name] = true
		} else {
			log.Println(err)
		}
	}
}

func (m *WorldMap) DestroyCity(cityToDestroy string) {
	city := m.Cities[cityToDestroy]
	if len(city.aliens) > 1 {
		if city.east != nil {
			city.east.west = nil
		}
		if city.north != nil {
			city.north.south = nil
		}
		if city.west != nil {
			city.west.east = nil
		}
		if city.south != nil {
			city.south.north = nil
		}
		delete(m.Cities, city.name)
		for alien := range city.aliens {
			delete(m.Aliens, alien)
		}
		aliens := ""
		for alien := range city.aliens {
			aliens += alien + " "
		}
		log.Printf("%s has been destroyed by aliens %s", cityToDestroy, aliens)
		city.aliens = nil
	}
}
