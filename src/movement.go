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
func (g *Game) tryJump() {
	if playerY == groundY+(playerHeight-10) || g.onBox() {
		g.player.newY = -14 * unit
	}
}

func (c *player) tryDive() {
	c.newY = 10 * unit
}

//updates the Horizontal movement of the character
func (g *Game) updateMovement() {
	g.player.x += g.player.newX
	g.player.y += g.player.newY
	if g.player.y > groundY*unit {
		g.player.y = groundY * unit
	}
	if g.player.newX > 0 {
		g.player.newX -= 6
	} else if g.player.newX < 0 {
		g.player.newX += 4
	}
	if g.player.newY < 20*unit && !g.onBox() {
		g.player.newY += 8
	}

	if g.insideBox() {
		g.player.y = (boxY - (2 * boxWidth)) * unit
		g.player.newY = 0
	}
}

//checks input for player movement
func (g *Game) updatePlayer() error {

	if g.hitAstroid() {
		g.mode = ModeGameOver
	}

	//Makes the player follow the ground if no input
	g.player.x -= g.cameraMovement * unit

	if (ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft)) && !g.hitBoxBack() {
		g.player.newX = -6 * unit
	}
	if (ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyArrowRight)) && !g.hitBox() {
		g.player.newX = 12 * unit
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) || inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) || inpututil.IsKeyJustPressed(ebiten.KeyW) {
		g.tryJump()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) || inpututil.IsKeyJustPressed(ebiten.KeyS) {
		g.player.tryDive()
	}

	g.updateMovement()
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
	boxX     = 0
	boxY     = 0
	playerX  = 0
	playerY  = 0
	astroidX = -200
	astroidY = -200
)

const (
	playerWidth  = 60
	playerHeight = 64
)

//Checks if player hits astroids
func (g *Game) hitAstroid() bool {
	for i := 0; i <= boxWidth; i++ {
		if playerX+playerWidth == astroidX+i || playerX == astroidX+i ||
			playerX+playerWidth == astroidX+(boxWidth/2)+(i/2) || playerX == astroidX+(boxWidth/2)+(i/2) {
			for j := 0; j <= unit; j++ {
				if playerY+playerHeight == astroidY+j || playerY == astroidY+j ||
					playerY+playerHeight == astroidY+(boxWidth/2)+(j/2) || playerX == astroidX+(boxWidth/2)+(j/2) {
					return true
					break
				}
			}
		}
	}
	return false
}

//checks if player hits box
func (g *Game) hitBox() bool {

	for i := 0; i <= unit; i++ {
		if playerX+i+playerWidth-10 == boxX && playerY >= (boxY-boxWidth) {
			g.player.newX = 0
			//g.player.x += int(math.Abs(float64(playerX - boxX)))
			return true
			break
		}
	}

	return false
}

//checks if player hits box with back
func (g *Game) hitBoxBack() bool {

	for i := 0; i <= unit; i++ {
		if playerX-i == boxX+boxWidth && playerY >= (boxY-boxWidth) {
			g.player.newX = 0
			//g.player.x += int(math.Abs(float64(playerX - boxX)))
			return true
			break
		}
	}

	return false
}

//Checks if player is on top of a box
func (g *Game) onBox() bool {

	for i := 0; i <= boxWidth; i++ {
		if playerX+playerWidth == boxX+i || playerX == boxX+i ||
			playerX+playerWidth == boxX+(boxWidth/2)+(i/2) || playerX == boxX+(boxWidth/2)+(i/2) {
			for j := 0; j <= unit; j++ {
				if playerY+j+playerHeight == boxY-boxWidth {
					println("on box")
					if inpututil.IsKeyJustPressed(ebiten.KeySpace) || inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) || inpututil.IsKeyJustPressed(ebiten.KeyW) {
						g.player.newY = -14 * unit
					} else {
						g.player.newY = 0
					}
					//g.player.y = boxY - boxWidth
					return true
					break
				}
			}
		}
	}

	return false
}

func (g *Game) insideBox() bool {
	for i := boxX; i < (boxX + boxWidth); i++ {
		if playerX == i {
			for j := boxY; j > (boxY - boxWidth); j-- {
				if playerY == j {
					return true
				}
			}
		}
	}
	return false
}
