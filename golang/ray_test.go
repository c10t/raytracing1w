package main

import (
	"testing"
)

func TestPointAt(t *testing.T) {
	o := Vec3{X: 3, Y: 4, Z: 5}
	d := Vec3{X: 0, Y: 1, Z: 2}
	r := Ray{Origin: o, Direction: d}
	p := r.PointAt(1.5)
	if p.X != 3 || p.Y != 5.5 || p.Z != 8 {
		t.Fatalf("PointAt(1.5) should be (3, 5.5, 8) for the given Ray but actual is (%v, %v, %v)", p.X, p.Y, p.Z)
	}
}
