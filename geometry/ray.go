package geometry

type Ray struct {
	Origin    Vec3
	Direction Vec3
}

func (r *Ray) PointAtParameter(t float64) Vec3 {
	return r.Origin.Add(
		r.Direction.MultiplyScalar(t),
	)
}
