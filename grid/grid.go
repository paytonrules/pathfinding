package grid

import (
	"errors"

	"github.com/paytonrules/pathfinding/graph"
)

type Grid struct {
	rooms         [][]*Room
	width, height int
}

func New(x, y int) *Grid {
	grid := &Grid{}
	grid.width, grid.height = x, y
	grid.rooms = make([][]*Room, x)
	for i := range grid.rooms {
		grid.rooms[i] = make([]*Room, y)

		for j := range grid.rooms[i] {
			grid.rooms[i][j] = &Room{X: i, Y: j, grid: grid}
		}
	}
	return grid
}

func (g *Grid) EachRoom(f func(*Room)) {
	for _, r := range g.rooms {
		for _, n := range r {
			f(n)
		}
	}
}

func (g *Grid) Nodes() []graph.Node {
	nodes := make([]graph.Node, 0)
	g.EachRoom(func(r *Room) {
		nodes = append(nodes, r)
	})
	return nodes
}

func (g *Grid) isValidRoom(x, y int) error {
	if x >= g.width || x < 0 {
		return errors.New("X dimension invalid")
	} else if y >= g.height || y < 0 {
		return errors.New("Y dimension invalid")
	}
	return nil
}

func (g *Grid) RoomAt(x, y int) (*Room, error) {
	err := g.isValidRoom(x, y)
	if err != nil {
		return nil, err
	}

	return g.rooms[x][y], nil
}

func (g *Grid) setRoom(x, y int, r *Room) error {
	err := g.isValidRoom(x, y)

	if err == nil {
		g.rooms[x][y] = r
	}
	return err
}

func (g *Grid) adjacentTo(r *Room) []graph.Node {
	adjacentRooms := make([]graph.Node, 0, 8)
	if r.Blocked {
		return adjacentRooms
	}

	for x := r.X - 1; x <= r.X+1; x++ {
		for y := r.Y - 1; y <= r.Y+1; y++ {
			if x != r.X || y != r.Y {
				room, err := g.RoomAt(x, y)
				if err == nil && room.Blocked == false {
					adjacentRooms = append(adjacentRooms, room)
				}
			}
		}
	}

	return adjacentRooms
}
