package graph

type Location struct {
	X, Y int
}

type Graph struct {
	nodeList [][]int
}

func New(x, y int) *Graph {
	graph := &Graph{}
	graph.nodeList = make([][]int, x)
	for i := range graph.nodeList {
		graph.nodeList[i] = make([]int, y)
	}
	return graph
}

func (g *Graph) Rows() [][]int {
	return g.nodeList
}

func (g *Graph) Column(r int) []int {
	return g.nodeList[r]
}

func GetPath(from, to *Location) interface{} {
	return nil
}
