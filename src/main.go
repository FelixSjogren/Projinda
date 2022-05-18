package main

import (
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
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
	titleFontSize = fontSize * 1.5
	fontSize      = 24
	smallFontSize = fontSize / 2
)

var (
	titleArcadeFont font.Face
	arcadeFont      font.Face
	smallArcadeFont font.Face
)

//a game type, wich is needed to run the game
type Game struct {
	mode Mode

	//Player
	player *player

	// Camera
	cameraX        int
	cameraY        int
	cameraMovement int

	// Boxes
	boxTileYs []int
	boxTileXs []int

	//astroids
	astroidYpos int
	astroidXpos int

	gameoverCount int

	count int
}

func NewGame() *Game {
	g := &Game{}
	g.init()
	return g
}

func main() {
	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle("Wall Game")
	if err := ebiten.RunGame(NewGame()); err != nil {
		panic(err)
	}
}

//needed for ebiten.RunGame
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return windowWidth, windowHeight
}

//This is looped infinitely in ebiten.RunGame, checks which mode Game is in and preforms the correct actions
func (g *Game) Update() error {
	g.count++
	switch g.mode {
	case ModeTitle:
		if g.isSpacePressed() {
			if g.player == nil {
				g.player = &player{x: 200 * unit, y: (groundY - 40) * unit}
			}
			g.cameraMovement = 6
			g.mode = ModeGame
			//time.Sleep(time.Millisecond * 100)
		}
	case ModeGame:
		g.updatePlayer()
	case ModeGameOver:

		if g.gameoverCount > 0 {
			g.gameoverCount--
		}
		if g.gameoverCount == 0 && g.isSpacePressed() {
			g.init()
			g.mode = ModeTitle
			g.player.x = 200 * unit
			g.player.y = (groundY - 40) * unit
			boxX = 0
			boxY = 0
			playerX = 0
			playerY = 0
			astroidX = -200
			astroidY = -200
		}
	}
	return nil
}

//Also looped infinitely from ebiten.RunGame, Draws what is needen on the screen
func (g *Game) Draw(screen *ebiten.Image) {
	g.drawSky(screen)
	g.drawFire(screen)
	g.drawGround(screen)

	if g.mode != ModeTitle {
		g.drawPL(screen)
	}

	var titleTexts []string
	var texts []string
	switch g.mode {
	case ModeTitle:
		titleTexts = []string{"WALL GAME"}
		texts = []string{"", "", "", "", "", "", "", "Press space to play", ""}
	case ModeGameOver:
		titleTexts = []string{"", "GAME OVER!"}
		texts = []string{"", "", "", "", "", "", "", "Press space to return to title-screen", ""}
		g.drawDeadPL(screen)
	}
	for i, l := range titleTexts {
		x := (windowWidth - len(l)*titleFontSize) / 2
		text.Draw(screen, l, titleArcadeFont, x, (i+4)*titleFontSize, color.Black)
	}
	for i, l := range texts {
		x := (windowWidth - len(l)*fontSize) / 2
		text.Draw(screen, l, arcadeFont, x, (i+4)*fontSize, color.Black)
	}

	if g.mode == ModeTitle {
		msg := []string{
			"Incredible game by",
			"Felix & John",
		}
		for i, l := range msg {
			x := (windowWidth - len(l)*smallFontSize) / 2
			text.Draw(screen, l, smallArcadeFont, x, windowHeight-4+(i-1)*smallFontSize, color.White)
		}
	}
}

//Checks if spacebar is pressed
func (g *Game) isSpacePressed() bool {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		return true
	}
	return false
}

func init() {
	tt, err := opentype.Parse(fonts.PressStart2P_ttf)
	if err != nil {
		log.Fatal(err)
	}
	const dpi = 72
	titleArcadeFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    titleFontSize,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
	arcadeFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    fontSize,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
	smallArcadeFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    smallFontSize,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
