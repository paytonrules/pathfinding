package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNodeHasLocation(t *testing.T) {
	node := NewNode(&Point{X: 20, Y: 30})

	assert.Equal(t, 20, node.X)
	assert.Equal(t, 30, node.Y)
}

func TestGraphEdgeHasTwoNodesAndACost(t *testing.T) {
	nodeOne := NewNode(&Point{})
	nodeTwo := NewNode(&Point{})
	edge := NewGraphEdge(nodeOne, nodeTwo, 20)

	assert.Equal(t, nodeOne, edge.From())
	assert.Equal(t, nodeTwo, edge.To())
	assert.Equal(t, 20, edge.Cost())
}
