package main

import (
	"gosdl/config"
	"gosdl/renderer"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(config.WindowWidth, config.WindowHeight)
	ebiten.SetWindowTitle(config.WindowTitle)
	ebiten.SetTPS(config.Framerate)
	if err := ebiten.RunGame(&renderer.Game{}); err != nil {
		log.Fatal(err)
	}
}
