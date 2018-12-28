package main

import "math"

type Sphere struct {
	Center Vec3
	Radius float64
}

func (s *Sphere) Hit(r Ray, tmin float64, tmax float64) (bool, HitRecord) {
	oc := r.Origin.Sub(s.Center)
	a := Dot(r.Direction, r.Direction)
	b := Dot(oc, r.Direction)
	c := Dot(oc, oc) - (s.Radius * s.Radius)
	discriminant := b*b - a*c

	if discriminant > 0 {
		candidate1 := (-b - math.Sqrt(discriminant)) / a

		if tmin < candidate1 && candidate1 < tmax {
			t := candidate1
			p := r.PointAt(candidate1)
			n := p.Sub(s.Center).Shrink(s.Radius)
			return true, HitRecord{At: t, Point: p, Normal: n}
		}

		candidate2 := (-b + math.Sqrt(discriminant)) / a

		if tmin < candidate2 && candidate2 < tmax {
			t := candidate2
			p := r.PointAt(candidate2)
			n := p.Sub(s.Center).Shrink(s.Radius)
			return true, HitRecord{At: t, Point: p, Normal: n}
		}
	}

	return false, HitRecord{}
}

func NewSphere(x, y, z, r float64) *Sphere {
	return &Sphere{Center: Vec3{X: x, Y: y, Z: z}, Radius: r}
}
