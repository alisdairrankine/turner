package main

import (
	"math/rand"

	"github.com/alisdairrankine/turner"
	"github.com/alisdairrankine/turner/material"
)

func bigWorld() turner.Hitable {
	scene := []turner.Hitable{}
	scene = append(scene, turner.Sphere{turner.Vec3{X: 0, Y: -1000, Z: 0}, 1000, material.Diffuse(turner.Vec3{X: 0.5, Y: 0.5, Z: 0.5})})
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
					turner.Sphere{centre, 0.2, mat},
				)
			}
		}
	}
	scene = append(scene, turner.Sphere{turner.Vec3{0, 1, 0}, 1, material.Dielectric(1.5)},
		turner.Sphere{turner.Vec3{-4, 1, 0}, 1, material.Diffuse(turner.Vec3{0.4, 0.2, 0.1})},
		turner.Sphere{turner.Vec3{4, 1, 0}, 1, material.Metal(turner.Vec3{0.7, 0.6, 0.5}, 0)})
	return turner.NewHitList(scene...)
}
