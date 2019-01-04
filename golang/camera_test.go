package main

import (
	"testing"
)

func TestGetRay(t *testing.T) {
	lookF := Vec3{3, 3, 2}
	lookA := Vec3{0, 0, -1}
	distToFocus := Sub(lookF, lookA).Length()

	vup := Vec3{0, 1, 0}
	cam := NewVerticalCamera(lookF, lookA, vup, 90, 2, 2, distToFocus)

	r := cam.GetRay(0.5, 0.5)
	if o := r.Origin; o.X == 0 || o.Y == 0 || o.Z == 0 {
		t.Fatalf("Camera.GetRay.Origin seems wrong: (%v, %v, %v)", o.X, o.Y, o.Z)
	}
	if d := r.Direction; d.X == 0 || d.Y == 0 || d.Z == -1 {
		t.Fatalf("Camera.GetRay.Direction seems wrong: (%v, %v, %v)", d.X, d.Y, d.Z)
	}
}
