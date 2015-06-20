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

func TestRows(t *testing.T) {
	graph := New(2, 90)

	assert.Equal(t, 2, len(graph.Rows()))
	assert.Equal(t, 90, len(graph.Rows()[0]))
	assert.Equal(t, 90, len(graph.Rows()[1]))
}

func TestColumn(t *testing.T) {
	graph := New(2, 3)

	assert.Equal(t, 3, len(graph.Column(0)))
}
