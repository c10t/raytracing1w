package main

import (
	"testing"
)

func TestGetRandomInUnitSphere(t *testing.T) {
	p := RandomInUnitSphere()
	if p.SquaredLength() >= 1 {
		t.Fatalf("The point is out of unit sphere: (%v, %v, %v)", p.X, p.Y, p.Z)
	}
}

func TestColor(t *testing.T) {
	s1 := NewSphere(0, 0, -1, 0.5, Metal{Albedo: Vec3{0.8, 0.3, 0.3}})
	s2 := NewSphere(0, -100.5, -1, 100, Metal{Albedo: Vec3{0.8, 0.6, 0.2}})
	world := World{s1, s2}

	cam := NewCamera()
	u := float64(50) / float64(200)
	v := float64(50) / float64(100)
	r := cam.GetRay(u, v)

	c := color(&r, &world, 0)
	if c.SquaredLength() == 0 {
		t.Fatalf("color seems strange: (%v, %v, %v)", c.X, c.Y, c.Z)
	}
}
