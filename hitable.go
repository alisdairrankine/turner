package turner

type Hitable interface {
	Hit(r *Ray, tMin, tMax float64) (*HitRecord, bool)
	BoundingBox(t0, t1 float64) (*AABB, bool)
}

type HitRecord struct {
	T      float64
	P      Vec3
	Normal Vec3
	Mat    Material
}
