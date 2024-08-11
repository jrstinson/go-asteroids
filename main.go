package main

import (
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"github.com/jrstinson/go-asteroids/frame"
)

const GAME_W int = 1280
const GAME_H int = 960

const SHIP_SIZE int = 100

type Game struct{}

func (g *Game) Update() error {
	return nil

}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Asteroids")
	space := frame.NewSpace("assets/frame.png", image.Rect(0, 0, GAME_W, GAME_H))

	space.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return GAME_W, GAME_H
}

func main() {
	ebiten.SetWindowTitle("Asteroids!")
	ebiten.SetWindowSize(GAME_W, GAME_H)
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
