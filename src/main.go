package main

import (
	"github.com/FelixSjogren/Projinda/src/character/mechanics"
)



const (
	windowHeight = 640
	windowWidth  = 960
)

var (
	background = color.White
)

type Game struct {
	player mechanics.*player
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return windowWidth, windowHeight
}

func (c *char) draw(screen *ebiten.Image) {
	s := charImg

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(1, 1)
	op.GeoM.Translate(float64(c.x)/unit, float64(c.y)/unit)
	screen.DrawImage(s, op)
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(background)
	g.player.draw(screen)
}

func main() {
	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle("Wall Game")
	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}