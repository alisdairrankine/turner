package turner

import (
	"math"
)

type Camera struct {
	LowerLeft  Vec3
	Horizontal Vec3
	Vertical   Vec3
	Origin     Vec3
	U          Vec3
	V          Vec3
	W          Vec3
	LensRadius float64
}

func NewCamera(lookfrom, lookat, vUp Vec3, vfov, aspect, aperture, focusDistance float64) Camera {
	lensRadius := aperture / 2
	theta := vfov * math.Pi / 180
	halfHeight := math.Tan(theta / 2)
	halfWidth := aspect * halfHeight
	origin := lookfrom
	w := lookfrom.Minus(lookat).UnitVector()
	u := vUp.Cross(w).UnitVector()
	v := w.Cross(u)
	lowerLeft := origin.Minus(u.MultiplyScalar(halfWidth * focusDistance)).Minus(v.MultiplyScalar(halfHeight * focusDistance)).Minus(w.MultiplyScalar(focusDistance))
	horizontal := u.MultiplyScalar(2 * halfWidth * focusDistance)
	vertical := v.MultiplyScalar(2 * halfHeight * focusDistance)
	return Camera{
		LowerLeft:  lowerLeft,
		Horizontal: horizontal,
		Vertical:   vertical,
		Origin:     origin,
		U:          u,
		V:          v,
		W:          w,
		LensRadius: lensRadius,
	}
}

func (c *Camera) Ray(u, v float64) Ray {
	rd := RandomPointInUnitDisc().MultiplyScalar(c.LensRadius)
	offset := c.U.MultiplyScalar(rd.X).Add(c.V.MultiplyScalar(rd.Y))
	return Ray{
		c.Origin.Add(offset),
		c.LowerLeft.Add(c.Horizontal.MultiplyScalar(u)).Add(c.Vertical.MultiplyScalar(v)).Minus(c.Origin).Minus(offset),
	}
}
