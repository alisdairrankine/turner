package turner

type Scene struct {
	Hitables []Hitable
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
