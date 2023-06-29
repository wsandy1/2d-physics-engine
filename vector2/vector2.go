package vector2

import "math"

type Vec2 struct {
	X, Y float32
}

// rotate by theta degrees about origin (positive = anticlockwise)
func (v *Vec2) Rotate(theta float64) Vec2 {
	return Vec2{
		v.X*float32(math.Cos(-theta)) + v.Y*float32(-math.Sin(-theta)),
		v.X*float32(math.Sin(-theta)) + v.Y*float32(math.Cos(-theta)),
	}
}

func Add(v1 Vec2, v2 Vec2) Vec2 {
	x := v1.X + v2.X
	y := v1.Y + v2.Y
	vr := Vec2{x, y}
	return vr
}

// subtract v2 from v1
// ie returns v1-v2
func Sub(v1 Vec2, v2 Vec2) Vec2 {
	x := v1.X - v2.X
	y := v1.Y - v2.Y
	vr := Vec2{x, y}
	return vr
}

func Mul(v1 Vec2, v2 Vec2) Vec2 {

	return Vec2{v1.X * v2.X, v1.Y * v2.Y}
}

func ConstMul(v1 Vec2, c float32) Vec2 {

	return Vec2{v1.X * c, v1.Y * c}
}

// vec1 / vec2
func Div(v1 Vec2, v2 Vec2) Vec2 {

	return Vec2{v1.X / v2.X, v1.Y / v2.Y}
}

func ConstDiv(v1 Vec2, c float32) Vec2 {

	return Vec2{v1.X / c, v1.Y / c}
}

// calculate magnitude of a vector
func Magnitude(v Vec2) float32 {
	return float32(math.Sqrt(math.Pow(float64(v.X), 2) + math.Pow(float64(v.Y), 2)))
}

// gives the dot product of two vectors
func Dot(v1 Vec2, v2 Vec2) float32 {
	return v1.X*v2.X + v1.Y*v2.Y
}
