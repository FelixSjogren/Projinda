package visuals

import (
	"image/color"
	_ "image/png"
	"log"

	"github.com/FelixSjogren/Projinda/src/mechanics/character/movement"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	playerImg *ebiten.Image
)

const (
	windowHeight = 640
	windowWidth  = 960
)

var (
	background = color.White
)

func (g *movement.Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return windowWidth, windowHeight
}

func (c *movement.player) draw(screen *ebiten.Image) {
	s := playerImg

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(1, 1)
	op.GeoM.Translate(float64(c.x)/movement.unit, float64(c.y)/movement.unit)
	screen.DrawImage(s, op)
}

func (g *movement.Game) Draw(screen *ebiten.Image) {
	screen.Fill(background)
	g.player.draw(screen)
}

func init() {
	var err error
	playerImg, _, err = ebitenutil.NewImageFromFile("./images/character_ball.png")
	if err != nil {
		log.Fatal(err)
	}
}
