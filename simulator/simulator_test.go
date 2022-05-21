package simulator

import (
	"math/rand"
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
	mockWorld.EXPECT().MoveAlien(aliens["Honey"], gomock.Any()).Times(10)
	// (1 call + 1 call for every alien) * number of simulation steps
	mockWorld.EXPECT().GetAliens().Times((1 + 1) * 10).Return(aliens)
	mockWorld.EXPECT().DestroyCity("Dubai").Times(int(1 + 10))
	simulator := InitSimulation(mockWorld, rand.New(rand.NewSource(0)), 1)
	simulator.Simulate()
}
