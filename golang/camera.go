package main

import (
	"math"
)

type Camera struct {
	LowerLeftCorner Vec3
	Horizontal      Vec3
	Vertical        Vec3
	Origin          Vec3
}

func NewCamera() *Camera {
	l := Vec3{-2, -1, -1}
	h := Vec3{4, 0, 0}
	v := Vec3{0, 2, 0}
	o := Vec3{0, 0, 0}
	return &Camera{l, h, v, o}
}

func NewVerticalCamera(vfov, aspect float64) *Camera {
	theta := vfov * math.Pi / 180
	halfHeight := math.Tan(theta / 2)
	halfWidth := aspect * halfHeight
	l := Vec3{X: -halfWidth, Y: -halfHeight, Z: -1.0}
	h := Vec3{X: 2 * halfWidth, Y: 0, Z: 0}
	v := Vec3{X: 0, Y: 2 * halfHeight, Z: 0}
	o := Vec3{X: 0, Y: 0, Z: 0}

	return &Camera{LowerLeftCorner: l, Horizontal: h, Vertical: v, Origin: o}
}

func (c *Camera) GetRay(u, v float64) Ray {
	r1 := c.Horizontal.Scale(u)
	r2 := c.Vertical.Scale(v)
	d := c.LowerLeftCorner.Add(r1).Add(r2).Sub(c.Origin)
	return Ray{c.Origin, d}
}
