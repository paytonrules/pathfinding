package grid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGridCreation(t *testing.T) {
	grid := New(100, 90)

	assert.Equal(t, 100, len(grid.rooms))
	for _, value := range grid.rooms {
		assert.Equal(t, 90, len(value))
	}
}

func TestEachRoom(t *testing.T) {
	grid := New(2, 2)
	rooms := make([]*Room, 0, 4)

	grid.EachRoom(func(n *Room) {
		rooms = append(rooms, n)
	})

	assert.Equal(t, 4, len(rooms))
}

func TestEachRoomHasItsLocation(t *testing.T) {
	grid := New(2, 2)

	room, _ := grid.RoomAt(0, 0)
	assert.Equal(t, 0, room.X)
	assert.Equal(t, 0, room.Y)

	room, _ = grid.RoomAt(1, 1)
	assert.Equal(t, 1, room.X)
	assert.Equal(t, 1, room.Y)
}

func TestRoomAt(t *testing.T) {
	grid := New(3, 3)

	node, _ := grid.RoomAt(1, 1)
	assert.NotNil(t, node)
}

func TestInvalidLocationInX(t *testing.T) {
	grid := New(3, 3)

	node, err := grid.RoomAt(3, 2)
	assert.Nil(t, node)
	assert.NotNil(t, err)

	node, err = grid.RoomAt(-1, 2)
	assert.Nil(t, node)
	assert.NotNil(t, err)
}

func TestInvalidLocationInY(t *testing.T) {
	grid := New(3, 3)

	node, err := grid.RoomAt(2, 3)
	assert.Nil(t, node)
	assert.NotNil(t, err)

	node, err = grid.RoomAt(2, -1)
	assert.Nil(t, node)
	assert.NotNil(t, err)
}
