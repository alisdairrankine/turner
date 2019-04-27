package geometry

import (
	"math"

	"github.com/alisdairrankine/turner"
)

type Sphere struct {
	Centre turner.Vec3
	Radius float64
	Mat    turner.Material
}

func (s Sphere) Hit(ray *turner.Ray, tMin, tMax float64) (*turner.HitRecord, bool) {
	oc := ray.Origin.Minus(s.Centre)
	a := ray.Direction.Dot(ray.Direction)
	b := oc.Dot(ray.Direction)
	c := oc.Dot(oc) - (s.Radius * s.Radius)
	discriminant := (b * b) - (a * c)
	rec := &turner.HitRecord{Mat: s.Mat}
	if discriminant > 0 {
		temp := (-b - math.Sqrt(b*b-a*c)) / a
		if temp < tMax && temp > tMin {
			rec.T = temp
			rec.P = ray.PointAtParameter(temp)
			rec.Normal = rec.P.Minus(s.Centre).DivideScalar(s.Radius)
			return rec, true
		}
		temp = (-b + math.Sqrt(b*b-a*c)) / a
		if temp < tMax && temp > tMin {
			rec.T = temp
			rec.P = ray.PointAtParameter(temp)
			rec.Normal = rec.P.Minus(s.Centre).DivideScalar(s.Radius)
			return rec, true
		}
	}

	return rec, false
}

func (s Sphere) BoundingBox(t0, t1 float64) (*turner.AABB, bool) {
	return &turner.AABB{Min: s.Centre.Minus(turner.Vec3{s.Radius, s.Radius, s.Radius}), Max: s.Centre.Add(turner.Vec3{s.Radius, s.Radius, s.Radius})}, true
}
