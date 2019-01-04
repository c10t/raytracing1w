package main

import (
	"math"
	"math/rand"
)

type Camera struct {
	LowerLeftCorner Vec3
	Horizontal      Vec3
	Vertical        Vec3
	Origin          Vec3
	LensRadius      float64
	ViewDirection
}

// orthonormal basis to describe camera's orientation (look from -> look at)
type ViewDirection struct {
	U Vec3
	V Vec3
	W Vec3
}

func NewVerticalCamera(lookFrom, lookAt, vup Vec3, vfov, aspect, aperture, focusDist float64) *Camera {
	lensRadius := aperture / 2

	theta := vfov * math.Pi / 180
	halfHeight := math.Tan(theta / 2)
	halfWidth := aspect * halfHeight

	orig := lookFrom
	w := lookFrom.Sub(lookAt).UnitVector()
	u := Cross(vup, w).UnitVector()
	v := Cross(w, u)

	x := u.Scale(focusDist).Scale(halfWidth)
	y := v.Scale(focusDist).Scale(halfHeight)
	z := w.Scale(focusDist)

	llc := orig.Sub(x).Sub(y).Sub(z)
	hor := x.Scale(2)
	ver := y.Scale(2)

	viewDirection := ViewDirection{U: u, V: v, W: w}

	return &Camera{
		LowerLeftCorner: llc,
		Horizontal:      hor,
		Vertical:        ver,
		Origin:          orig,
		LensRadius:      lensRadius,
		ViewDirection:   viewDirection,
	}
}

func (c *Camera) GetRay(s, t float64) Ray {
	rd := RandomInUnitDisk().Scale(c.LensRadius)
	offset := Add(c.ViewDirection.U.Scale(rd.X), c.ViewDirection.V.Scale(rd.Y))
	r1 := c.Horizontal.Scale(s)
	r2 := c.Vertical.Scale(t)
	d := c.LowerLeftCorner.Add(r1).Add(r2).Sub(c.Origin).Sub(offset)
	return Ray{c.Origin.Add(offset), d}
}

func RandomInUnitDisk() Vec3 {
	var p = Vec3{1, 1, 0}
	for p.SquaredLength() >= 1.0 {
		p.X = 2*rand.Float64() - 1
		p.Y = 2*rand.Float64() - 1
	}

	return p
}
