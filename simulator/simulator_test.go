package simulator

import (
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
	simulator := InitSimulation(mockWorld, 123)
	assert.Assert(t, simulator.worldMap != nil)
	assert.Assert(t, simulator.aliensCount == 123)
	assert.Assert(t, simulator.worldMap.GetAliens()["dude"].Name == "dude")
	assert.Assert(t, simulator.worldMap.GetAliens()["dude"].City == "Somewhere")
}
