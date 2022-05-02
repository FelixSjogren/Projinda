package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	playerImg *ebiten.Image
)

func (c *player) drawPL(screen *ebiten.Image) {
	s := playerImg

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(1, 1)
	op.GeoM.Translate(float64(c.x)/unit, float64(c.y)/unit)
	screen.DrawImage(s, op)
}

func init() {
	var err error
	playerImg, _, err = ebitenutil.NewImageFromFile("./images/character_ball.png")
	if err != nil {
		log.Fatal(err)
	}
}
