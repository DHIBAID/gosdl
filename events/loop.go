package events

import (
	"fmt"
	"sdl-renderer/config"
	"sdl-renderer/draw"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

func Loop(surface *sdl.Surface, window *sdl.Window) {
	running := true

	// initialize rectangle position and velocity
	var posX, posY float32 = 100.0, 100.0
	var velX, velY float32 = 0.0, -5.0
	rect := sdl.Rect{W: 200, H: 200}

	// Starting line reference
	referenceColor := sdl.MapRGB(surface.Format, 255, 255, 255)                  // white color
	surface.FillRect(&sdl.Rect{X: 100, Y: 100, W: rect.W, H: 2}, referenceColor) // fixed Y position

	// load font
	font, err := ttf.OpenFont("./fonts/firacodenerd.ttf", 24) // Replace with the actual font path
	if err != nil {
		panic(err)
	}
	defer font.Close()

	var frameCount int
	var startTime uint32 = uint32(sdl.GetTicks64())

	for running {
		// handle events
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			if _, ok := event.(*sdl.QuitEvent); ok {
				running = false
				window.Destroy()
				return
			}
		}

		// draw frame
		draw.StartDrawing(surface, window, &rect, &posX, &posY, &velX, &velY)

		// calculate FPS
		frameCount++
		elapsedTime := uint32(sdl.GetTicks64()) - startTime
		fps := float64(frameCount) / (float64(elapsedTime) / 1000.0)

		// render FPS counter with antialiasing
		fpsText := fmt.Sprintf("FPS: %.2f", fps)
		fpsSurface, err := font.RenderUTF8Blended(fpsText, sdl.Color{R: 255, G: 255, B: 255, A: 255})
		if err != nil {
			panic(err)
		}
		defer fpsSurface.Free()

		fpsRect := sdl.Rect{X: 10, Y: 10, W: fpsSurface.W, H: fpsSurface.H}
		fpsSurface.Blit(nil, surface, &fpsRect)

		// update window every frame
		window.UpdateSurface()

		// delay to control frame rate
		sdl.Delay(1000 / config.Framerate)
	}
}
