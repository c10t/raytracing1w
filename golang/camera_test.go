package main

import (
	"testing"
)

func TestGetRay(t *testing.T) {
	c := NewCamera()
	r := c.GetRay(0.5, 0.5)
	if o := r.Origin; o.X != 0 || o.Y != 0 || o.Z != 0 {
		t.Fatalf("Camera.GetRay.Origin seems wrong: (%v, %v, %v)", o.X, o.Y, o.Z)
	}
	if d := r.Direction; d.X != 0 || d.Y != 0 || d.Z != -1 {
		t.Fatalf("Camera.GetRay.Direction seems wrong: (%v, %v, %v)", d.X, d.Y, d.Z)
	}
}
