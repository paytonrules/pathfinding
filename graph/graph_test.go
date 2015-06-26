package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type simpleNode struct {
	adjacentNodes []Node
}

func (n *simpleNode) Adjacent() []Node {
	return n.adjacentNodes
}

type simpleGrid struct {
	nodeList []Node
}

func (g *simpleGrid) Nodes() []Node {
	return g.nodeList
}

func TestGraphWithNoNodes(t *testing.T) {
	grid := &simpleGrid{}
	graph := NewFromGrid(grid)

	assert.Equal(t, 0, len(graph.AdjacencyList))
}

func TestGraphWithNoAdjacentNodes(t *testing.T) {
	node := &simpleNode{}
	grid := &simpleGrid{nodeList: []Node{node}}

	graph := NewFromGrid(grid)

	assert.Equal(t, 1, len(graph.AdjacencyList))
	assert.Equal(t, 0, len(graph.AdjacencyList[node]))
}

func TestGraphWithOneAdjacentNode(t *testing.T) {
	adjacentNode := &simpleNode{}
	node := &simpleNode{adjacentNodes: []Node{adjacentNode}}
	grid := &simpleGrid{nodeList: []Node{node}}

	graph := NewFromGrid(grid)

	assert.Equal(t, 1, len(graph.AdjacencyList))
	assert.Equal(t, []Node{adjacentNode}, graph.AdjacencyList[node])
}

func TestGraphCanEnumerateItsAdjacencyLists(t *testing.T) {
	adjacentNode := &simpleNode{}
	node := &simpleNode{adjacentNodes: []Node{adjacentNode}}
	grid := &simpleGrid{nodeList: []Node{node}}

	graph := NewFromGrid(grid)

	nodes := map[Node][]Node{}
	graph.EachAdjacencyList(func(n Node, l []Node) {
		nodes[n] = l
	})

	expected := map[Node][]Node{
		node: []Node{adjacentNode},
	}
	assert.Equal(t, expected, nodes)
}
