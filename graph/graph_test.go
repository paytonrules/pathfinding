package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGraphCreation(t *testing.T) {
	graph := New(100, 90)

	assert.Equal(t, 100, len(graph.nodeList))
	for _, value := range graph.nodeList {
		assert.Equal(t, 90, len(value))
	}
}

func TestEachNode(t *testing.T) {
	graph := New(2, 2)
	nodes := make([]*Node, 0, 4)

	graph.EachNode(func(n *Node) {
		nodes = append(nodes, n)
	})

	assert.Equal(t, 4, len(nodes))
}

func TestColumn(t *testing.T) {
	graph := New(2, 3)

	assert.Equal(t, 3, len(graph.Column(0)))
}
