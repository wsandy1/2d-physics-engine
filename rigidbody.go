package main

import (
	"math"

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

type Triangle struct {
	One   vector2.Vec2
	Two   vector2.Vec2
	Three vector2.Vec2
	Area  float64
	CoM   vector2.Vec2
}

func (b *RigidBody) Accelerate(v vector2.Vec2) {
	b.acceleration = vector2.Add(b.acceleration, v)
}

func CalculateCentreOfMass(vertices []vector2.Vec2) vector2.Vec2 {
	triangles := []Triangle{}
	for i := 2; i < len(vertices); i++ {
		// make triangles by drawing a line to each non-adjacent vertex
		current_triangle := Triangle{One: vertices[0], Two: vertices[i-1], Three: vertices[i], Area: 0, CoM: vector2.Vec2{X: 0, Y: 0}}

		// calculate side lengths
		var side_lengths [3]float32
		side_lengths[0] = vector2.Magnitude(vector2.Sub(current_triangle.Two, current_triangle.One))
		side_lengths[1] = vector2.Magnitude(vector2.Sub(current_triangle.Three, current_triangle.Two))
		side_lengths[2] = vector2.Magnitude(vector2.Sub(current_triangle.Three, current_triangle.One))

		// calculate triangle area using Heron's formula area, where s is the semi-perimeter of a triangle with sides of length A, B, C; area = sqrt(s * (s-a)(s-b)(s-c))
		semi_perimeter := 0.5 * (side_lengths[0] + side_lengths[1] + side_lengths[2])
		current_triangle.Area = math.Sqrt(float64(semi_perimeter * (semi_perimeter - side_lengths[0]) * (semi_perimeter - side_lengths[1]) * (semi_perimeter - side_lengths[2])))

		// calculate the midpoint of the edge between the first and second points
		midpoint := vector2.Add(current_triangle.One, vector2.ConstMul(vector2.Sub(current_triangle.Two, current_triangle.One), 0.5))
		//calculate CoM, which is 1/3 of the way along the meridian
		current_triangle.CoM = vector2.Add(midpoint, vector2.ConstMul(vector2.Sub(current_triangle.Three, midpoint), 0.333333))

		triangles = append(triangles, current_triangle)
	}

	// calculate sum of triangle areas * by their respective CoM vectors
	// as well as the total area of the polygon
	var total_vec vector2.Vec2
	var total_area float64

	for _, t := range triangles {
		total_vec = vector2.Add(total_vec, vector2.ConstMul(t.CoM, float32(t.Area)))
		total_area += t.Area
	}

	// return overall CoM
	return vector2.ConstDiv(total_vec, float32(total_area))
}
