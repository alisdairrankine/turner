package turner

type Camera struct {
	LowerLeftCorner Vec3
	Horizontal      Vec3
	Vertical        Vec3
	Origin          Vec3
}

func DefaultCamera() *Camera {

	lowerLeft := Vec3{-2.0, -1.0, -1.0}
	horizontal := Vec3{4.0, 0.0, 0.0}
	vertical := Vec3{0.0, 2.0, 0.0}
	origin := Vec3{0.0, 0.0, 0.0}
	return &Camera{
		lowerLeft,
		horizontal,
		vertical,
		origin,
	}
}

func (c *Camera) Ray(u, v float64) Ray {
	return Ray{
		c.Origin,
		c.LowerLeftCorner.Add(c.Horizontal.MultiplyScalar(u)).Add(c.Vertical.MultiplyScalar(v)),
	}
}
