package physics

func ApplyGravity(velY *float32, gravity float32) {
	*velY += gravity
}