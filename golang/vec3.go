package main

import "math"

type Vec3 struct {
	X, Y, Z float64
}

var UnitVector = Vec3{X: 1, Y: 1, Z: 1}

func (v Vec3) Length() float64 {
	return math.Sqrt(v.SquaredLength())
}

func (v Vec3) SquaredLength() float64 {
	return Dot(v, v)
}

func (v Vec3) UnitVector() Vec3 {
	return v.Shrink(v.Length())
}

func Add(v Vec3, w Vec3) Vec3 {
	return Vec3{v.X + w.X, v.Y + w.Y, v.Z + w.Z}
}

func Sub(v Vec3, w Vec3) Vec3 {
	return Vec3{v.X - w.X, v.Y - w.Y, v.Z - w.Z}
}

func (v Vec3) Slide(t float64) Vec3 {
	return Vec3{v.X + t, v.Y + t, v.Z + t}
}

func (v Vec3) Scale(t float64) Vec3 {
	return Vec3{v.X * t, v.Y * t, v.Z * t}
}

func (v Vec3) Shrink(t float64) Vec3 {
	return Vec3{v.X / t, v.Y / t, v.Z / t}
}

func (v Vec3) Add(w Vec3) Vec3 {
	return Vec3{v.X + w.X, v.Y + w.Y, v.Z + w.Z}
}

func (v Vec3) Sub(w Vec3) Vec3 {
	return Vec3{v.X - w.X, v.Y - w.Y, v.Z - w.Z}
}

func (v Vec3) Mul(w Vec3) Vec3 {
	return Vec3{v.X * w.X, v.Y * w.Y, v.Z * w.Z}
}

func Dot(v Vec3, w Vec3) float64 {
	return v.X*w.X + v.Y*w.Y + v.Z*w.Z
}
