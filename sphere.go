package turner

import "math"

type Sphere struct {
	Centre Vec3
	Radius float64
}

func (s Sphere) Hit(ray *Ray, tMin, tMax float64) (*HitRecord, bool) {
	oc := ray.Origin.Minus(s.Centre)
	a := ray.Direction.Dot(ray.Direction)
	b := oc.Dot(ray.Direction)
	c := oc.Dot(oc) - (s.Radius * s.Radius)
	discriminant := (b * b) - (a * c)
	rec := &HitRecord{}
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
