package simulator

import "fmt"

type Alien struct {
	name string
	city string
}

type City struct {
	name      string
	destroyed bool
	east      *City
	north     *City
	west      *City
	south     *City
}

type WorldMap struct {
	cities map[string]*City
	aliens map[string]*Alien
}

func initWorldMap() WorldMap {
	worldMap := WorldMap{}
	worldMap.cities = make(map[string]*City)
	worldMap.aliens = make(map[string]*Alien)
	return worldMap
}

func (m *WorldMap) addCity(name string, east string, north string, west string, south string) {
	city := City{name: name, destroyed: false}
	if east != "" {
		eastCity := m.cities[east]
		if eastCity == nil {
			eastCity = &City{name: east, east: nil, north: nil, west: &city, south: nil}
		}
		city.east = eastCity
		eastCity.west = &city
		m.cities[east] = eastCity
	}
	if north != "" {
		northCity := m.cities[north]
		if northCity == nil {
			northCity = &City{name: north, east: nil, north: nil, west: nil, south: &city}
		}
		city.north = northCity
		northCity.south = &city
		m.cities[north] = northCity
	}
	if west != "" {
		westCity := m.cities[west]
		if westCity == nil {
			westCity = &City{name: west, east: &city, north: nil, west: nil, south: nil}
		}
		city.west = westCity
		westCity.east = &city
		m.cities[west] = westCity
	}
	if south != "" {
		southCity := m.cities[south]
		if southCity == nil {
			southCity = &City{name: south, east: nil, north: &city, west: nil, south: nil}
		}
		city.south = southCity
		southCity.north = &city
		m.cities[south] = southCity
	}

	m.cities[name] = &city
}

func (m *WorldMap) addAlien(alien *Alien) {
	if m.cities[alien.city] == nil {
		panic(fmt.Errorf("trying to unleash an alien %s into non-existing city %s", alien.name, alien.city))
	}
	m.aliens[alien.name] = alien
}

func (m *WorldMap) destroyCity(alienFirst *Alien, alienSecond *Alien) {
	if alienFirst.city == alienSecond.city {
		city := m.cities[alienFirst.city]
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
		delete(m.cities, city.name)
		delete(m.aliens, alienFirst.name)
		delete(m.aliens, alienSecond.name)
	}
}
