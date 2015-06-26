package main

import (
	"os"

	"github.com/paytonrules/pathfinding/graph"
	"github.com/paytonrules/pathfinding/grid"
	"github.com/veandco/go-sdl2/sdl"
)

var pixelSize int32 = 3
var roomSize int32 = 12

type gameObject struct {
	geometry []sdl.Rect
	location *grid.Room
}

var alien *gameObject = &gameObject{
	geometry: []sdl.Rect{
		{2, 0, 1, 1},
		{8, 0, 1, 1},
		{3, 1, 1, 1},
		{7, 1, 1, 1},
		{2, 2, 7, 1},
		{1, 3, 2, 1},
		{4, 3, 3, 1},
		{8, 3, 2, 1},
		{0, 4, 11, 1},
		{0, 5, 1, 1},
		{2, 5, 7, 1},
		{10, 5, 1, 1},
		{0, 6, 1, 1},
		{2, 6, 1, 1},
		{8, 6, 1, 1},
		{10, 6, 1, 1},
		{3, 7, 2, 1},
		{6, 7, 2, 1},
	},
}

var gunner *gameObject = &gameObject{
	geometry: []sdl.Rect{
		{5, 0, 1, 1},
		{4, 1, 3, 2},
		{1, 3, 9, 1},
		{0, 4, 11, 3},
	},
}

func (g gameObject) Draw(r *sdl.Renderer) {
	tGeometry := make([]sdl.Rect, cap(g.geometry))
	copy(tGeometry, g.geometry)
	for index := range g.geometry {
		tGeometry[index].X += int32(g.location.X) * roomSize
		tGeometry[index].Y += int32(g.location.Y) * roomSize
	}
	for index := range tGeometry {
		tGeometry[index].X *= pixelSize
		tGeometry[index].Y *= pixelSize
		tGeometry[index].H *= pixelSize
		tGeometry[index].W *= pixelSize
	}
	r.FillRects(tGeometry)
}

func Draw(r *sdl.Renderer, g *grid.Grid) {
	g.EachRoom(func(room *grid.Room) {
		x := int32(room.X) * pixelSize * roomSize
		y := int32(room.Y) * pixelSize * roomSize
		width := roomSize * pixelSize
		height := roomSize * pixelSize
		rect := &sdl.Rect{x, y, width, height}

		if room.Blocked {
			r.FillRect(rect)
		} else {
			r.DrawRect(rect)
		}
	})
}

func DrawGraph(r *sdl.Renderer, g *graph.Graph) {
	nodesDrawn := make(map[graph.Node]struct{})
	g.EachAdjacencyList(func(n graph.Node, l []graph.Node) {
		for _, d := range l {
			_, drawn := nodesDrawn[d]

			if !drawn {
				startRoom := n.(*grid.Room)
				endRoom := d.(*grid.Room)
				x1 := int32(startRoom.X) * pixelSize * roomSize
				y1 := int32(startRoom.Y) * pixelSize * roomSize
				x2 := int32(endRoom.X) * roomSize * pixelSize
				y2 := int32(endRoom.Y) * roomSize * pixelSize

				r.DrawLine(int(x1), int(y1), int(x2), int(y2))
			}
		}

		nodesDrawn[n] = struct{}{}
	})
}

func DrawPath(r *sdl.Renderer, nodes []graph.Node) {
	for index, node := range nodes {
		if index != 0 {
			previousNode := nodes[index-1]
			startRoom := previousNode.(*grid.Room)
			endRoom := node.(*grid.Room)
			x1 := int32(startRoom.X) * pixelSize * roomSize
			y1 := int32(startRoom.Y) * pixelSize * roomSize
			x2 := int32(endRoom.X) * roomSize * pixelSize
			y2 := int32(endRoom.Y) * roomSize * pixelSize

			r.DrawLine(int(x1), int(y1), int(x2), int(y2))
		}
	}
}

func setupMap(g *grid.Grid) {
	var blockedRooms = [...][2]int{
		{1, 3},
		{3, 0},
		{3, 1},
		{3, 2},
		{3, 3},
		{3, 4},
		{3, 7},
		{3, 9},
		{3, 10},
		{3, 11},
		{4, 7},
		{4, 9},
		{4, 10},
		{4, 11},
		{5, 7},
		{5, 9},
		{5, 10},
		{5, 11},
		{6, 3},
		{6, 4},
		{6, 5},
		{7, 3},
		{7, 4},
		{7, 5},
		{7, 6},
		{7, 7},
		{7, 8},
		{8, 3},
		{8, 4},
		{8, 5},
		{8, 6},
		{8, 7},
		{8, 8},
		{10, 6},
		{10, 7},
		{10, 8},
		{10, 9},
		{11, 6},
		{11, 7},
		{11, 8},
		{11, 9},
		{11, 11},
		{12, 1},
		{12, 2},
		{12, 3},
		{12, 6},
		{12, 7},
		{12, 8},
		{12, 9},
		{12, 11},
		{13, 1},
		{13, 2},
		{13, 3},
		{14, 10},
		{14, 11},
		{15, 10},
		{15, 11},
		{16, 10},
		{16, 11},
	}

	for _, val := range blockedRooms {
		room, _ := g.RoomAt(val[0], val[1])
		room.Blocked = true
	}
}

func main() {
	window, err := sdl.CreateWindow("Monster Path",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		640, 480, sdl.WINDOW_SHOWN)

	if err != nil {
		os.Exit(1)
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		os.Exit(2)
	}
	defer renderer.Destroy()

	var event sdl.Event
	running := true

	myGrid := grid.New(17, 12)
	setupMap(myGrid)
	gunner.location, _ = myGrid.RoomAt(15, 7)
	alien.location, _ = myGrid.RoomAt(1, 1)
	graph := graph.NewFromGrid(myGrid)
	path := graph.Path(gunner.location, alien.location)
	for running {
		for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
			}
		}

		renderer.SetDrawColor(205, 201, 201, 255)
		renderer.Clear()

		renderer.SetDrawColor(10, 10, 10, 255)
		Draw(renderer, myGrid)
		DrawGraph(renderer, graph)
		renderer.SetDrawColor(200, 0, 0, 255)
		DrawPath(renderer, path)

		renderer.SetDrawColor(0, 100, 0, 255)
		alien.Draw(renderer)
		gunner.Draw(renderer)
		renderer.Present()
	}
}
