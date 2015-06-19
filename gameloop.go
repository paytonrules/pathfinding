package main

import (
	"os"

	"github.com/paytonrules/pathfinding/graph"
	"github.com/veandco/go-sdl2/sdl"
)

var pixelSize int32 = 3
var roomSize int32 = 12

type gameObject struct {
	geometry []sdl.Rect
	location *graph.Location
}

var alien *gameObject = &gameObject{
	location: &graph.Location{X: 70, Y: 100},
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
	location: &graph.Location{X: 20, Y: 80},
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
		tGeometry[index].X += int32(g.location.X)
		tGeometry[index].Y += int32(g.location.Y)
	}
	for index := range tGeometry {
		tGeometry[index].X *= pixelSize
		tGeometry[index].Y *= pixelSize
		tGeometry[index].H *= pixelSize
		tGeometry[index].W *= pixelSize
	}
	r.FillRects(tGeometry)
}

func Draw(r *sdl.Renderer, g *graph.Graph) {
}

func DrawPaths(r *sdl.Renderer, p interface{}) {
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

	myGraph := graph.Create(100, 100)
	for running {
		for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
			}
		}

		path := graph.GetPath(gunner.location, alien.location)

		renderer.SetDrawColor(0, 0, 0, 255)
		renderer.Clear()

		renderer.SetDrawColor(255, 255, 255, 255)
		Draw(renderer, myGraph)
		DrawPaths(renderer, path)

		alien.Draw(renderer)
		gunner.Draw(renderer)
		renderer.Present()
	}
}
