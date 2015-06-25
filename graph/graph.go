package graph

type Node interface {
	Adjacent() []Node
}

type Graph struct {
	AdjacencyList map[Node][]Node
}

type Grid interface {
	Nodes() []Node
}

func NewFromGrid(g Grid) *Graph {
	graph := &Graph{}
	graph.AdjacencyList = make(map[Node][]Node)

	for _, node := range g.Nodes() {
		graph.AdjacencyList[node] = node.Adjacent()
	}
	return graph
}
