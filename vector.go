package turner

import (
	"image/color"
	"math"
)

type Vec3 struct {
	X, Y, Z float64
}

func (v Vec3) Add(vec Vec3) Vec3 {
	return Vec3{v.X + vec.X, v.Y + vec.Y, v.Z + vec.Z}
}

func (v Vec3) Minus(vec Vec3) Vec3 {
	return Vec3{v.X - vec.X, v.Y - vec.Y, v.Z - vec.Z}
}

func (v Vec3) MultiplyVector(vec Vec3) Vec3 {
	return Vec3{v.X * vec.X, v.Y * vec.Y, v.Z * vec.Z}
}

func (v Vec3) MultiplyScalar(m float64) Vec3 {
	return Vec3{v.X * m, v.Y * m, v.Z * m}
}

func (v Vec3) DivideVector(vec Vec3) Vec3 {
	return Vec3{v.X / vec.X, v.Y / vec.Y, v.Z / vec.Z}
}

func (v Vec3) DivideScalar(m float64) Vec3 {
	return Vec3{v.X / m, v.Y / m, v.Z / m}
}

func (v Vec3) Dot(vec Vec3) float64 {
	return v.X*vec.X + v.Y*vec.Y + v.Z*vec.Z
}

func (v Vec3) Cross(vec Vec3) Vec3 {
	return Vec3{
		(v.Y * vec.Z) - (v.Z * vec.Y),
		(v.Z * vec.X) - (v.X * vec.Z),
		(v.X * vec.Y) - (v.Y * vec.X),
	}
}

func (v Vec3) Length() float64 {
	return math.Sqrt(v.SquaredLength())
}

func (v Vec3) SquaredLength() float64 {
	return (v.X * v.X) + (v.Y * v.Y) + (v.Z * v.Z)
}

func (v Vec3) UnitVector() Vec3 {
	return v.DivideScalar(v.Length())
}

func (v Vec3) Colour() color.Color {
	return color.RGBA{
		uint8(255.99 * v.X),
		uint8(255.99 * v.Y),
		uint8(255.99 * v.Z),
		255,
	}
}

func (v Vec3) Reflect(normal Vec3) Vec3 {
	return v.Minus(normal.MultiplyScalar(v.Dot(normal) * 2))
}

func (v Vec3) Refract(normal Vec3, niOverNt float64) (bool, Vec3) {
	uv := v.UnitVector()
	refracted := Vec3{}
	dt := uv.Dot(normal)
	discriminant := 1.0 - niOverNt*niOverNt*(1-dt*dt)
	if discriminant > 0 {
		refracted = uv.Minus(normal.MultiplyScalar(dt)).MultiplyScalar(niOverNt).Minus(normal.MultiplyScalar(math.Sqrt(discriminant)))
		return true, refracted
	}
	return false, refracted
}
