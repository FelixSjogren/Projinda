package main

import (
	"image/color"
	_ "image/png"

	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	windowHeight = 640
	windowWidth  = 960
)

const (
	unit    = 16
	groundY = 420
)

var (
	background = color.White
)

var (
	charImg *ebiten.Image
)

func init() {
	var err error
	charImg, _, err = ebitenutil.NewImageFromFile("./Images/character_ball.png")
	if err != nil {
		log.Fatal(err)
	}
}

type char struct {
	x    int
	y    int
	newX int
	newY int
}

func (c *char) updateMovement() {
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

func (c *char) tryJump() {
	c.newY = -10 * unit
}

type Game struct {
	player *char
}

func (g *Game) Update() error {
	if g.player == nil {
		g.player = &char{x: 50 * unit, y: groundY * unit}
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.player.newX = -4 * unit
	} else if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.player.newX = 4 * unit
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) || inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) || inpututil.IsKeyJustPressed(ebiten.KeyW) {
		g.player.tryJump()
	}

	g.player.updateMovement()
	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return windowWidth, windowHeight
}

func (c *char) draw(screen *ebiten.Image) {
	s := charImg

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(1, 1)
	op.GeoM.Translate(float64(c.x)/unit, float64(c.y)/unit)
	screen.DrawImage(s, op)
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(background)
	g.player.draw(screen)
}

func main() {
	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle("Wall Game")
	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}
