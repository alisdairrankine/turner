package main

import (
	"fmt"
	"image"
	"image/png"
	"math"
	"math/rand"
	"os"
	"time"

	"github.com/alisdairrankine/turner"
)

func main() {
	width := 500
	height := 250
	samples := 100
	rand.Seed(time.Now().Unix())

	img := image.NewRGBA(image.Rect(
		0, 0, width, height,
	))

	camera := turner.DefaultCamera()
	world := newWorld()

	fmt.Println("starting raytrace")
	for y := height - 1; y >= 0; y-- {
		for x := 0; x < width; x++ {
			colour := turner.Vec3{0, 0, 0}
			for s := 0; s < samples; s++ {
				u := (float64(x) + rand.Float64()) / float64(width)
				v := (float64(height-y) + rand.Float64()) / float64(height)
				ray := camera.Ray(u, v)

				colour = colour.Add(colourFromRay(ray, world))

			}
			img.Set(x, y, colour.DivideScalar(float64(samples)).Colour())

		}
	}
	fmt.Println("Saving file")

	file, _ := os.Create("image.png")

	png.Encode(file, img)
}

func colourFromRay(ray turner.Ray, world turner.Hitable) turner.Vec3 {
	if hitRec, hit := world.Hit(&ray, 0, math.MaxFloat64); hit {
		return turner.Vec3{hitRec.Normal.X + 1, hitRec.Normal.Y + 1, hitRec.Normal.Z + 1}.MultiplyScalar(0.5)
	}

	unitDir := ray.Direction.UnitVector()
	t := 0.5 * (unitDir.Y + 1)
	vec1 := turner.Vec3{1, 1, 1}
	vec2 := turner.Vec3{0.5, 0.7, 1.0}
	colour := vec1.MultiplyScalar(1.0 - t).Add(vec2.MultiplyScalar(t))
	return colour
}

func newWorld() turner.Hitable {
	return turner.NewHitList(
		turner.Sphere{turner.Vec3{0, -100.5, -1}, 100},
		turner.Sphere{turner.Vec3{0, 0, -1}, 0.5},
	)
}
