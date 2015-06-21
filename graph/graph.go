package graph

type Location struct {
	X, Y int
}

type Graph struct {
	nodeList [][]*Node
}

type Node struct {
}

func New(x, y int) *Graph {
	graph := &Graph{}
	graph.nodeList = make([][]*Node, x)
	for i := range graph.nodeList {
		graph.nodeList[i] = make([]*Node, y)
	}
	return graph
}

func (g *Graph) Rows() [][]*Node {
	return g.nodeList
}

func (g *Graph) Column(r int) []*Node {
	return g.nodeList[r]
}

func (g *Graph) EachNode(f func(*Node)) {
	for _, r := range g.nodeList {
		for _, n := range r {
			f(n)
		}
	}
}

func GetPath(from, to *Location) interface{} {
	return nil
}
