package main

import (
	"math/rand"

	"github.com/alisdairrankine/turner"
)

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
