package geometry

type Hitable interface {
	Hit(r *Ray, tMin, tMax float64) (*HitRecord, bool)
}

type HitRecord struct {
	T      float64
	P      Vec3
	Normal Vec3
}
