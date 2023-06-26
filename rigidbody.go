package main

import (
	"github.com/wsandy1/2d-physics-engine/graphics"
	"github.com/wsandy1/2d-physics-engine/vector2"
)

// rigid body with positions and rotations for verlet integration
// centre of mass is at {0,0}
type RigidBody struct {
	mass float32
	moi  float32 // moment of inertia

	current_position vector2.Vec2
	last_position    vector2.Vec2

	acceleration  vector2.Vec2
	rotation      float32
	last_rotation float32

	//vertices     []vector2.Vec2
	poly         graphics.Polygon
	point_forces []PointForce
}

func (b *RigidBody) Accelerate(v vector2.Vec2) {
	b.acceleration = vector2.Add(b.acceleration, v)
}
