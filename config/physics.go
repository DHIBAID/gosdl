package config

// Physics constants for the particle system
const (
	WallDamping        = 0.98 // Used for velocity damping on wall collisions
	FluidViscosity     = 0.5  // Used for velocity matching between particles
	AttractionForce    = 0.5  // Base attraction force constant
	Gravity            = 0.5  // Downward force applied to particles (positive for downward)
	InteractionRadius  = 7.0 // h: interaction radius (pixels)
	RepulsionStiffness = 1.0  // kRep: repulsion stiffness (lower for stability)
	GravityX           = 0.0  // gravity x-component
	GravityY           = 0.5  // gravity y-component (positive for downward)
	ParticleMass       = 1.0  // mass of each particle
)

func Timestep() float64 {
	return 1.0 / Framerate
}
