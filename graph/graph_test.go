package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGraphCreation(t *testing.T) {
	graph := Create(100, 90)

	assert.Equal(t, 100, len(graph.nodeList))
	for _, value := range graph.nodeList {
		assert.Equal(t, 90, len(value))
	}
}
