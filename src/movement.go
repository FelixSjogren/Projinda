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
	if g.player == nil {
		g.player = &player{x: 100 * unit, y: (groundY + 4) * unit}
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.player.newX = -4 * unit
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.player.newX = 4 * unit
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) || inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) || inpututil.IsKeyJustPressed(ebiten.KeyW) {
		g.player.tryJump()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) || inpututil.IsKeyJustPressed(ebiten.KeyS) {
		g.player.tryDive()
	}

	g.player.updateMovement()
	if g.hit() {
		g.mode = ModeGameOver
		g.gameoverCount = 30
	}
	return nil
}

//checks if Player has hit tha back wall
func (g *Game) hit() bool {
	if g.player.x <= 0 {
		return true
	}
	return false
}
