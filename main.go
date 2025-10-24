package main

import (
	"log"
	"sdl-renderer/config"
	"sdl-renderer/events"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

func main() {
	log.Println("Starting SDL2 example")

	if err := sdl.Init(sdl.INIT_VIDEO); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	if err := ttf.Init(); err != nil {
		panic(err)
	}
	defer ttf.Quit()

	window, err := sdl.CreateWindow(config.WindowTitle, config.WindowX, config.WindowY, config.WindowWidth, config.WindowHeight, config.WindowFlags)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	// get the surface
	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}

	events.Loop(surface, window)
}
