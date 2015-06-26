package grid

import (
	"testing"

	"github.com/paytonrules/pathfinding/graph"
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

func TestReturnsEachRoomForNodes(t *testing.T) {
	grid := New(2, 2)
	nodes := grid.Nodes()

	assert.Equal(t, 4, len(nodes))
}

func TestOneRoomHasNothingAdjacentToIt(t *testing.T) {
	grid := New(1, 1)
	room, _ := grid.RoomAt(0, 0)

	assert.Equal(t, 0, len(grid.adjacentTo(room)))
}

func TestTwoVerticalRoomsAreAdjacentToEachOther(t *testing.T) {
	grid := New(1, 2)
	roomOne, _ := grid.RoomAt(0, 0)
	roomTwo, _ := grid.RoomAt(0, 1)

	assert.Equal(t, []graph.Node{roomTwo}, grid.adjacentTo(roomOne))
	assert.Equal(t, []graph.Node{roomOne}, grid.adjacentTo(roomTwo))
}

func TestTwoHorizontalRoomsAreAdjacentToEachOther(t *testing.T) {
	grid := New(2, 1)
	roomOne, _ := grid.RoomAt(0, 0)
	roomTwo, _ := grid.RoomAt(1, 0)

	assert.Equal(t, []graph.Node{roomTwo}, grid.adjacentTo(roomOne))
	assert.Equal(t, []graph.Node{roomOne}, grid.adjacentTo(roomTwo))
}

func TestDiagonallyAdjacentRoomsAreIncluded(t *testing.T) {
	grid := New(2, 2)
	room, _ := grid.RoomAt(0, 0)
	diagRoom, _ := grid.RoomAt(1, 1)

	assert.Contains(t, grid.adjacentTo(room), diagRoom)
}

func TestAboveAndBelowAreIncluded(t *testing.T) {
	grid := New(1, 3)
	room, _ := grid.RoomAt(0, 1)
	above, _ := grid.RoomAt(0, 0)
	below, _ := grid.RoomAt(0, 0)

	assert.Contains(t, grid.adjacentTo(room), above, below)
}

func TestRightAndLeftAreIncluded(t *testing.T) {
	grid := New(3, 1)
	room, _ := grid.RoomAt(1, 0)
	left, _ := grid.RoomAt(0, 0)
	right, _ := grid.RoomAt(2, 0)

	assert.Contains(t, grid.adjacentTo(room), left, right)
}

func TestBlockedRoomsAreNotAdjacent(t *testing.T) {
	grid := New(1, 2)
	room, _ := grid.RoomAt(0, 0)
	blockedRoom, _ := grid.RoomAt(0, 1)
	blockedRoom.Blocked = true

	assert.Equal(t, 0, len(grid.adjacentTo(room)))
}

func TestBlockedRoomsHaveNoAdjacentRooms(t *testing.T) {
	grid := New(1, 2)
	blockedRoom, _ := grid.RoomAt(0, 1)
	blockedRoom.Blocked = true

	assert.Equal(t, 0, len(grid.adjacentTo(blockedRoom)))
}
