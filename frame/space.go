package frame

import (
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"github.com/jrstinson/go-asteroids/entities"
	"github.com/jrstinson/go-asteroids/util"
)

const UNIVERSE_W int = 6400
const UNIVERSE_H int = 4800

type Space struct {
	play_area *ebiten.Image
	view_area *ebiten.Image
	origin    image.Point
	ship      *entities.Ship
}

func NewSpace(img_p string, view_s image.Rectangle, ship *entities.Ship) *Space {
	img, _, err := ebitenutil.NewImageFromFile(img_p)
	if err != nil {
		log.Fatal(err)
	}

	play_area := ebiten.NewImage(UNIVERSE_W, UNIVERSE_H)

	op := &ebiten.DrawImageOptions{}

	for x := 0; x < UNIVERSE_W; x += img.Bounds().Dx() {
		for y := 0; y < UNIVERSE_H; y += img.Bounds().Dy() {
			op.GeoM.Reset()
			op.GeoM.Translate(float64(x), float64(y))
			play_area.DrawImage(img, op)
		}
	}

	view_area_rect := image.Rect(UNIVERSE_W/2-view_s.Max.X/2, UNIVERSE_H/2-view_s.Max.Y/2, UNIVERSE_W/2+view_s.Max.X/2, UNIVERSE_H/2+view_s.Max.Y/2)

	view_area := play_area.SubImage(view_area_rect).(*ebiten.Image)

	origin := image.Point{X: UNIVERSE_W / 2, Y: UNIVERSE_H / 2}

	return &Space{
		play_area: play_area,
		view_area: view_area,
		origin:    origin,
		ship:      ship,
	}
}

func (s *Space) Update() {
}

func (s *Space) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Reset()

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		// create a vector pointing up with a magnitude of 5
		// rotate the vector by the angle of the ship
		// add the vector to origin
		// set origin to the new point

		v := util.Vector{X: float64(s.origin.X), Y: float64(s.origin.Y) - 10}

		ship := *s.ship

		v = v.RotateAround(ship.GetAngle(), util.Vector{X: float64(s.origin.X), Y: float64(s.origin.Y)})

		s.origin = image.Point{X: int(v.X), Y: int(v.Y)}

		// if moving the origin would put any part of the view area outside of the play area, move the ship within the play area
		if s.origin.X-s.view_area.Bounds().Dx()/2 < 0 {
			s.origin.X = s.view_area.Bounds().Dx() / 2
		}

		if s.origin.Y-s.view_area.Bounds().Dy()/2 < 0 {
			s.origin.Y = s.view_area.Bounds().Dy() / 2
		}

		if s.origin.X+s.view_area.Bounds().Dx()/2 > s.play_area.Bounds().Dx() {
			s.origin.X = s.play_area.Bounds().Dx() - s.view_area.Bounds().Dx()/2
		}

		if s.origin.Y+s.view_area.Bounds().Dy()/2 > s.play_area.Bounds().Dy() {
			s.origin.Y = s.play_area.Bounds().Dy() - s.view_area.Bounds().Dy()/2
		}

	}

	view_area_rect := image.Rect(s.origin.X-s.view_area.Bounds().Dx()/2, s.origin.Y-s.view_area.Bounds().Dy()/2, s.origin.X+s.view_area.Bounds().Dx()/2, s.origin.Y+s.view_area.Bounds().Dy()/2)

	if view_area_rect.In(s.play_area.Bounds()) {
		s.view_area = s.play_area.SubImage(view_area_rect).(*ebiten.Image)
	}
	screen.DrawImage(s.view_area, op)
}

func (s *Space) Layout(outside_width, outside_height int) (int, int) {
	return outside_width, outside_height
}

func (s *Space) GetViewArea() *ebiten.Image {
	return s.view_area
}

func (s *Space) GetPlayArea() *ebiten.Image {
	return s.play_area
}
