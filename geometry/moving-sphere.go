package geometry

import (
	"math"

	"github.com/alisdairrankine/turner"
)

type MovingSphere struct {
	StartCentre turner.Vec3
	EndCentre   turner.Vec3
	Radius      float64
	Mat         turner.Material
	StartTime   float64
	EndTime     float64
}

func (s MovingSphere) centre(time float64) turner.Vec3 {
	return s.StartCentre.Add(s.EndCentre.Minus(s.StartCentre)).MultiplyScalar((time - s.StartTime) / (s.EndTime - s.StartTime))
}

func (s MovingSphere) Hit(ray *turner.Ray, tMin, tMax float64) (*turner.HitRecord, bool) {
	oc := ray.Origin.Minus(s.centre(ray.Time))
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
			rec.Normal = rec.P.Minus(s.centre(ray.Time)).DivideScalar(s.Radius)
			return rec, true
		}
		temp = (-b + math.Sqrt(b*b-a*c)) / a
		if temp < tMax && temp > tMin {
			rec.T = temp
			rec.P = ray.PointAtParameter(temp)
			rec.Normal = rec.P.Minus(s.centre(ray.Time)).DivideScalar(s.Radius)
			return rec, true
		}
	}

	return rec, false
}
