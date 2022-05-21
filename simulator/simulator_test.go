package simulator

import (
	"math/rand"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/luckychess/invasion/world"
	mock_world "github.com/luckychess/invasion/world/mock"
	"gotest.tools/assert"
)

func TestInit(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockWorld := mock_world.NewMockWorldMap(ctrl)
	testAliens := map[string]*world.Alien{
		"dude": {Name: "dude", City: "Somewhere"},
	}
	mockWorld.EXPECT().GetAliens().Return(testAliens).Times(2)
	simulator := InitSimulation(mockWorld, rand.New(rand.NewSource(0)), 123)
	assert.Assert(t, simulator.worldMap != nil)
	assert.Assert(t, simulator.aliensCount == 123)
	assert.Assert(t, simulator.worldMap.GetAliens()["dude"].Name == "dude")
	assert.Assert(t, simulator.worldMap.GetAliens()["dude"].City == "Somewhere")
}

func TestSimulateOneAlien(t *testing.T) {
	// one city, one alien
	ctrl := gomock.NewController(t)
	mockWorld := mock_world.NewMockWorldMap(ctrl)
	testCities := map[string]*world.City{
		"Dubai": {},
	}
	aliens := map[string]*world.Alien{
		"Honey": {Name: "Honey", City: "Dubai"},
	}

	mockWorld.EXPECT().GetCities().AnyTimes().Return(testCities)
	mockWorld.EXPECT().AddAlien(gomock.Any()).Times(1)
	mockWorld.EXPECT().MoveAlien(aliens["Honey"], gomock.Any()).Times(simulatorSteps)
	// (1 call + 1 call for every alien) * number of simulation steps
	mockWorld.EXPECT().GetAliens().Times((1 + 1) * simulatorSteps).Return(aliens)
	mockWorld.EXPECT().DestroyCity("Dubai").Times(int(1 + simulatorSteps))
	simulator := InitSimulation(mockWorld, rand.New(rand.NewSource(0)), 1)
	simulator.Simulate()
}

func TestSimulateTwoAliensOneCity(t *testing.T) {
	// 2 aliens, 1 city, city is instantly destroyed
	ctrl := gomock.NewController(t)
	mockWorld := mock_world.NewMockWorldMap(ctrl)
	testCities := map[string]*world.City{
		"Uglich": {},
	}
	mockWorld.EXPECT().GetCities().AnyTimes().Return(testCities)
	mockWorld.EXPECT().AddAlien(gomock.Any()).Times(2)
	mockWorld.EXPECT().MoveAlien(gomock.Any(), gomock.Any()).Times(0)
	destroyMock := mockWorld.EXPECT().DestroyCity("Uglich").Times(1)
	mockWorld.EXPECT().GetAliens().AnyTimes().After(destroyMock).Return(nil)
	simulator := InitSimulation(mockWorld, rand.New(rand.NewSource(0)), 2)
	simulator.Simulate()
}

func TestSimulateThreeAliensThreeCities(t *testing.T) {
	// 3 aliens, 3 cities, nothing happens
	ctrl := gomock.NewController(t)
	mockWorld := mock_world.NewMockWorldMap(ctrl)
	A := world.City{Name: "A"}
	B := world.City{Name: "B"}
	C := world.City{Name: "C"}
	A.East = &B
	B.West = &A
	B.East = &C
	C.West = &B
	testCities := map[string]*world.City{
		"A": &A,
		"B": &B,
		"C": &C,
	}
	aliens := map[string]*world.Alien{
		"DudeA": {Name: "DudeA", City: "A"},
		"DudeB": {Name: "DudeB", City: "B"},
		"DudeC": {Name: "DudeC", City: "C"},
	}

	mockWorld.EXPECT().GetCities().Times(3*3 + 1 + simulatorSteps).Return(testCities)
	mockWorld.EXPECT().AddAlien(gomock.Any()).Times(3)
	mockWorld.EXPECT().GetAliens().Times(2 * simulatorSteps).Return(aliens)
	mockWorld.EXPECT().MoveAlien(gomock.Any(), gomock.Any()).Times(3 * simulatorSteps)
	mockWorld.EXPECT().DestroyCity(gomock.Any()).Times(3 + 3*simulatorSteps)
	simulator := InitSimulation(mockWorld, rand.New(rand.NewSource(0)), 3)
	simulator.Simulate()
}

func TestStopSimulation(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockWorld := mock_world.NewMockWorldMap(ctrl)
	A := world.City{Name: "A"}
	B := world.City{Name: "B"}
	C := world.City{Name: "C"}
	A.East = &B
	B.West = &A
	B.East = &C
	C.West = &B
	testCities := map[string]*world.City{
		"A": &A,
		"B": &B,
		"C": &C,
	}
	mockWorld.EXPECT().GetCities().Times(1).Return(testCities)
	// StopSimulation doesn't require previous calls to StartSimulation
	simulator := InitSimulation(mockWorld, rand.New(rand.NewSource(0)), 1)
	result := simulator.StopSimulation()
	assert.Assert(t, strings.Contains(result, "A east=B"))
	assert.Assert(t, strings.Contains(result, "B east=C west=A"))
	assert.Assert(t, strings.Contains(result, "C west=B"))
}
