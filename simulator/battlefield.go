package simulator

type Alien struct {
	id int
	name string
	position *City
}

type City struct {
	name string
	destroyed bool
	east *City
	north *City
	west *City
	south *City
}

type WorldMap struct {
	cities map[string]*City
}

func(m *WorldMap) init() {
	m.cities = make(map[string]*City)
}

func(m *WorldMap) addCity(name string, east string, north string, west string, south string) {
	city := City{name: name, destroyed: false}
	if east != "" {
		eastCity := m.cities[east]
		if eastCity != nil {
			city.east = eastCity
			eastCity.west = &city
		}
		m.cities[east] = eastCity
	}
	if north != "" {
		northCity := m.cities[north]
		if northCity != nil {
			city.north = northCity
			northCity.south = &city
		}
		m.cities[north] = northCity
	}
	if west != "" {
		westCity := m.cities[west]
		if westCity != nil {
			city.west = westCity
			westCity.east = &city
		}
		m.cities[west] = westCity
	}
	if south != "" {
		southCity := m.cities[south]
		if southCity != nil {
			city.south = southCity
			southCity.north = &city
		}
		m.cities[south] = southCity
	}

	m.cities[name] = &city
}
