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
	cities []City
}