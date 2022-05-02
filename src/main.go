package main

import (
	"github.com/hajimehoshi/ebiten"
)

type Game struct {
	player movement.player
}

func main() {
	ebiten.SetWindowSize(visuals.windowWidth, visuals.windowHeight)
	ebiten.SetWindowTitle("Wall Game")
	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}
