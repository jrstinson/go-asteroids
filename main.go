package main

import (
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/jrstinson/go-asteroids/entities"
	"github.com/jrstinson/go-asteroids/frame"
)

const GAME_W int = 1280
const GAME_H int = 960

const SHIP_SIZE int = 100

type Game struct {
	Space *frame.Space
	Ship  *entities.Ship
}

func (g *Game) Update() error {
	return nil

}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Space.Draw(screen)
	g.Ship.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return GAME_W, GAME_H
}

func main() {
	ebiten.SetWindowSize(GAME_W, GAME_H)
	ebiten.SetWindowTitle("Asteroids")

	ship := entities.NewShip("assets/ship.png")

	space := frame.NewSpace("assets/frame.png", image.Rect(0, 0, GAME_W, GAME_H), ship)

	game := &Game{
		Space: space,
		Ship:  ship,
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
