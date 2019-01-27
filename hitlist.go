package turner

type HitList struct {
	Hitables []Hitable
}

func NewHitList(hitables ...Hitable) HitList {
	return HitList{
		Hitables: hitables,
	}
}

func (h HitList) Hit(ray *Ray, tMin, tMax float64) (*HitRecord, bool) {
	var rec *HitRecord
	hit := false
	closest := tMax
	for _, hitable := range h.Hitables {
		if tempRec, isHit := hitable.Hit(ray, tMin, closest); isHit {
			rec = tempRec
			hit = true
		}

	}
	return rec, hit
}
