package main

type HitRecord struct {
	At            float64
	Point, Normal Vec3
}

type Hitable interface {
	Hit(r Ray, tmin float64, tmax float64) (bool, HitRecord)
}
