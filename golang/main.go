package main

import (
	"fmt"
	"log"
	"math"
	"time"
)

func main() {
	const timeLayout = "20060102-1504-05"
	t := time.Now()
	filename := t.Format(timeLayout)
	log.Println(fmt.Sprintf("Writing to %v.ppm ...", filename))

	// for _, line := range lerp(200, 100) {
	//   fmt.Println(line)
	// }
}

func color(r *Ray, w *World) Vec3 {
	hit, rec := w.Hit(*r, 0, math.MaxFloat64)
	if hit {
		return rec.Normal.Slide(1).Scale(0.5)
	}

	unitDirection := r.Direction.UnitVector()
	t := 0.5 * (unitDirection.Y + 1)
	v1 := Vec3{1, 1, 1}
	v2 := Vec3{0.5, 0.7, 1}
	return Add(v1.Scale(1-t), v2.Scale(t))
}

func lerp(nx, ny int) []string {
	result := []string{"P3", fmt.Sprintf("%d %d", nx, ny), "255"}

	lowerLeftCorner := Vec3{-2, -1, -1}
	horizontal := Vec3{4, 0, 0}
	vertical := Vec3{0, 2, 0}
	origin := Vec3{0, 0, 0}

	s1 := NewSphere(0, 0, -1, 0.5)
	s2 := NewSphere(0, -100.5, -1, 100)
	world := World{s1, s2}

	for j := ny - 1; j > -1; j-- {
		for i := 0; i < nx; i++ {
			u := float64(i) / float64(nx)
			v := float64(j) / float64(ny)
			d := lowerLeftCorner.Add(horizontal.Scale(u)).Add(vertical.Scale(v))
			r := Ray{Origin: origin, Direction: d}
			c := color(&r, &world)
			ir := int(255.99 * c.X)
			ig := int(255.99 * c.Y)
			ib := int(255.99 * c.Z)
			result = append(result, fmt.Sprintf("%d %d %d", ir, ig, ib))
		}
	}

	return result
}
