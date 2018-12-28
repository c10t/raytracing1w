package main

import (
	"testing"
)

func TestHitAtTwoPoint(t *testing.T) {
	s := NewSphere(0, 0, 0, 1)
	r := Ray{Origin: Vec3{0.5, 0, -10}, Direction: Vec3{0, 0, 1}}
	isHit, rec := s.Hit(r, 0, 99999)
	if !isHit {
		t.Fatalf("The Ray should be hit to the given Sphere")
	}
	if rec.At < 0 {
		t.Fatalf("HitRecord.At should be positive: %v", rec.At)
	}
	if p := rec.Point; p.X != 0.5 || p.Y != 0 || p.Z >= 0 {
		t.Fatalf("HitRecord.Point seems wrong: (%v, %v, %v)", p.X, p.Y, p.Z)
	}
	if n := rec.Normal; n.X == 0 || n.Y != 0 || n.Z == 0 {
		t.Fatalf("HitRecord.Normal seems wrong: (%v, %v, %v)", n.X, n.Y, n.Z)
	}
}

func TestHitAtOnePoint(t *testing.T) {
	s := NewSphere(0, 0, 0, 1)
	r := Ray{Origin: Vec3{0, 0.5, 0}, Direction: Vec3{0, 0, 1}}
	isHit, rec := s.Hit(r, 0, 99999)
	if !isHit {
		t.Fatalf("The Ray should be hit to the given Sphere")
	}
	if rec.At < 0 {
		t.Fatalf("HitRecord.At should be positive: %v", rec.At)
	}
	if p := rec.Point; p.X != 0 || p.Y != 0.5 || p.Z <= 0 {
		t.Fatalf("HitRecord.Point seems wrong: (%v, %v, %v)", p.X, p.Y, p.Z)
	}
	if n := rec.Normal; n.X != 0 || n.Y == 0 || n.Z == 0 {
		t.Fatalf("HitRecord.Normal seems wrong: (%v, %v, %v)", n.X, n.Y, n.Z)
	}
}

func TestNoHitOnBoundary(t *testing.T) {
	s := NewSphere(0, 0, 0, 1)
	r := Ray{Origin: Vec3{0, 1, -10}, Direction: Vec3{0, 0, 1}}
	isHit, _ := s.Hit(r, 0, 99999)
	if isHit {
		t.Fatalf("The Ray should not be hit to the given Sphere")
	}
}
