package main

import (
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"github.com/jrstinson/go-asteroids/frame"
)

const GAME_W int = 6400
const GAME_H int = 4800

const SHIP_SIZE int = 100

func init() {

}

type Game struct{}

func (g *Game) Update() error {
	return nil

}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Asteroids")
	space := frame.NewSpace("assets/tile.png", image.Rect(0, 0, GAME_W, GAME_H), 0)
	space.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetWindowTitle("Asteroids!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
