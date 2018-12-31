package main

type HitRecord struct {
	At            float64
	Point, Normal Vec3
	Material
}

type Hitable interface {
	Hit(r Ray, tmin float64, tmax float64) (bool, HitRecord)
}

type World []Hitable

func (w *World) Hit(r Ray, tmin float64, tmax float64) (bool, HitRecord) {
	hitAnything := false
	closestSoFar := tmax
	finalRecord := HitRecord{}

	for _, item := range *w {
		if item != nil {
			hit, rec := item.Hit(r, tmin, closestSoFar)

			if hit {
				hitAnything = true
				closestSoFar = rec.At
				finalRecord = rec
			}
		}
	}

	return hitAnything, finalRecord
}
