package graph

type Point struct {
	X, Y int32
}

type GraphEdge struct {
	from *Node
	to   *Node
	cost int
}

type Node struct {
	*Point
}

func NewNode(p *Point) *Node {
	return &Node{p}
}

func NewGraphEdge(from, to *Node, cost int) *GraphEdge {
	return &GraphEdge{from: from, to: to, cost: cost}
}

func (g *GraphEdge) From() *Node {
	return g.from
}

func (g *GraphEdge) To() *Node {
	return g.to
}

func (g *GraphEdge) Cost() int {
	return g.cost
}
