package turner

type Material interface {
	Scatter(rayIn Ray, hitRec *HitRecord) (hit bool, Attenuation Vec3, scatteredRay Ray)
}
