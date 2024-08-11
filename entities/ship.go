package entities

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const SHIP_SIZE int = 100

type Ship struct {
	img *ebiten.Image
}

func NewShip(img_p string) *Ship {
	img, _, err := ebitenutil.NewImageFromFile(img_p)
	if err != nil {
		log.Fatal(err)
	}

	return &Ship{
		img: img,
	}
}

func (s *Ship) Update() {
}

func (s *Ship) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	screen.DrawImage(s.img, op)
}

func (s *Ship) Layout(outside_width, outside_height int) (int, int) {
	return outside_width, outside_height
}
