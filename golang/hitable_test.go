package main

import (
	"testing"
)

func TestHitTheWorldWithSphere(t *testing.T) {
	s1 := NewSphere(1, 0, 0, 1)
	s2 := NewSphere(0, 1, 0, 1)
	w := World{s1, s2}
	r := Ray{Origin: Vec3{2, -1, 0}, Direction: Vec3{-1, 1, 0}}
	hit, rec := w.Hit(r, 0, 9999)
	if !hit {
		t.Fatalf("The Ray should be hit to the given World")
	}
	if rec.At <= 0 {
		t.Fatalf("HitRecord.At should be positive: %v", rec.At)
	}
	if p := rec.Point; p.X <= 0 || p.Y >= 0 || p.Z != 0 {
		t.Fatalf("HitRecord.Point seems wrong: (%v, %v, %v)", p.X, p.Y, p.Z)
	}
	if n := rec.Normal; n.X <= 0 || n.Y >= 0 || n.Z != 0 {
		t.Fatalf("HitRecord.Normal seems wrong: (%v, %v, %v)", n.X, n.Y, n.Z)
	}
}
