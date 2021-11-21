package main

import (
	"AirWars2D"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

func main() {
	ebiten.SetWindowSize(640,480)
	ebiten.SetWindowTitle("AirWars2D")
	if err := ebiten.RunGame(&AirWars2D.Game{}); err != nil {
		log.Fatal(err)
	}
}
