package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

type Mode int

const (
	ModeTitle Mode = iota
	ModeGame
	ModeGameOver
)

const (
	windowHeight  = 720
	windowWidth   = 1280
	tileSize      = 212
	titleFontSize = fontSize * 1.5
	fontSize      = 24
	smallFontSize = fontSize / 2
)

var (
	titleArcadeFont font.Face
	arcadeFont      font.Face
	smallArcadeFont font.Face
)

type Game struct {
	mode Mode

	//Player
	player *player

	// Camera
	cameraX int
	cameraY int

	gameoverCount int
}

var (
	background = color.White
)

func main() {
	newGame := &Game{}
	newGame.mode = ModeGame
	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle("Wall Game")
	if err := ebiten.RunGame(newGame); err != nil {
		panic(err)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return windowWidth, windowHeight
}

func (g *Game) Update() error {
	switch g.mode {
	case ModeTitle:
		if g.isKeyPressed() {
			g.mode = ModeGame
		}
	case ModeGame:
		g.updatePlayer()
	case ModeGameOver:
		if g.gameoverCount > 0 {
			g.gameoverCount--
		} else if g.gameoverCount == 0 && g.isKeyPressed() {
			g.init()
			g.mode = ModeGame
			g.player.x = 100 * unit
			g.player.y = (groundY + 4) * unit
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(background)
	var titleTexts []string
	var texts []string
	switch g.mode {
	case ModeTitle:
		titleTexts = []string{"WALL GAME"}
		texts = []string{"", "", "", "", "", "", "", "Press space to play", ""}
	case ModeGameOver:
		texts = []string{"", "GAME OVER!"}
		texts = []string{"", "", "", "", "", "", "", "Press space to play again", ""}
	}
	for i, l := range titleTexts {
		x := (windowWidth - len(l)*titleFontSize) / 2
		text.Draw(screen, l, titleArcadeFont, x, (i+4)*titleFontSize, color.Black)
	}
	for i, l := range texts {
		x := (windowWidth - len(l)*fontSize) / 2
		text.Draw(screen, l, arcadeFont, x, (i+4)*fontSize, color.Black)
	}

	g.player.drawPL(screen)
	g.drawGround(screen)
}

func (g *Game) isKeyPressed() bool {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		return true
	}
	return false
}
