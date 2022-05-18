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
	playerImg       *ebiten.Image
	playerRunForImg *ebiten.Image
	playerJumpImg   *ebiten.Image
	playerDeadImg   *ebiten.Image
	groundImg       *ebiten.Image
	boxImg          *ebiten.Image
	astroidImg      *ebiten.Image
	skyImg          *ebiten.Image
	fireImg         *ebiten.Image
)

const (
	tileSize        = 212
	boxWidth        = 100
	boxStartOffsetX = 8
	boxIntervalX    = 8
	boxGapY         = 0

	frameNum    = 4
	frameOX     = 0
	frameOY     = 0
	frameWidth  = 45
	frameHeight = 720

	playerFrameNum    = 6
	playerFrameOX     = 0
	playerFrameOY     = 0
	playerFrameWidth  = 60
	playerFrameHeight = 68
)

func (g *Game) drawDeadPL(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(1, 1)
	op.GeoM.Translate(float64(g.player.x)/unit, float64(g.player.y)/unit+(playerHeight-20))
	screen.DrawImage(playerDeadImg, op)
}

//draws player
func (g *Game) drawPL(screen *ebiten.Image) {

	op := &ebiten.DrawImageOptions{}

	if g.player.newY < 0 && g.player.newX < 0 {
		op.GeoM.Scale(-1, 1)
		op.GeoM.Translate(float64(g.player.x)/unit, float64(g.player.y)/unit+(playerHeight-20))
		screen.DrawImage(playerJumpImg, op)
		playerX = int(op.GeoM.Element(0, 2))
		playerY = int(op.GeoM.Element(1, 2))
		println(playerY)
	} else if g.player.newY < 0 {
		op.GeoM.Scale(1, 1)
		op.GeoM.Translate(float64(g.player.x)/unit, float64(g.player.y)/unit+(playerHeight-20))
		screen.DrawImage(playerJumpImg, op)
		playerX = int(op.GeoM.Element(0, 2))
		playerY = int(op.GeoM.Element(1, 2))
	} else if g.player.newX > 0 {
		op.GeoM.Scale(1, 1)
		op.GeoM.Translate(float64(g.player.x)/unit-23, float64(g.player.y)/unit+14)
		op.GeoM.Translate(playerFrameWidth/2, playerFrameHeight/2)
		i := (g.count / 5) % playerFrameNum
		sx, sy := playerFrameOX+i*playerFrameWidth, playerFrameOY
		screen.DrawImage(playerRunForImg.SubImage(image.Rect(sx, sy, sx+playerFrameWidth, sy+playerFrameHeight)).(*ebiten.Image), op)
		playerX = int(op.GeoM.Element(0, 2))
		playerY = int(op.GeoM.Element(1, 2))
	} else if g.player.newX < 0 {
		op.GeoM.Scale(-1, 1)
		op.GeoM.Translate(float64(g.player.x)/unit-23, float64(g.player.y)/unit+14)
		op.GeoM.Translate(playerFrameWidth/2, playerFrameHeight/2)
		i := (g.count / 5) % playerFrameNum
		sx, sy := playerFrameOX+i*playerFrameWidth, playerFrameOY
		screen.DrawImage(playerRunForImg.SubImage(image.Rect(sx, sy, sx+playerFrameWidth, sy+playerFrameHeight)).(*ebiten.Image), op)
		playerX = int(op.GeoM.Element(0, 2))
		playerY = int(op.GeoM.Element(1, 2))
	} else {
		op.GeoM.Scale(1, 1)
		op.GeoM.Translate(float64(g.player.x)/unit, float64(g.player.y)/unit+(playerHeight-20))
		screen.DrawImage(playerImg, op)
		playerX = int(op.GeoM.Element(0, 2))
		playerY = int(op.GeoM.Element(1, 2))
	}
}

// draws the sky
func (g *Game) drawSky(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Reset()
	screen.DrawImage(skyImg.SubImage(image.Rect(0, 0, windowWidth, windowHeight)).(*ebiten.Image), op)
}

//draws the fire
func (g *Game) drawFire(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Reset()
	op.GeoM.Translate(-float64(frameWidth), -float64(frameHeight))
	op.GeoM.Translate(frameWidth, frameHeight)
	i := (g.count / 5) % frameNum
	sx, sy := frameOX+i*frameWidth, frameOY
	screen.DrawImage(fireImg.SubImage(image.Rect(sx, sy, sx+frameWidth, sy+frameHeight)).(*ebiten.Image), op)

}

