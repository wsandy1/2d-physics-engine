package main

import "github.com/wsandy1/2d-physics-engine/vector2"

type PointForce struct {
	origin vector2.Vec2 // where force is coming from
	dirmag vector2.Vec2 // where force is going
}
