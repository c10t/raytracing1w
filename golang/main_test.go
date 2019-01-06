package main

import (
	"fmt"
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

	lookF := Vec3{3, 3, 2}
	lookA := Vec3{0, 0, -1}
	distToFocus := Sub(lookF, lookA).Length()

	vup := Vec3{0, 1, 0}
	cam := NewVerticalCamera(lookF, lookA, vup, 90, 2, 2, distToFocus)

	u := float64(50) / float64(200)
	v := float64(50) / float64(100)
	r := cam.GetRay(u, v)

	c := color(&r, &world, 0)
	if c.SquaredLength() == 0 {
		t.Fatalf("color seems strange: (%v, %v, %v)", c.X, c.Y, c.Z)
	}
}

func TestLerp(t *testing.T) {
	nx := 100
	ny := 50
	ns := 50

	result := lerp(nx, ny, ns)

	if result[0] != "P3" {
		t.Fatalf("first line should be 'P3': %v", result[0])
	}
	if result[1] != fmt.Sprintf("%v %v", nx, ny) {
		t.Fatalf("second line should be 'nx ny': %v", result[1])
	}
	if result[2] != "255" {
		t.Fatalf("third line should be '255': %v", result[2])
	}
	for i := 3; i < nx*ny+3; i++ {
		if line := result[i]; len(line) > 11 {
			t.Fatalf("line %v should be triplet of three digit numbers: %v", len(line), line)
		}
	}
}

func BenchmarkColorAtPoint(b *testing.B) {
	b.ResetTimer()

	nx := 200
	ny := 100
	ns := 100

	result := make([]string, nx*ny+3)
	result[0] = "P3"
	result[1] = fmt.Sprintf("%d %d", nx, ny)
	result[2] = "255"

	world := makeTheWorld()
	lookF := Vec3{3, 1, 2}
	lookA := Vec3{0, 0, -1}
	distToFocus := Sub(lookF, lookA).Length()
	aperture := 0.1

	vup := Vec3{0, 1, 0}
	aspect := float64(nx) / float64(ny)
	cam := NewVerticalCamera(lookF, lookA, vup, 90, aspect, aperture, distToFocus)

	i := 10
	j := 10

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		colorAtPoint(result, nx, ny, ns, i, j, cam, &world)
	}
}

func BenchmarkMakeTheWorld(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		makeTheWorld()
	}
}

func BenchmarkLerp(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		lerp(200, 100, 100)
	}
}
