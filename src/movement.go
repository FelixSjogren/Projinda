package main

import (
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

//A struct for the players movement and position
type player struct {
	x    int
	y    int
	newX int
	newY int
}

const (
	//one unit is 16 pixels, used in movement
	unit = 16
	//the Y-position of the ground
	groundY = 420
)

//Makes character jump
//Checks if the player is able to jump, checks if on ground or on a box
func (g *Game) tryJump() {
	if playerY == groundY+(playerHeight-20) || g.onBox() {

		g.player.newY = -14 * unit
	}
}

func (c *player) tryDive() {
	c.newY = 10 * unit
}

//updates movement of the player
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
	//if the player is in the air, it should go down (applies gravity)
	if g.player.newY < 20*unit && !g.onBox() {
		g.player.newY += 8
	}
	//failsafe, If the character somehow is inside the box, it is teleported ontop of it.
	if g.insideBox() {
		g.player.y = (boxY - (2 * boxWidth)) * unit
		g.player.newY = 0
	}
}

//checks input for player movement
func (g *Game) updatePlayer() error {

	//checks astroid hit.
	if g.hitAstroid() {
		g.mode = ModeGameOver
	}

	//Makes the player follow the ground if no input
	g.player.x -= g.cameraMovement * unit * 2

	//go left
	if (ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft)) && !g.hitBoxBack() {
		g.player.newX = -6 * unit
	}
	//go right
	if (ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyArrowRight)) && !g.hitBox() {
		g.player.newX = 12 * unit
	}

	//jump
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) || inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) || inpututil.IsKeyJustPressed(ebiten.KeyW) {
		g.tryJump()
	}

	//dive
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

//positions of player, boxes and astroids
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
	playerHeight = 68
)

//Checks if player hits astroids
func (g *Game) hitAstroid() bool {
	for i := 0; i <= boxWidth; i++ {
		if playerX+playerWidth == astroidX+i || playerX == astroidX+i ||
			playerX+playerWidth == astroidX+(boxWidth/2)+(i/2) || playerX == astroidX+(boxWidth/2)+(i/2) {
			for j := 0; j <= unit; j++ {
				if playerY-playerHeight == astroidY+j-20 || playerY == astroidY+j-20 ||
					playerY-playerHeight == astroidY+(boxWidth/2)+(j/2)-20 || playerY == astroidY+(boxWidth/2)+(j/2)-20 {

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
		if playerX+i+playerWidth == boxX && playerY >= (boxY-boxWidth) {
			g.player.newX = 0
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
					if inpututil.IsKeyJustPressed(ebiten.KeySpace) || inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) || inpututil.IsKeyJustPressed(ebiten.KeyW) {
						g.player.newY = -14 * unit
					} else {
						g.player.newY = 0
					}
					return true
					break
				}
			}
		}
	}

	return false
}

//Checks if player is inside a box
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
