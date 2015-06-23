package grid

import "errors"

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
			grid.rooms[i][j] = &Room{X: i, Y: j}
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

func (g *Grid) RoomAt(x, y int) (*Room, error) {
	if x >= g.width || x < 0 {
		return nil, errors.New("X dimension invalid")
	} else if y >= g.height || y < 0 {
		return nil, errors.New("Y dimension invalid")
	}
	return g.rooms[x][y], nil
}

func GetPath(from, to *Room) interface{} {
	return nil
}
