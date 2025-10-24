package draw

import (
	"sdl-renderer/physics"

	"github.com/veandco/go-sdl2/sdl"
)

func StartDrawing(surface *sdl.Surface, window *sdl.Window, rect *sdl.Rect, posX *float32, posY *float32, velX *float32, velY *float32) {
	// fill background black
	surface.FillRect(nil, 0)

	// draw starting reference line at a fixed position
	referenceColor := sdl.MapRGB(surface.Format, 255, 255, 255)                                // white color
	surface.FillRect(&sdl.Rect{X: 75, Y: 100, W: int32(3 * rect.W / 2), H: 2}, referenceColor) // fixed Y position

	// apply gravity only when the rectangle is in free fall
	if *posY+float32(rect.H) < float32(surface.H) {
		physics.ApplyGravity(velY, 5.0)
	}

	// update rectangle position
	*posY += *velY
	*posX += *velX

	// bounce off walls and clamp position
	if *posX <= 0 {
		*posX = 0
		*velX = -*velX
	} else if *posX+float32(rect.W) >= float32(surface.W) {
		*posX = float32(surface.W - rect.W)
		*velX = -*velX
	}

	if *posY+float32(rect.H) >= float32(surface.H) {
		*posY = float32(surface.H - rect.H)
		*velY = -*velY // fully invert velocity for elastic collision
	} else if *posY <= 0 {
		*posY = 0
		*velY = -*velY
	}

	// update rect position for drawing
	rect.X = int32(*posX)
	rect.Y = int32(*posY)

	// draw rectangle
	pixel := sdl.MapRGB(surface.Format, 255, 0, 255)
	surface.FillRect(rect, pixel)
}
