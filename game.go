package AirWars2D

import "github.com/hajimehoshi/ebiten/v2"

type Game struct {
	world *World
	ready bool
}

func (g* Game) Init() {
}


func (g* Game) Update() error {
	return nil
}

func (g* Game) Draw(screen *ebiten.Image) {
	if !g.ready {
		g.world = NewWorld()
		g.ready = true
	}
}

func (g* Game) Layout(outsidewidth,outsidheight int) (screenwidth,screenheight int) {
	return 640,480
}

