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

func TestUnitVector(t *testing.T) {}

func TestAdd(t *testing.T) {
	v := Vec3{X: 1.5, Y: -0.2, Z: -1}
	w := Add(v, UnitVector)
	if w.X != 2.5 || w.Y != 0.8 || w.Z != 0 {
		t.Fatalf("(1.5, -0.2, -1) + (1, 1, 1) != (2.5, 0.8, 0)")
	}
}

func TestSub(t *testing.T) {
	v := Vec3{X: 0.5, Y: 1, Z: 1.7}
	w := Sub(v, UnitVector)
	if w.X != -0.5 || w.Y != 0 || w.Z != 0.7 {
		t.Fatalf("(0.5, 1, 1.7) - (1, 1, 1) != (-0.5, 0, 0.7)")
	}
}

func TestSlide(t *testing.T) {
	v := Vec3{X: 1, Y: -1, Z: 0}
	w := v.Slide(-1)
	if w.X != 0 || w.Y != -2 || w.Z != -1 {
		t.Fatalf("(1, -1, 0) + (-1) != (0, -2, -1): (%v, %v, %v)", w.X, w.Y, w.Z)
	}
}

func TestScale(t *testing.T) {
	v := Vec3{X: 1, Y: -1, Z: 0}
	w := v.Scale(-1.5)
	if w.X != -1.5 || w.Y != 1.5 || w.Z != 0 {
		t.Fatalf("-1.5 * (1, -1, 0) != (-1.5, 1.5, 0): (%v, %v, %v)", w.X, w.Y, w.Z)
	}
}

func TestShrink(t *testing.T) {
	v := Vec3{X: 3, Y: -3, Z: 0}
	w := v.Shrink(-1.5)
	if w.X != -2 || w.Y != 2 || w.Z != 0 {
		t.Fatalf("(3, -3, 0) / (-1.5) != (-2, 2, 0): (%v, %v, %v)", w.X, w.Y, w.Z)
	}
}

func TestAddMethod(t *testing.T) {
	v := Vec3{X: 1.5, Y: -0.2, Z: -1}
	w := v.Add(UnitVector)
	if w.X != 2.5 || w.Y != 0.8 || w.Z != 0 {
		t.Fatalf("(1.5,-0.2,-1) + (1,1,1) != (2.5,0.8,0): (%v,%v,%v)", w.X, w.Y, w.Z)
	}
}

func TestSubMethod(t *testing.T) {
	v := Vec3{X: -1.5, Y: 0.2, Z: 1}
	w := v.Sub(UnitVector)
	if w.X != -2.5 || w.Y != -0.8 || w.Z != 0 {
		t.Fatalf("(-1.5,0.2,1) - (1,1,1) != (-2.5,-0.8,0): (%v,%v,%v)", w.X, w.Y, w.Z)
	}
}

func TestMulMethod(t *testing.T) {
	v := Vec3{X: -1.5, Y: 0.2, Z: 1}
	w := v.Mul(Vec3{-2, -10, 0})
	if w.X != 3 || w.Y != -2 || w.Z != 0 {
		t.Fatalf("(-1.5,0.2,1) * (-2,-10,0) != (3,-2,0): (%v,%v,%v)", w.X, w.Y, w.Z)
	}
}

func TestDot(t *testing.T) {
	v := Vec3{X: -1, Y: 2, Z: 3}
	w := Vec3{X: 4, Y: 5, Z: 0}
	if d := Dot(v, w); d != 6 {
		t.Fatalf("(-1,2,3) . (4,5,0) != 6 %v", d)
	}
}
