package main

import (
	"math"
	"testing"
)

func TestLength(t *testing.T) {
	v := Vec3{X: 3.0, Y: 4.0, Z: 5.0}
	if v.Length() != math.Sqrt(50) {
		t.Fatalf("Length should be return sqrt(50) for the vector (3, 4, 5)")
	}
}

func TestSquaredLength(t *testing.T) {
	v := Vec3{X: 3.0, Y: 4.0, Z: 5.0}
	if v.SquaredLength() != 50 {
		t.Fatalf("SquaredLength should be return 50 for the vector (3, 4, 5)")
	}
}
