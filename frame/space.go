package frame

import (
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Space struct {
	tile_img    *ebiten.Image
	tile_buffer int
	view_area   *ebiten.Image
	play_area   *ebiten.Image
	player_pos  *image.Point
}

func NewSpace(img_p string, max_s image.Rectangle, tile_buff int) *Space {
	var err error
	var ti *ebiten.Image

	ti, _, err = ebitenutil.NewImageFromFile(img_p)

	if err != nil {
		log.Fatal(err)
	}

	play_area := ebiten.NewImage(max_s.Dx(), max_s.Dy())

	// draw the center tile of the play area
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(max_s.Dx()/2), float64(max_s.Dy()/2))
	play_area.DrawImage(ti, op)

	view_area := play_area.SubImage(image.Rect(0, 0, max_s.Dx(), max_s.Dy())).(*ebiten.Image)

	return &Space{
		tile_img:    ti,
		tile_buffer: tile_buff,
		view_area:   view_area,
		play_area:   play_area,
		player_pos:  &image.Point{max_s.Dx() / 2, max_s.Dy() / 2},
	}
}

func (s *Space) Update() {
}

func (s *Space) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0)
	screen.DrawImage(s.view_area, op)
}

func (s *Space) Layout(outside_width, outside_height int) (int, int) {
	return outside_width, outside_height
}
