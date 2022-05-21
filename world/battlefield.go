package world

import (
	"fmt"
	"log"
	"math/rand"
)

// Alien structure contains name and current city name of alien.
type Alien struct {
	Name string
	City string
}

// City contains name of the city and pointers to cities in other directions.
// It also contains all aliens currently in the city.
type City struct {
	Name   string
	East   *City
	North  *City
	West   *City
	South  *City
	Aliens map[string]bool
}

// GetDirections is a helper function which returns a slice of
// possible directions to go from the city.
func (c *City) GetDirections() []string {
	directions := make([]string, 0)
	if c.East != nil {
		directions = append(directions, "east")
	}
	if c.North != nil {
		directions = append(directions, "north")
	}
	if c.West != nil {
		directions = append(directions, "west")
	}
	if c.South != nil {
		directions = append(directions, "south")
	}
	return directions
}

// GetNeighbour returns name of the city in given direction or error
// if no city in this direction exists.
func (c *City) GetNeighbour(direction string) (string, error) {
	switch direction {
	case "east":
		if c.East != nil {
			return c.East.Name, nil
		}
	case "north":
		if c.North != nil {
			return c.North.Name, nil
		}
	case "west":
		if c.West != nil {
			return c.West.Name, nil
		}
	case "south":
		if c.South != nil {
			return c.South.Name, nil
		}
	default:
		return "", fmt.Errorf("wrong direction %s", direction)
	}
	return "", fmt.Errorf("no cities in %s direction", direction)
}

// WorldMap interface describes actions available for the world.
type WorldMap interface {
	// GetCities returns all cities in the world.
	GetCities() map[string]*City
	// GetCities returns all aliens in the world.
	GetAliens() map[string]*Alien
	// AddCity adds a new city to the world and also creates or updates information
	// about neighbours of the given city.
	AddCity(name string, east string, north string, west string, south string)
	// AddAlien adds alien into the world.
	AddAlien(alien *Alien) error
	// MoveAlien moves given alien in a random direction
	// if there are directions to move.
	MoveAlien(alien *Alien, rng *rand.Rand)
	// Destroy city deletes city and all aliens in it if there are 2 or
	// more aliens in the city.
	DestroyCity(cityToDestroy string)
}

type worldMapImpl struct {
	Cities map[string]*City
	Aliens map[string]*Alien
}

// InitWorldMap creates an empty world map with no cities and aliens.
func InitWorldMap() WorldMap {
	worldMap := worldMapImpl{}
	worldMap.Cities = make(map[string]*City)
	worldMap.Aliens = make(map[string]*Alien)
	return &worldMap
}

func (m *worldMapImpl) GetCities() map[string]*City {
	return m.Cities
}

func (m *worldMapImpl) GetAliens() map[string]*Alien {
	return m.Aliens
}

func (m *worldMapImpl) AddCity(name string, east string, north string, west string, south string) {
	city := &City{Name: name, Aliens: make(map[string]bool)}
	if m.Cities[name] != nil {
		city = m.Cities[name]
	}
	if east != "" {
		eastCity := m.Cities[east]
		if eastCity == nil {
			eastCity = &City{Name: east, East: nil, North: nil, West: city, South: nil, Aliens: make(map[string]bool)}
		}
		city.East = eastCity
		eastCity.West = city
		m.Cities[east] = eastCity
	}
	if north != "" {
		northCity := m.Cities[north]
		if northCity == nil {
			northCity = &City{Name: north, East: nil, North: nil, West: nil, South: city, Aliens: make(map[string]bool)}
		}
		city.North = northCity
		northCity.South = city
		m.Cities[north] = northCity
	}
	if west != "" {
		westCity := m.Cities[west]
		if westCity == nil {
			westCity = &City{Name: west, East: city, North: nil, West: nil, South: nil, Aliens: make(map[string]bool)}
		}
		city.West = westCity
		westCity.East = city
		m.Cities[west] = westCity
	}
	if south != "" {
		southCity := m.Cities[south]
		if southCity == nil {
			southCity = &City{Name: south, East: nil, North: city, West: nil, South: nil, Aliens: make(map[string]bool)}
		}
		city.South = southCity
		southCity.North = city
		m.Cities[south] = southCity
	}

	m.Cities[name] = city
}

func (m *worldMapImpl) AddAlien(alien *Alien) error {
	if m.Cities[alien.City] == nil {
		return (fmt.Errorf("trying to unleash an alien %s into non-existing city %s", alien.Name, alien.City))
	}
	m.Aliens[alien.Name] = alien
	m.Cities[alien.City].Aliens[alien.Name] = true
	return nil
}

func (m *worldMapImpl) MoveAlien(alien *Alien, rng *rand.Rand) {
	city := m.Cities[alien.City]
	directions := city.GetDirections()
	if len(directions) > 0 {
		direction := directions[rng.Intn(len(directions))]
		newCity, err := city.GetNeighbour(direction)
		if err == nil {
			alien.City = newCity
			delete(city.Aliens, alien.Name)
			m.Cities[alien.City].Aliens[alien.Name] = true
		} else {
			log.Println(err)
		}
	}
}

func (m *worldMapImpl) DestroyCity(cityToDestroy string) {
	city := m.Cities[cityToDestroy]
	if len(city.Aliens) > 1 {
		if city.East != nil {
			city.East.West = nil
		}
		if city.North != nil {
			city.North.South = nil
		}
		if city.West != nil {
			city.West.East = nil
		}
		if city.South != nil {
			city.South.North = nil
		}
		delete(m.Cities, city.Name)
		for alien := range city.Aliens {
			delete(m.Aliens, alien)
		}
		aliens := ""
		for alien := range city.Aliens {
			aliens += alien + " "
		}
		log.Printf("%s has been destroyed by aliens %s", cityToDestroy, aliens)
		city.Aliens = nil
	}
}
