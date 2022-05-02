package mechanics

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
	c.newY = -10 * unit
}

//updates the Horizontal movement of the character
func (c *player) updateMovement() {
	c.x += c.newX
	c.y += c.newY
	if c.y > groundY*unit {
		c.y = groundY * unit
	}
	if c.newX > 0 {
		c.newX -= 4
	} else if c.newX < 0 {
		c.newX += 4
	}
	if c.newY < 20*unit {
		c.newY += 8
	}
}

func (g *Game) Update() error {
	if g.player == nil {
		g.player = &player{x: 50 * unit, y: groundY * unit}
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.player.newX = -4 * unit
		g.player.updateMovement()
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.player.newX = 4 * unit
		g.player.updateMovement()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) || inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) || inpututil.IsKeyJustPressed(ebiten.KeyW) {
		g.player.tryJump()
	}
	return nil
}
