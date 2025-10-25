package renderer

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

// DrawSmoothCircle draws an antialiased circle using a pre-rendered texture.
func DrawSmoothCircle(screen *ebiten.Image, x, y, radius float64, col color.Color) {
	circleImage := ebiten.NewImage(int(radius*2), int(radius*2))
	circleImage.Fill(col)

	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(x-radius, y-radius)
	screen.DrawImage(circleImage, opts)
}
