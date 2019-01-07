package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"sync"
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

	rand.Seed(time.Now().UnixNano())

	ppm := lerp(200, 100, 100)

	writeStart := time.Now()
	for _, line := range ppm {
		writer.WriteString(line + "\n")
	}
	writer.Flush()

	end := time.Now()
	log.Println(fmt.Sprintf("Write: %v sec", end.Sub(writeStart).Seconds()))
	log.Println(fmt.Sprintf("Total: %v sec", end.Sub(t).Seconds()))
}

func RandomInUnitSphere() Vec3 {
	var p = Vec3{1, 1, 1}
	for p.SquaredLength() >= 1.0 {
		p.X = 2*rand.Float64() - 1
		p.Y = 2*rand.Float64() - 1
		p.Z = 2*rand.Float64() - 1
	}

	return p
}

func color(r *Ray, w *World, depth int) Vec3 {
	hit, rec := w.Hit(*r, 0.001, math.MaxFloat64)
	if hit {
		isScattered, attenuation, scatteredRay := rec.Material.Scatter(*r, rec)
		if depth < 50 && isScattered {
			return color(&scatteredRay, w, depth+1).Mul(attenuation)
		} else {
			return Vec3{0, 0, 0}
		}
	}

	unitDirection := r.Direction.UnitVector()
	t := 0.5 * (unitDirection.Y + 1)
	v1 := Vec3{1, 1, 1}
	v2 := Vec3{0.5, 0.7, 1}
	return Add(v1.Scale(1-t), v2.Scale(t))
}

func lerp(nx, ny, ns int) []string {
	result := make([]string, nx*ny+3)
	result[0] = "P3"
	result[1] = fmt.Sprintf("%d %d", nx, ny)
	result[2] = "255"

	/*
		s1 := NewSphere(0, 0, -1, 0.5, Lambertian{Albedo: Vec3{0.1, 0.2, 0.5}})
		s2 := NewSphere(0, -100.5, -1, 100, Lambertian{Albedo: Vec3{0.8, 0.8, 0.0}})
		s3 := NewSphere(1, 0, -1, 0.5, NewMetal(Vec3{0.8, 0.6, 0.2}, 0.1))
		s4 := NewSphere(-1, 0, -1, 0.5, Dielectric{refractiveIndex: 1.5})
		s5 := NewSphere(-1, 0, -1, -0.45, Dielectric{refractiveIndex: 1.5})
		world := World{s1, s2, s3, s4, s5}
	*/

	world := makeTheWorld()
	lookF := Vec3{3, 1, 2}
	lookA := Vec3{0, 0, -1}
	distToFocus := Sub(lookF, lookA).Length()
	aperture := 0.1

	vup := Vec3{0, 1, 0}
	aspect := float64(nx) / float64(ny)
	cam := NewVerticalCamera(lookF, lookA, vup, 90, aspect, aperture, distToFocus)

	startLoop := time.Now()

	var wg sync.WaitGroup

	for j := ny - 1; j > -1; j-- {
		for i := 0; i < nx; i++ {
			wg.Add(1)
			go func(i, j int) {
				defer wg.Done()
				colorAtPoint(result, nx, ny, ns, i, j, cam, &world)
			}(i, j)
		}
	}

	wg.Wait()
	endLoop := time.Now()
	timeForLoop := endLoop.Sub(startLoop).Seconds()
	log.Println(fmt.Sprintf("lerp loop takes %v second", timeForLoop))

	return result
}

func colorAtPoint(ppm []string, nx, ny, ns, i, j int, cam *Camera, w *World) {
	col := Vec3{0, 0, 0}
	for s := 0; s < ns; s++ {
		u := (float64(i) + rand.Float64()) / float64(nx)
		v := (float64(j) + rand.Float64()) / float64(ny)
		r := cam.GetRay(u, v)
		col = col.Add(color(&r, w, 0))
	}
	col = col.Shrink(float64(ns))

	ir := int(255.99 * math.Sqrt(col.X))
	ig := int(255.99 * math.Sqrt(col.Y))
	ib := int(255.99 * math.Sqrt(col.Z))

	lineIndex := nx*(ny-1-j) + i + 3
	ppm[lineIndex] = fmt.Sprintf("%d %d %d", ir, ig, ib)
}

func makeTheWorld() World {
	spheres := make([]Hitable, 500)

	spheres[0] = NewSphere(0, -1000, 0, 1000, Lambertian{Albedo: Vec3{0.5, 0.5, 0.5}})

	idx := 1
	for a := -11; a < 11; a++ {
		for b := -11; b < 11; b++ {
			center := Vec3{
				X: float64(a) + 0.9*rand.Float64(),
				Y: 0.2,
				Z: float64(b) + 0.9*rand.Float64(),
			}
			if center.Sub(Vec3{4, 0.2, 0}).Length() > 0.9 {
				m := createMaterial(rand.Float64())
				spheres[idx] = NewSphere(center.X, center.Y, center.Z, 0.2, m)
				idx++
			}
		}
	}

	spheres[idx] = NewSphere(0, 1, 0, 1.0, Dielectric{refractiveIndex: 1.5})
	idx++
	spheres[idx] = NewSphere(-4, 1, 0, 1.0, Lambertian{Albedo: Vec3{0.4, 0.2, 0.1}})
	idx++
	spheres[idx] = NewSphere(4, 1, 0, 1.0, Metal{Vec3{0.7, 0.6, 0.5}, 0.0})
	idx++

	return spheres
}

func createMaterial(factor float64) Material {
	switch {
	case factor < 0.8:
		albedo := Vec3{
			X: rand.Float64() * rand.Float64(),
			Y: rand.Float64() * rand.Float64(),
			Z: rand.Float64() * rand.Float64(),
		}
		return Lambertian{Albedo: albedo}
	case factor < 0.95:
		albedo := Vec3{
			X: 0.5 * (1 + rand.Float64()),
			Y: 0.5 * (1 + rand.Float64()),
			Z: 0.5 * (1 + rand.Float64()),
		}
		return Metal{Albedo: albedo, fuzz: 0.5 * rand.Float64()}
	default:
		return Dielectric{refractiveIndex: 1.5}
	}
}
