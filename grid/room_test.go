package grid

import (
	"testing"

	"github.com/paytonrules/pathfinding/graph"
	"github.com/stretchr/testify/assert"
)

func TestRoomStartsAsUnblocked(t *testing.T) {
	room := &Room{}

	assert.False(t, room.Blocked)
	assert.Equal(t, 0, room.X)
	assert.Equal(t, 0, room.Y)
}

func TestRoomGetsAdjacentNodesFromGrid(t *testing.T) {
	roomOne := &Room{X: 0, Y: 0}
	roomTwo := &Room{X: 0, Y: 1}
	grid := New(1, 2)
	grid.setRoom(0, 0, roomOne)
	grid.setRoom(0, 1, roomTwo)
	roomOne.grid = grid

	assert.Equal(t, []graph.Node{roomTwo}, roomOne.Adjacent())
}
