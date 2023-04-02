package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

const (
	screenWidth  = 1280
	screenHeight = 720
)

var (
	normalFont font.Face
)

func init() {
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 200
	normalFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingVertical,
	})

	if err != nil {
		log.Fatal(err)
	}
}

type Game struct {
	dialogues []string
	index     int
	counter   int
}

func NewGame(dialogues []string) *Game {
	return &Game{dialogues: dialogues}
}

func (g *Game) Update() error {
	g.counter++
	if g.counter >= 60 {
		g.index++
		g.index %= len(g.dialogues)
		g.counter = 0
	}
	return nil
}
func (g *Game) Draw(screen *ebiten.Image) {
	bound := text.BoundString(normalFont, g.dialogues[g.index])
	text.Draw(screen, fmt.Sprintf("%v %v", bound.Max.X, bound.Max.Y), normalFont, 0, 160, color.White)

	heightSize := bound.Dy()
	log.Println(bound.Bounds(), heightSize)

	text.Draw(screen, g.dialogues[g.index], normalFont, 0, -bound.Min.Y, color.White)
}
func (g *Game) Layout(outsideWidth, outsizeHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Message Window")

	game := NewGame([]string{
		"こんにちは。私は ganyariya です。",
		"これは ebiten でノベルゲームのテキストを表示するテストです。",
		"どのように実装していくのがいいんだろう...",
	})
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
