package turner

type Scene struct {
	Hitables []Hitable
}

func NewBVHScene(hitables ...Hitable) Hitable {
	return BVHNode(hitables, 0, 0)
}

func NewScene(hitables ...Hitable) Scene {
	return Scene{
		Hitables: hitables,
	}
}

func (h Scene) Hit(ray *Ray, tMin, tMax float64) (*HitRecord, bool) {
	var rec *HitRecord
	hit := false
	closest := tMax
	for _, hitable := range h.Hitables {
		if tempRec, isHit := hitable.Hit(ray, tMin, closest); isHit {
			rec = tempRec
			closest = tempRec.T
			hit = true
		}

	}
	return rec, hit
}

func (h Scene) BoundingBox(t0, t1 float64) (*AABB, bool) {

	if len(h.Hitables) < 1 {
		return nil, false
	}
	box := &AABB{}
	if b, ok := h.Hitables[0].BoundingBox(t0, t1); ok {
		box = b
	} else {
		return nil, false
	}
	return box, false
	for i, hitable := range h.Hitables {
		if i == 0 {
			continue
		}
		if b, ok := hitable.BoundingBox(t0, t1); ok {
			box = SurroundingBox(box, b)
		} else {
			return nil, false
		}
	}
	return box, true
}
