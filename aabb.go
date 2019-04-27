package turner

type AABB struct {
	Min Vec3
	Max Vec3
}

func FMin(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func FMax(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func (h AABB) Hit(ray *Ray, tMin, tMax float64) (*HitRecord, bool) {

	if !hitDimension(h.Min.X, h.Max.X, ray.Origin.X, ray.Direction.X, tMin, tMax) ||
		!hitDimension(h.Min.Y, h.Max.Y, ray.Origin.Y, ray.Direction.Y, tMin, tMax) ||
		!hitDimension(h.Min.Z, h.Max.Z, ray.Origin.Z, ray.Direction.Z, tMin, tMax) {
		return nil, false
	}
	return nil, true
}

func hitDimension(min, max, origin, direction, tMin, tMax float64) bool {
	t0 := FMin(
		(min-origin)/direction,
		(max-origin)/direction,
	)
	t1 := FMax(
		(min-origin)/direction,
		(max-origin)/direction,
	)
	min2 := FMax(t0, tMin)
	max2 := FMin(t1, tMax)
	if max2 <= min2 {
		return false
	}
	return true
}

func SurroundingBox(a, b *AABB) *AABB {

	small := Vec3{
		FMin(a.Min.X, b.Min.X),
		FMin(a.Min.Y, b.Min.Y),
		FMin(a.Min.Z, b.Min.Z),
	}

	big := Vec3{
		FMin(a.Max.X, b.Max.X),
		FMin(a.Max.Y, b.Max.Y),
		FMin(a.Max.Z, b.Max.Z),
	}

	return &AABB{small, big}
}
