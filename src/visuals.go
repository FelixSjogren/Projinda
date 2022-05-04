package main

import (
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	playerImg *ebiten.Image
	groundImg *ebiten.Image
)

const (
	tileSize = 212
)

//draws player
func (p *player) drawPL(screen *ebiten.Image) {
	s := playerImg

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(1, 1)
	op.GeoM.Translate(float64(p.x)/unit, float64(p.y)/unit)
	screen.DrawImage(s, op)

}

//draws the ground and moves camera
func (g *Game) drawGround(screen *ebiten.Image) {
	g.cameraX += 2
	const (
		newX = windowWidth / tileSize
		newY = windowHeight/tileSize + 0.5
	)

	op := &ebiten.DrawImageOptions{}

	for i := -2; i < newX+1; i++ {
		// ground
		op.GeoM.Reset()
		op.GeoM.Translate(float64(i*(tileSize-1)-floorMod(g.cameraX, tileSize)),
			float64((newY-1)*tileSize-floorMod(g.cameraY, tileSize)))
		screen.DrawImage(groundImg.SubImage(image.Rect(0, 0, tileSize, tileSize)).(*ebiten.Image), op)
	}
}

//Helper function for ground placement
func floorDiv(x, y int) int {
	d := x / y
	if d*y == x || x >= 0 {
		return d
	}
	return d - 1
}

//Also helper for ground placement
func floorMod(x, y int) int {
	return x - floorDiv(x, y)*y
}

//sets camera position
func (g *Game) init() {
	g.cameraX = -240
	g.cameraY = 0
}

//gets needed images
func init() {
	var err error
	playerImg, _, err = ebitenutil.NewImageFromFile("./images/character_ball.png")
	if err != nil {
		log.Fatal(err)
	}
	groundImg, _, err = ebitenutil.NewImageFromFile("./images/ground.png")
	if err != nil {
		log.Fatal(err)
	}
}
