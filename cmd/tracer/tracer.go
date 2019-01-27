package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"math/rand"
	"os"
	"time"

	"github.com/alisdairrankine/turner"
)

func main() {
	rand.Seed(time.Now().Unix())

	width := 1024
	height := 768
	samples := 100

	photonCount := width * height * samples

	img := image.NewRGBA(image.Rect(
		0, 0, width, height,
	))

	lookfrom := turner.Vec3{13, 2, 3}
	lookat := turner.Vec3{0, 0, 0}
	camera := turner.NewCamera(
		lookfrom,
		lookat,
		turner.Vec3{0, 1, 0},
		20,
		float64(width)/float64(height),
		0.2,
		10,
	)
	world := bigWorld()

	fmt.Printf("starting raytrace: %d Photons ready\n", photonCount)
	i := 0
	fmt.Println()
	for y := height - 1; y >= 0; y-- {
		for x := 0; x < width; x++ {
			if i%1000 == 0 {
				fmt.Printf("%d/%d\r", i, photonCount)
			}
			colour := colourFromPixel(x, y, width, height, world, camera, samples)
			img.Set(x, y, colour)
			i += samples
		}
	}
	fmt.Printf("%d/%d\n", i, width*height)
	fmt.Println("Saving file")

	file, _ := os.Create("image.png")

	png.Encode(file, img)
}

func colourFromPixel(x, y, width, height int, world turner.Hitable, camera *turner.Camera, samples int) color.Color {
	colour := turner.Vec3{0, 0, 0}
	for s := 0; s < samples; s++ {
		u := (float64(x) + rand.Float64()) / float64(width)
		v := (float64(height-y) + rand.Float64()) / float64(height)
		ray := camera.Ray(u, v)
		colour = colour.Add(colourFromRay(ray, world, 0))
	}
	return colour.DivideScalar(float64(samples)).Colour()

}

func colourFromRay(ray turner.Ray, world turner.Hitable, depth int) turner.Vec3 {
	if hitRec, hit := world.Hit(&ray, 0.01, math.MaxFloat64); hit {
		if depth < 50 {
			if hit, attenuation, scatteredRay := hitRec.Mat.Scatter(ray, hitRec); hit {
				return attenuation.MultiplyVector(colourFromRay(scatteredRay, world, depth+1))
			}
		}
		return turner.Vec3{}
	}
	unitDir := ray.Direction.UnitVector()
	t := 0.5 * (unitDir.Y + 1)
	return turner.Vec3{1, 1, 1}.MultiplyScalar(1.0 - t).Add(turner.Vec3{0.5, 0.7, 1.0}.MultiplyScalar(t))

}

func newWorld() turner.Hitable {
	return turner.NewHitList(
		turner.Sphere{turner.Vec3{0, 0, -1}, 0.5, turner.Lambertian{turner.Vec3{0.8, 0.3, 0.3}}},
		turner.Sphere{turner.Vec3{0, -100.5, -1}, 100, turner.Lambertian{turner.Vec3{0.8, 0.8, 0}}},
		turner.Sphere{turner.Vec3{1, 0, -1}, 0.5, turner.Metal{turner.Vec3{0.8, 0.6, 0.2}, 0.3}},
		turner.Sphere{turner.Vec3{-1, 0, -1}, 0.5, turner.Dielectric{1.5}},
		turner.Sphere{turner.Vec3{-1, 0, -1}, -0.45, turner.Dielectric{1.5}},
	)
}

func bigWorld() turner.Hitable {
	scene := []turner.Hitable{}
	scene = append(scene, turner.Sphere{turner.Vec3{0, -1000, 0}, 1000, turner.Lambertian{turner.Vec3{0.5, 0.5, 0.5}}})
	for a := -11; a < 11; a++ {
		for b := -11; b < 11; b++ {
			chooseMat := rand.Float64()
			centre := turner.Vec3{float64(a) + 0.9*rand.Float64(), 0.2, float64(b) + 0.9*rand.Float64()}
			var mat turner.Material
			if centre.Minus(turner.Vec3{4, 0.2, 0}).Length() > 0.9 {
				if chooseMat < 0.8 {
					mat = turner.Lambertian{turner.Vec3{rand.Float64() * rand.Float64(), rand.Float64() * rand.Float64(), rand.Float64() * rand.Float64()}}
				} else if chooseMat < 0.95 {
					mat = turner.Metal{turner.Vec3{0.5 * (1 + rand.Float64()), 0.5 * (1 + rand.Float64()), 0.5 * (1 + rand.Float64())}, 0.5 * rand.Float64()}
				} else {
					mat = turner.Dielectric{1.5}
				}
				scene = append(scene,
					turner.Sphere{centre, 0.2, mat},
				)
			}
		}
	}
	scene = append(scene, turner.Sphere{turner.Vec3{0, 1, 0}, 1, turner.Dielectric{1.5}},
		turner.Sphere{turner.Vec3{-4, 1, 0}, 1, turner.Lambertian{turner.Vec3{0.4, 0.2, 0.1}}},
		turner.Sphere{turner.Vec3{4, 1, 0}, 1, turner.Metal{turner.Vec3{0.7, 0.6, 0.5}, 0}})
	return turner.NewHitList(scene...)
}

func Render() {

}
