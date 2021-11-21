package AirWars2D

import "AirWars2D/Math"

type Entity struct {
	velocity Math.Vector2D
}

func (e* Entity) SetVelocity(velocity Math.Vector2D) {
	e.velocity = velocity
}

func (e* Entity) GetVelocity() Math.Vector2D {
	return e.velocity
}
