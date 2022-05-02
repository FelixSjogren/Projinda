package main

import (
	_ "image/png"

	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	charImg *ebiten.Image
)

func init() {
	var err error
	charImg, _, err = ebitenutil.NewImageFromFile("./images/character_ball.png")
	if err != nil {
		log.Fatal(err)
	}
}
