package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"time"
)

func main() {
	const timeLayout = "20060102-1504-05"
	t := time.Now()
	filename := t.Format(timeLayout) + ".ppm"
	log.Println(fmt.Sprintf("Writing to %v ...", filename))

	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	for _, line := range lerp(200, 100, 100, time.Now().UnixNano()) {
		writer.WriteString(line + "\n")
	}
	writer.Flush()
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

func lerp(nx, ny, ns int, seed int64) []string {
	rand.Seed(seed)

	result := []string{"P3", fmt.Sprintf("%d %d", nx, ny), "255"}

	s1 := NewSphere(0, 0, -1, 0.5)
	s2 := NewSphere(0, -100.5, -1, 100)
	world := World{s1, s2}

	cam := NewCamera()

	for j := ny - 1; j > -1; j-- {
		for i := 0; i < nx; i++ {
			col := Vec3{0, 0, 0}
			for s := 0; s < ns; s++ {
				u := (float64(i) + rand.Float64()) / float64(nx)
				v := (float64(j) + rand.Float64()) / float64(ny)
				r := cam.GetRay(u, v)
				col = col.Add(color(&r, &world))
			}
			col = col.Shrink(float64(ns))

			ir := int(255.99 * col.X)
			ig := int(255.99 * col.Y)
			ib := int(255.99 * col.Z)
			result = append(result, fmt.Sprintf("%d %d %d", ir, ig, ib))
		}
	}

	return result
}
