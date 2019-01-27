package main

import (
	"math/rand"
	"time"

	"github.com/alisdairrankine/turner"
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
	)

	renderer := turner.NewRenderer(
		bigWorld(),
		camera,
		width,
		height,
		"image.png",
		samples,
		bounces,
	)
	renderer.Render()
}
