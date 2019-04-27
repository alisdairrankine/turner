package main

import (
	"math/rand"
	"time"

	"github.com/alisdairrankine/turner"

	"github.com/alisdairrankine/turner/geometry"
	"github.com/alisdairrankine/turner/material"
)

func main() {
	rand.Seed(time.Now().Unix())

	width := 800
	height := 600
	samples := 100
	bounces := 50

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
		0,
		1,
	)

	renderer := turner.NewRenderer(
		bigWorld(),
		camera,
		width,
		height,
		"render.png",
		samples,
		bounces,
	)
	renderer.Render()
}

func bigWorld() turner.Hitable {
	scene := []turner.Hitable{}
	scene = append(scene, geometry.Sphere{turner.Vec3{X: 0, Y: -1000, Z: 0}, 1000, material.Diffuse(turner.Vec3{X: 0.5, Y: 0.5, Z: 0.5})})
	for a := -11; a < 11; a++ {
		for b := -11; b < 11; b++ {
			chooseMat := rand.Float64()
			centre := turner.Vec3{
				X: float64(a) + 0.9*rand.Float64(),
				Y: 0.2,
				Z: float64(b) + 0.9*rand.Float64(),
			}
			var mat turner.Material
			if centre.Minus(turner.Vec3{X: 4, Y: 0.2, Z: 0}).Length() > 0.9 {
				if chooseMat < 0.8 {
					mat = material.Diffuse(
						turner.Vec3{
							X: rand.Float64() * rand.Float64(),
							Y: rand.Float64() * rand.Float64(),
							Z: rand.Float64() * rand.Float64()},
					)

				} else if chooseMat < 0.95 {
					mat = material.Metal(
						turner.Vec3{
							X: 0.5 * (1 + rand.Float64()),
							Y: 0.5 * (1 + rand.Float64()),
							Z: 0.5 * (1 + rand.Float64())},
						0.5*rand.Float64(),
					)

				} else {
					mat = material.Dielectric(1.5)

				}

				scene = append(scene,
					geometry.Sphere{centre, 0.2, mat},
				)
			}
		}
	}
	scene = append(scene, geometry.Sphere{turner.Vec3{0, 1, 0}, 1, material.Dielectric(1.5)},
		geometry.Sphere{turner.Vec3{-4, 1, 0}, 1, material.Metal(turner.Vec3{0.7, 0.6, 0.5}, 0)},
		geometry.Sphere{turner.Vec3{4, 1, 0}, 1, material.Diffuse(turner.Vec3{0.4, 0.2, 0.1})})
	return turner.NewScene(scene...)
}
