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

func NewVerticalCamera(lookFrom, lookAt, vup Vec3, vfov, aspect float64) *Camera {
	theta := vfov * math.Pi / 180
	halfHeight := math.Tan(theta / 2)
	halfWidth := aspect * halfHeight

	orig := lookFrom
	w := lookFrom.Sub(lookAt).UnitVector()
	u := Cross(vup, w).UnitVector()
	v := Cross(w, u)

	x := u.Scale(halfWidth)
	y := v.Scale(halfHeight)

	llc := orig.Sub(x).Sub(y).Sub(w)
	hor := x.Scale(2)
	ver := y.Scale(2)

	return &Camera{LowerLeftCorner: llc, Horizontal: hor, Vertical: ver, Origin: orig}
}

func (c *Camera) GetRay(u, v float64) Ray {
	r1 := c.Horizontal.Scale(u)
	r2 := c.Vertical.Scale(v)
	d := c.LowerLeftCorner.Add(r1).Add(r2).Sub(c.Origin)
	return Ray{c.Origin, d}
}
