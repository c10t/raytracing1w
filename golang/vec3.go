package main

import "math"

type Vec3 struct {
	X, Y, Z float64
}

func (v Vec3) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func (v Vec3) SquaredLength() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}
