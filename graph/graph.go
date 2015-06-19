package graph

type Location struct {
	X, Y int
}

type Graph struct {
	nodeList [][]int
}

func Create(x, y int) *Graph {
	graph := &Graph{}
	graph.nodeList = make([][]int, 100)
	for i := range graph.nodeList {
		graph.nodeList[i] = make([]int, 90)
	}
	return graph
}

func GetPath(from, to *Location) interface{} {
	return nil
}
