package world

import "fmt"

type Alien struct {
	Name string
	City string
}

type City struct {
	name  string
	east  *City
	north *City
	west  *City
	south *City
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
	city := City{name: name}
	if east != "" {
		eastCity := m.Cities[east]
		if eastCity == nil {
			eastCity = &City{name: east, east: nil, north: nil, west: &city, south: nil}
		}
		city.east = eastCity
		eastCity.west = &city
		m.Cities[east] = eastCity
	}
	if north != "" {
		northCity := m.Cities[north]
		if northCity == nil {
			northCity = &City{name: north, east: nil, north: nil, west: nil, south: &city}
		}
		city.north = northCity
		northCity.south = &city
		m.Cities[north] = northCity
	}
	if west != "" {
		westCity := m.Cities[west]
		if westCity == nil {
			westCity = &City{name: west, east: &city, north: nil, west: nil, south: nil}
		}
		city.west = westCity
		westCity.east = &city
		m.Cities[west] = westCity
	}
	if south != "" {
		southCity := m.Cities[south]
		if southCity == nil {
			southCity = &City{name: south, east: nil, north: &city, west: nil, south: nil}
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
	return nil
}

func (m *WorldMap) DestroyCity(alienFirst *Alien, alienSecond *Alien) {
	if alienFirst.City == alienSecond.City {
		city := m.Cities[alienFirst.City]
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
		delete(m.Aliens, alienFirst.Name)
		delete(m.Aliens, alienSecond.Name)
	}
}
