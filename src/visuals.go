package main

import (
	"image"
	_ "image/png"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	playerImg *ebiten.Image
	groundImg *ebiten.Image
	boxImg    *ebiten.Image
)

const (
	tileSize        = 212
	boxWidth        = 100
	boxStartOffsetX = 8
	boxIntervalX    = 8
	boxGapY         = 5
)

//draws player
func (p *player) drawPL(screen *ebiten.Image) {
	s := playerImg

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(1, 1)
	op.GeoM.Translate(float64(p.x)/unit, float64(p.y)/unit)
	screen.DrawImage(s, op)

}

//draws the ground and box and moves camera
func (g *Game) drawGround(screen *ebiten.Image) {
	g.cameraX += 2
	const (
		newX        = windowWidth / tileSize
		newY        = windowHeight/tileSize + 0.5
		boxTileSrcX = 128
		boxTileSrcY = 192
	)

	op := &ebiten.DrawImageOptions{}

	for i := -2; i < newX+1; i++ {
		// ground
		op.GeoM.Reset()
		op.GeoM.Translate(float64(i*(tileSize-1)-floorMod(g.cameraX, tileSize/2)),
			float64((newY-1)*tileSize-floorMod(g.cameraY, tileSize)))
		screen.DrawImage(groundImg.SubImage(image.Rect(0, 0, tileSize, tileSize)).(*ebiten.Image), op)

		// box
		if tileY, ok := g.boxAt(floorDiv(g.cameraX, tileSize) + i); ok {

			for j := 0; j < tileY; j++ {
				op.GeoM.Reset()
				op.GeoM.Scale(1, -1)
				op.GeoM.Translate(float64(i*tileSize-floorMod(g.cameraX, tileSize)),
					//om man sätter tilesize i början till boxWidth blir det en rad
					float64((j+1)*tileSize-floorMod(g.cameraY, boxWidth)))
				//xy-pos typ men lite oklart om man sätter första argumentet i translate till o får man ett torn
				op.GeoM.Translate(float64(-(j)*boxWidth), float64(groundY-(boxWidth*(j+1))))
				var r image.Rectangle
				r = image.Rect(0, 0, boxWidth, boxWidth)
				screen.DrawImage(boxImg.SubImage(r).(*ebiten.Image), op)
			}
		}
	}
}

//Helper for box
func (g *Game) boxAt(tileX int) (tileY int, ok bool) {
	if (tileX - boxStartOffsetX) <= 0 {
		return 0, false
	}
	if floorMod(tileX-boxStartOffsetX, boxIntervalX) != 0 {
		return 0, false
	}
	idx := floorDiv(tileX-boxStartOffsetX, boxIntervalX)
	return g.boxTileYs[idx%len(g.boxTileYs)], true
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
	g.boxTileYs = make([]int, 256)
	for i := range g.boxTileYs {
		g.boxTileYs[i] = rand.Intn(6) + 2
	}
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
	boxImg, _, err = ebitenutil.NewImageFromFile("./images/box.png")
	if err != nil {
		log.Fatal(err)
	}
}
