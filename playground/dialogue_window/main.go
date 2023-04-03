package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 1280
	screenHeight = 720
)

var (
	emptyImage = ebiten.NewImage(3, 3)
)

func init() {
	emptyImage.Fill(color.RGBA{R: 200, B: 100, G: 10, A: 255})
	emptyImage.Fill(color.White)
}

type Game struct {
}

func NewGame() *Game {
	return &Game{}
}

func (g *Game) Update() error {
	return nil
}
func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DrawRect(screen, 0, 0, screenWidth/2, screenHeight/2, color.RGBA{R: 200, A: 255, G: 200, B: 30})
	ebitenutil.DrawRect(screen, screenWidth/2, screenHeight/2, screenWidth/2, screenHeight/2, color.White)
}
func (g *Game) Layout(outsideWidth, outsizeHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.RunGame(NewGame())
}