//draws the ground and box and moves camera
func (g *Game) drawGround(screen *ebiten.Image) {
	g.cameraX += g.cameraMovement
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
		op.GeoM.Translate(float64(i*(tileSize-1)-floorMod(g.cameraX, tileSize))+50,
			float64((newY-1)*tileSize-floorMod(g.cameraY, tileSize)))
		screen.DrawImage(groundImg.SubImage(image.Rect(0, 0, tileSize, tileSize)).(*ebiten.Image), op)

		// box
		if _, ok := g.boxPossibleAt(floorDiv(g.cameraX, tileSize) + i); ok {
			op.GeoM.Reset()
			op.GeoM.Scale(1, -1)
			op.GeoM.Translate(float64(i*tileSize-floorMod(g.cameraX, tileSize)),
				//om man sätter tilesize i början till boxWidth blir det en rad
				float64(tileSize-floorMod(g.cameraY, boxWidth)))
			//xy-pos typ men lite oklart om man sätter första argumentet i translate till o får man ett torn
			op.GeoM.Translate(float64(boxWidth), float64(groundY-(boxWidth)))
			var r image.Rectangle
			r = image.Rect(0, 0, boxWidth, boxWidth)
			screen.DrawImage(boxImg.SubImage(r).(*ebiten.Image), op)
			boxX = int(op.GeoM.Element(0, 2))
			boxY = int(op.GeoM.Element(1, 2))
		}

		//astroid
		if _, ok := g.boxPossibleAt(floorDiv(g.cameraX-500, tileSize) + i); ok {
			op.GeoM.Reset()
			op.GeoM.Scale(1, 1)
			op.GeoM.Translate(float64(i*tileSize-floorMod(g.cameraX-500, tileSize)),
				//om man sätter tilesize i början till boxWidth blir det en rad
				float64(tileSize-floorMod(g.cameraY, boxWidth)))
			//xy-pos typ men lite oklart om man sätter första argumentet i translate till o får man ett torn
			op.GeoM.Translate(float64(boxWidth), float64(groundY-(3*boxWidth)))
			var r image.Rectangle
			r = image.Rect(0, 0, boxWidth, boxWidth)
			screen.DrawImage(astroidImg.SubImage(r).(*ebiten.Image), op)
			astroidX = int(op.GeoM.Element(0, 2))
			astroidY = int(op.GeoM.Element(1, 2))
		}
	}
}

//Helper for box
func (g *Game) boxPossibleAt(tileX int) (tileY int, ok bool) {
	if (tileX - boxStartOffsetX) <= 0 {
		return 0, false
	}
	if floorMod(tileX-boxStartOffsetX, boxIntervalX) != 0 {
		return 0, false
	}
	idx := floorDiv(tileX-boxStartOffsetX, boxIntervalX)
	return g.boxTileXs[idx%len(g.boxTileXs)], true
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
	g.boxTileXs = make([]int, 256)
	for i := range g.boxTileYs {
		g.boxTileYs[i] = rand.Intn(6) + 2
	}
	for j := range g.boxTileXs {
		g.boxTileXs[j] = rand.Intn(6) + 2
	}
}

//gets needed images
func init() {
	var err error
	playerImg, _, err = ebitenutil.NewImageFromFile("./images/player.png")
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
	astroidImg, _, err = ebitenutil.NewImageFromFile("./images/astroid.png")
	if err != nil {
		log.Fatal(err)
	}
	skyImg, _, err = ebitenutil.NewImageFromFile("./images/sky.png")
	if err != nil {
		log.Fatal(err)
	}
	fireImg, _, err = ebitenutil.NewImageFromFile("./images/fire_anim.png")
	if err != nil {
		log.Fatal(err)
	}
	playerRunForImg, _, err = ebitenutil.NewImageFromFile("./images/playerRunFor.png")
	if err != nil {
		log.Fatal(err)
	}
	playerJumpImg, _, err = ebitenutil.NewImageFromFile("./images/playerJump.png")
	if err != nil {
		log.Fatal(err)
	}
	playerDeadImg, _, err = ebitenutil.NewImageFromFile("./images/playerDead.png")
	if err != nil {
		log.Fatal(err)
	}
}
