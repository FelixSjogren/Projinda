package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	windowHeight = 640
	windowWidth  = 960
)

var (
	background = color.White
)

func main() {
	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle("Wall Game")
	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return windowWidth, windowHeight
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(background)
	g.player.drawPL(screen)
}
