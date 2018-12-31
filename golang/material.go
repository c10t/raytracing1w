package main

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
}

func (m Metal) Scatter(r Ray, rec HitRecord) (bool, Vec3, Ray) {
	reflected := reflect(r.Direction.UnitVector(), rec.Normal)
	scattered := Ray{Origin: rec.Point, Direction: reflected}
	attenuation := m.Albedo
	return (Dot(scattered.Direction, rec.Normal) > 0), attenuation, scattered
}

func reflect(v Vec3, n Vec3) Vec3 {
	return v.Sub(n.Scale(Dot(v, n) * 2))
}
