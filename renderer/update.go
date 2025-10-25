package renderer

import (
	"gosdl/config"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) Update() error {
	// Get mouse position and button state
	mouseX, mouseY := ebiten.CursorPosition()
	mousePressed := ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)

	// Draw mouse particle visualization

	// Define mouse particle properties
	var mouseParticle *Particle
	if mousePressed {
		mouseParticle = &Particle{
			X:    float64(mouseX),
			Y:    float64(mouseY),
			VX:   0,
			VY:   0,
			Size: 10, // Larger size for stronger interaction
		}
	}

	// Update particle positions and apply forces
	for i := range particles {
		p1 := &particles[i]

		// Calculate total force for the particle
		totalForceX := config.GravityX * config.ParticleMass
		totalForceY := config.GravityY * config.ParticleMass

		for j := range particles {
			if i != j {
				p2 := &particles[j]

				rijX := p1.X - p2.X
				rijY := p1.Y - p2.Y
				distanceSquared := rijX*rijX + rijY*rijY

				if distanceSquared <= config.InteractionRadius*config.InteractionRadius {
					distance := math.Sqrt(distanceSquared)
					if distance < 0.01 {
						distance = 0.01 // Clamp to avoid divide-by-zero
					}

					dirX := rijX / distance
					dirY := rijY / distance
					s := config.InteractionRadius - distance

					// Repulsion force
					fRepX := config.RepulsionStiffness * s * dirX
					fRepY := config.RepulsionStiffness * s * dirY

					// Viscosity force
					dvX := p1.VX - p2.VX
					dvY := p1.VY - p2.VY
					fViscX := -config.FluidViscosity * dvX * (s / config.InteractionRadius)
					fViscY := -config.FluidViscosity * dvY * (s / config.InteractionRadius)

					totalForceX += fRepX + fViscX
					totalForceY += fRepY + fViscY
				}
			}
		}

		// Interaction with mouse particle
		if mousePressed {
			rijX := p1.X - mouseParticle.X
			rijY := p1.Y - mouseParticle.Y
			distanceSquared := rijX*rijX + rijY*rijY

			if distanceSquared <= config.InteractionRadius*config.InteractionRadius {
				distance := math.Sqrt(distanceSquared)
				if distance < 0.01 {
					distance = 0.01 // Clamp to avoid divide-by-zero
				}

				dirX := rijX / distance
				dirY := rijY / distance
				s := config.InteractionRadius - distance

				// Repulsion force
				fRepX := config.RepulsionStiffness * s * dirX
				fRepY := config.RepulsionStiffness * s * dirY

				// Apply forces to particle
				totalForceX += fRepX
				totalForceY += fRepY
			}
		}

		// Integration (semi-implicit Euler)
		accelerationX := totalForceX / config.ParticleMass
		accelerationY := totalForceY / config.ParticleMass
		p1.VX += accelerationX * config.Timestep()
		p1.VY += accelerationY * config.Timestep()
		p1.X += p1.VX * config.Timestep()
		p1.Y += p1.VY * config.Timestep()

		// Ensure particles stay within the screen bounds
		if p1.X < p1.Size {
			p1.X = p1.Size
			p1.VX = -p1.VX * config.WallDamping // Reverse and damp velocity
		}
		if p1.X > 320-p1.Size {
			p1.X = 320 - p1.Size
			p1.VX = -p1.VX * config.WallDamping
		}
		if p1.Y < p1.Size {
			p1.Y = p1.Size
			p1.VY = -p1.VY * config.WallDamping
			// Force a small downward velocity if stuck
			if math.Abs(p1.VY) < 0.05 {
				p1.VY = -0.5
			}
		}
		if p1.Y > 240-p1.Size {
			p1.Y = 240 - p1.Size
			p1.VY = -p1.VY * config.WallDamping
		}

		// Always apply gravity after collision handling
		p1.VY += config.GravityY * config.Timestep()
	}

	// Ensure the mouse particle interacts with other particles
	if mousePressed {
		mouseParticle = &Particle{
			X:    float64(mouseX),
			Y:    float64(mouseY),
			VX:   0,
			VY:   0,
			Size: 10, // Larger size for stronger interaction
		}

		// Add mouse particle to the interaction loop
		for i := range particles {
			p1 := &particles[i]

			rijX := p1.X - mouseParticle.X
			rijY := p1.Y - mouseParticle.Y
			distanceSquared := rijX*rijX + rijY*rijY

			if distanceSquared <= config.InteractionRadius*config.InteractionRadius {
				distance := math.Sqrt(distanceSquared)
				if distance < 0.01 {
					distance = 0.01 // Clamp to avoid divide-by-zero
				}

				dirX := rijX / distance
				dirY := rijY / distance
				s := config.InteractionRadius - distance

				// Repulsion force
				fRepX := config.RepulsionStiffness * s * dirX
				fRepY := config.RepulsionStiffness * s * dirY

				// Apply forces to particle
				p1.VX += fRepX / config.ParticleMass
				p1.VY += fRepY / config.ParticleMass
			}
		}
	}

	return nil
}
