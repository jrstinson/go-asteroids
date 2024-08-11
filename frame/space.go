package frame

import (
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const UNIVERSE_W int = 6400
const UNIVERSE_H int = 4800

type Space struct {
	play_area  *ebiten.Image
	view_area  *ebiten.Image
	player_pos *image.Point
}

func NewSpace(img_p string, view_s image.Rectangle) *Space {
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

	return &Space{
		play_area:  play_area,
		view_area:  view_area,
		player_pos: &image.Point{UNIVERSE_W / 2, UNIVERSE_H / 2},
	}
}

func (s *Space) Update() {
}

func (s *Space) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	screen.DrawImage(s.play_area, op)
}

func (s *Space) Layout(outside_width, outside_height int) (int, int) {
	return outside_width, outside_height
}

func (s *Space) GetPlayerPos() *image.Point {
	return s.player_pos
}

func (s *Space) SetPlayerPos(p *image.Point) {
	s.player_pos = p
}

func (s *Space) GetViewArea() *ebiten.Image {
	return s.view_area
}

func (s *Space) GetPlayArea() *ebiten.Image {
	return s.play_area
}
