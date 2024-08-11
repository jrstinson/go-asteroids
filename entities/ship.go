package entities

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const SHIP_SIZE int = 100

type Ship struct {
	img   *ebiten.Image
	angle float64
}

func NewShip(img_p string) *Ship {
	img, _, err := ebitenutil.NewImageFromFile(img_p)
	if err != nil {
		log.Fatal(err)
	}

	return &Ship{
		img:   img,
		angle: 0,
	}
}

func (s *Ship) Update() {

}

func (s *Ship) Draw(screen *ebiten.Image) {
	//draw ship at center of screen
	op := &ebiten.DrawImageOptions{}

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		s.angle -= 0.1
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		s.angle += 0.1
	}

	op.GeoM.Reset()

	op.GeoM.Translate(-float64(SHIP_SIZE/2), -float64(SHIP_SIZE/2))

	op.GeoM.Rotate(s.angle)

	op.GeoM.Translate(float64(screen.Bounds().Dx()/2), float64(screen.Bounds().Dy()/2)+float64(SHIP_SIZE))

	screen.DrawImage(s.img, op)
}

func (s *Ship) Layout(outside_width, outside_height int) (int, int) {
	return outside_width, outside_height
}
