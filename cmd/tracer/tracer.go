package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/alisdairrankine/turner/geometry"
)

func main() {
	width := 200
	height := 100

	img := image.NewRGBA(image.Rect(
		0, 0, width, height,
	))

	lowerLeft := geometry.Vec3{-2.0, -1.0, -1.0}
	horizontal := geometry.Vec3{4.0, 0.0, 0.0}
	vertical := geometry.Vec3{0.0, 2.0, 0.0}
	origin := geometry.Vec3{0.0, 0.0, 0.0}

	fmt.Println("starting raytrace")
	for y := height - 1; y >= 0; y-- {

		for x := 0; x < width; x++ {
			u := float64(x) / float64(width)
			v := float64(height-y) / float64(height)
			ray := geometry.Ray{
				origin,
				lowerLeft.Add(horizontal.MultiplyScalar(u)).Add(vertical.MultiplyScalar(v)),
			}
			img.Set(x, y, colourFromRay(ray))

		}
	}
	fmt.Println("Saving file")

	file, _ := os.Create("image.png")

	png.Encode(file, img)
}

func colourFromRay(ray geometry.Ray) color.Color {
	t := hitSphere(geometry.Vec3{0, 0, -1}, 0.5, ray)
	if t > 0 {
		N := ray.PointAtParameter(t).Minus(geometry.Vec3{0, 0, -1}).UnitVector()
		return geometry.Vec3{N.X + 1, N.Y + 1, N.Z + 1}.MultiplyScalar(0.5).Colour()
	}

	unitDir := ray.Direction.UnitVector()
	t = 0.5 * (unitDir.Y + 1)
	vec1 := geometry.Vec3{1, 1, 1}
	vec2 := geometry.Vec3{0.5, 0.7, 1.0}
	colour := vec1.MultiplyScalar(1.0 - t).Add(vec2.MultiplyScalar(t))
	return colour.Colour()
}

func hitSphere(center geometry.Vec3, radius float64, ray geometry.Ray) float64 {
	oc := ray.Origin.Minus(center)
	a := ray.Direction.Dot(ray.Direction)
	b := oc.Dot(ray.Direction) * 2.0
	c := oc.Dot(oc) - (radius * radius)
	discriminant := (b * b) - (4 * a * c)
	if discriminant < 0 {
		return -1
	} else {
		return (-b - discriminant) / (2 * a)
	}
}
