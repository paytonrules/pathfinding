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

func (g *Graph) EachAdjacencyList(cb func(n Node, l []Node)) {
	for key, value := range g.AdjacencyList {
		cb(key, value)
	}
}

func (g *Graph) Path(from, to Node) []Node {
	return nil
}
