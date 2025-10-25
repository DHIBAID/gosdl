package renderer

import (
	"fmt"
	"gosdl/config"
	"image/color"
	"math"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Particle struct {
	X, Y   float64
	VX, VY float64
	Size   float64
}

// Initialize particles
var particles []Particle

func init() {
	rand.NewSource(time.Now().UnixNano())
	for i := 0; i < 500; i++ {
		particles = append(particles, Particle{
			X:    rand.Float64() * 320,
			Y:    rand.Float64() * 240,
			VX:   0,
			VY:   0,
			Size: 1,
		})
	}
}

func (p *Particle) CollidesWith(other *Particle) bool {
	dx := p.X - other.X
	dy := p.Y - other.Y
	distance := dx*dx + dy*dy
	radiusSum := p.Size + other.Size
	return distance < radiusSum*radiusSum
}

// Use the same rendering logic for fluid particles and mouse particle
func DrawMouseParticle(screen *ebiten.Image, mouseX, mouseY int) {
	mouseColor := color.RGBA{200, 150, 255, 255} // Light purple
	mouseSize := 10                              // Same size as the mouse particle

	if config.UseAntialiasing {
		DrawSmoothCircle(screen, float64(mouseX), float64(mouseY), float64(mouseSize), mouseColor)
	} else {
		for dx := -int(mouseSize); dx <= int(mouseSize); dx++ {
			for dy := -int(mouseSize); dy <= int(mouseSize); dy++ {
				if dx*dx+dy*dy <= int(mouseSize*mouseSize) {
					screen.Set(mouseX+dx, mouseY+dy, mouseColor)
				}
			}
		}
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Print the FPS to the terminal, reusing the same line
	// Print the FPS to the terminal, reusing the same line
	fps := ebiten.ActualTPS()
	fmt.Printf("\rFPS: %.2f", fps)

	// Draw the mouse particle
	mouseX, mouseY := ebiten.CursorPosition()
	DrawMouseParticle(screen, mouseX, mouseY)

	// Draw particles
	for i := range particles {
		p := &particles[i]
		p.X += p.VX
		p.Y += p.VY

		// Bounce off walls
		if p.X-p.Size < 0 || p.X+p.Size > 320 {
			p.VX = -p.VX
		}
		if p.Y-p.Size < 0 || p.Y+p.Size > 240 {
			p.VY = -p.VY
		}

		// Check for collisions with other particles
		for j := range particles {
			if i != j && p.CollidesWith(&particles[j]) {
				// Simple elastic collision response
				p.VX, particles[j].VX = particles[j].VX, p.VX
				p.VY, particles[j].VY = particles[j].VY, p.VY
			}
		}

		// Assign color based on velocity
		speed := math.Sqrt(p.VX*p.VX + p.VY*p.VY)
		col := velocityColor(speed)

		if config.UseAntialiasing {
			DrawSmoothCircle(screen, p.X, p.Y, p.Size, col)
		} else {
			for dx := -int(p.Size); dx <= int(p.Size); dx++ {
				for dy := -int(p.Size); dy <= int(p.Size); dy++ {
					if dx*dx+dy*dy <= int(p.Size*p.Size) {
						screen.Set(int(p.X)+dx, int(p.Y)+dy, col)
					}
				}
			}
		}
	}
}

// velocityColor maps speed to a gradient: blue (slow) -> green (moderate) -> red (fast)
func velocityColor(speed float64) color.Color {
	// Define speed thresholds
	minSpeed := 0.0
	midSpeed := 1.5
	maxSpeed := 4.0

	var r, g, b uint8
	if speed <= midSpeed {
		// Interpolate blue to green
		t := float64((speed - minSpeed) / (midSpeed - minSpeed))
		r = uint8(0*(1-t) + 0*t)
		g = uint8(0*(1-t) + 255*t)
		b = uint8(255*(1-t) + 0*t)
	} else if speed <= maxSpeed {
		// Interpolate green to red
		t := float64((speed - midSpeed) / (maxSpeed - midSpeed))
		r = uint8(0*(1-t) + 255*t)
		g = uint8(255*(1-t) + 0*t)
		b = uint8(0)
	} else {
		// Clamp to red
		r, g, b = 255, 0, 0
	}
	return color.RGBA{r, g, b, 255}
}
