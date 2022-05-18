package main

import (
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type player struct {
	x    int
	y    int
	newX int
	newY int
}

const (
	unit    = 16
	groundY = 420
)

//Makes character jump
//To Add: Kolla så man inte kan hoppa i all evighetet
func (c *player) tryJump() {
	c.newY = -14 * unit
}

func (c *player) tryDive() {
	c.newY = 10 * unit
}

//updates the Horizontal movement of the character
func (c *player) updateMovement() {
	c.x += c.newX
	c.y += c.newY
	if c.y > groundY*unit {
		c.y = groundY * unit
	}
	if c.newX > 0 {
		c.newX -= 6
	} else if c.newX < 0 {
		c.newX += 4
	}
	if c.newY < 20*unit {
		c.newY += 8
	}
}

//checks input for player movement
func (g *Game) updatePlayer() error {

	//Makes the player follow the ground if no input
	g.player.x -= g.cameraMovement * unit

	if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.player.newX = -6 * unit
	}
	if (ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyArrowRight)) && !g.hitBox() {
		g.player.newX = 6 * unit
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) || inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) || inpututil.IsKeyJustPressed(ebiten.KeyW) {
		g.player.tryJump()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) || inpututil.IsKeyJustPressed(ebiten.KeyS) {
		g.player.tryDive()
	}

	g.player.updateMovement()
	if g.hitWall() {
		g.mode = ModeGameOver
		g.gameoverCount = 30
	}
	return nil
}

//checks if Player has hit tha back wall
func (g *Game) hitWall() bool {
	if g.player.x <= 0 {
		return true
	}
	return false
}

//kolla det med tileY som return, kan va något där
func (g *Game) boxAt(tileX int) (tileY int, ok bool) {
	if (tileX - boxStartOffsetX) <= 0 {
		return 0, false
	}
	if floorMod(tileX-boxStartOffsetX, boxIntervalX) != 0 {
		return 0, false
	}
	idx := floorDiv(tileX-boxStartOffsetX, boxIntervalX)
	return g.boxTileXs[idx%len(g.boxTileXs)], true
}

//positions of player and boxes
var (
	boxX    = 0
	boxY    = 0
	playerX = 0
	playerY = 0
)

//checks if player hits box
func (g *Game) hitBox() bool {
	const (
		playerWidth  = 60
		playerHeight = 64
	)

	/* fmt.Println("camera: ", g.cameraX)
	fmt.Println("gubbe: ", playerX)
	fmt.Println("box: ", boxX) */

	for i := 0; i <= unit; i++ {
		if playerX+i+playerWidth == boxX && playerY >= (boxY-boxWidth) {
			println("Hit!")
			g.player.newX -= 6 * unit
			return true
			break
		}
	}
	return false
}
