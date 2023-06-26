package vector2

import "math"

type Vec2 struct {
	x, y float32
}

// rotate by theta degrees about origin (positive = anticlockwise)
func (v *Vec2) Rotate(theta float64) Vec2 {
	return Vec2{
		v.x*float32(math.Cos(-theta)) + v.y*float32(-math.Sin(-theta)),
		v.x*float32(math.Sin(-theta)) + v.y*float32(math.Cos(-theta)),
	}
}

func Add(v1 Vec2, v2 Vec2) Vec2 {
	x := v1.x + v2.y
	y := v1.y + v2.y
	vr := Vec2{x, y}
	return vr
}

// subtract v2 from v1
// ie returns v1-v2
func Sub(v1 Vec2, v2 Vec2) Vec2 {
	x := v1.x - v2.x
	y := v1.y - v2.y
	vr := Vec2{x, y}
	return vr
}

func Mul(v1 Vec2, v2 Vec2) Vec2 {

	return Vec2{v1.x * v2.x, v1.y * v2.y}
}

func ConstMul(v1 Vec2, c float32) Vec2 {

	return Vec2{v1.x * c, v1.y * c}
}

// vec1 / vec2
func Div(v1 Vec2, v2 Vec2) Vec2 {

	return Vec2{v1.x / v2.x, v1.y / v2.y}
}

func ConstDiv(v1 Vec2, c float32) Vec2 {

	return Vec2{v1.x / c, v1.y / c}
}
