package grid

import "github.com/paytonrules/pathfinding/graph"

type Room struct {
	Blocked bool
	grid    *Grid
	X, Y    int
}

func (r *Room) Adjacent() []graph.Node {
	return r.grid.adjacentTo(r)
}
