package grid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoomStartsAsUnblocked(t *testing.T) {
	room := &Room{}

	assert.False(t, room.Blocked)
	assert.Equal(t, 0, room.X)
	assert.Equal(t, 0, room.Y)
}
