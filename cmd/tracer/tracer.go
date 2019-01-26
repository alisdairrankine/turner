package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"

	"github.com/alisdairrankine/turner/geometry"
)

func main() {
	width := 500
	height := 250

	img := image.NewRGBA(image.Rect(
		0, 0, width, height,
	))

	lowerLeft := geometry.Vec3{-2.0, -1.0, -1.0}
	horizontal := geometry.Vec3{4.0, 0.0, 0.0}
	vertical := geometry.Vec3{0.0, 2.0, 0.0}
	origin := geometry.Vec3{0.0, 0.0, 0.0}

	world := newWorld()

	fmt.Println("starting raytrace")
	for y := height - 1; y >= 0; y-- {
		for x := 0; x < width; x++ {
			u := float64(x) / float64(width)
			v := float64(height-y) / float64(height)
			ray := geometry.Ray{
				origin,
				lowerLeft.Add(horizontal.MultiplyScalar(u)).Add(vertical.MultiplyScalar(v)),
			}
			img.Set(x, y, colourFromRay(ray, world))

		}
	}
	fmt.Println("Saving file")

	file, _ := os.Create("image.png")

	png.Encode(file, img)
}

func colourFromRay(ray geometry.Ray, world geometry.Hitable) color.Color {
	if hitRec, hit := world.Hit(&ray, 0, math.MaxFloat64); hit {
		return geometry.Vec3{hitRec.Normal.X + 1, hitRec.Normal.Y + 1, hitRec.Normal.Z + 1}.MultiplyScalar(0.5).Colour()
	}

	unitDir := ray.Direction.UnitVector()
	t := 0.5 * (unitDir.Y + 1)
	vec1 := geometry.Vec3{1, 1, 1}
	vec2 := geometry.Vec3{0.5, 0.7, 1.0}
	colour := vec1.MultiplyScalar(1.0 - t).Add(vec2.MultiplyScalar(t))
	return colour.Colour()
}

func newWorld() geometry.Hitable {
	return geometry.NewHitList(
		geometry.Sphere{geometry.Vec3{0, -100.5, -1}, 100},
		geometry.Sphere{geometry.Vec3{0, 0, -1}, 0.5},
	)
}
