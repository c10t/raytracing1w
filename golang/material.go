package main

import "math"

type Material interface {
	Scatter(r Ray, rec HitRecord) (bool, Vec3, Ray)
}

type Lambertian struct {
	Albedo Vec3
}

func (l Lambertian) Scatter(r Ray, rec HitRecord) (bool, Vec3, Ray) {
	target := Add(rec.Point, rec.Normal).Add(RandomInUnitSphere())
	scattered := Ray{Origin: rec.Point, Direction: target.Sub(rec.Point)}
	attenuation := l.Albedo
	return true, attenuation, scattered
}

type Metal struct {
	Albedo Vec3
	fuzz   float64
}

func (m Metal) Scatter(r Ray, rec HitRecord) (bool, Vec3, Ray) {
	reflected := reflect(r.Direction.UnitVector(), rec.Normal)
	perturbation := RandomInUnitSphere().Scale(m.fuzz)
	scattered := Ray{Origin: rec.Point, Direction: reflected.Add(perturbation)}
	attenuation := m.Albedo
	return (Dot(scattered.Direction, rec.Normal) > 0), attenuation, scattered
}

func reflect(v Vec3, n Vec3) Vec3 {
	return v.Sub(n.Scale(Dot(v, n) * 2))
}

type Dielectric struct {
	refractiveIndex float64
}

func (d Dielectric) Scatter(r Ray, rec HitRecord) (bool, Vec3, Ray) {
	reflected := reflect(r.Direction, rec.Normal)
	var scattered Ray
	var refracted Vec3
	var outwardNormal Vec3
	var ratioNiOverNt float64
	// 1. the glass surface absorbs nothing
	// 2. kill the blue channel which is the type of color bug
	//    that can be hard to find
	attenuation := Vec3{1, 1, 0}

	if Dot(r.Direction, rec.Normal) > 0 {
		outwardNormal = rec.Normal.Scale(-1)
		ratioNiOverNt = d.refractiveIndex
	} else {
		outwardNormal = rec.Normal
		ratioNiOverNt = 1.0 / d.refractiveIndex
	}

	isRefracted, refracted := refract(r.Direction, outwardNormal, ratioNiOverNt)
	if isRefracted {
		scattered = Ray{rec.Point, refracted}
		return true, attenuation, scattered
	} else {
		scattered = Ray{rec.Point, reflected}
		return false, attenuation, scattered
	}
}

func refract(v Vec3, n Vec3, ratio float64) (bool, Vec3) {
	uv := v.UnitVector()
	dt := Dot(uv, n)
	discriminant := 1.0 - ratio*ratio*(1-dt*dt)

	if discriminant > 0 {
		refracted := uv.Sub(n.Scale(dt)).Scale(ratio).Sub(n.Scale(math.Sqrt(discriminant)))
		return true, refracted
	} else {
		return false, Vec3{}
	}
}
