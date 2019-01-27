package turner

import (
	"math"
	"math/rand"
)

type Material interface {
	Scatter(rayIn Ray, hitRec *HitRecord) (hit bool, Attenuation Vec3, scatteredRay Ray)
}

type Lambertian struct {
	Albedo Vec3
}

func (m Lambertian) Scatter(rayIn Ray, hitRec *HitRecord) (bool, Vec3, Ray) {

	target := hitRec.P.Add(hitRec.Normal).Add(randomPointInUnitSphere())
	bounce := Ray{hitRec.P, target.Minus(hitRec.P)}
	return true, m.Albedo, bounce
}

func randomPointInUnitSphere() Vec3 {
	p := Vec3{}
	for {
		p = Vec3{rand.Float64(), rand.Float64(), rand.Float64()}.MultiplyScalar(2).Minus(Vec3{1, 1, 1})
		if p.SquaredLength() < 1 {
			break
		}
	}
	return p

}

type Metal struct {
	Albedo Vec3
	Fuzz   float64
}

func (m Metal) Scatter(rayIn Ray, hitRec *HitRecord) (bool, Vec3, Ray) {
	reflected := rayIn.Direction.UnitVector().Reflect(hitRec.Normal)
	scattered := Ray{hitRec.P, reflected.Add(randomPointInUnitSphere().MultiplyScalar(m.Fuzz))}

	return (scattered.Direction.Dot(hitRec.Normal) > 0), m.Albedo, scattered
}

func Refract(vec, normal Vec3, niOverNt float64) (bool, Vec3) {
	uv := vec.UnitVector()
	refracted := Vec3{}
	dt := uv.Dot(normal)
	discriminant := 1.0 - niOverNt*niOverNt*(1-dt*dt)
	if discriminant > 0 {
		refracted = uv.Minus(normal.MultiplyScalar(dt)).MultiplyScalar(niOverNt).Minus(normal.MultiplyScalar(math.Sqrt(discriminant)))
		return true, refracted
	}
	return false, refracted
}

type Dielectric struct {
	RefractiveIndex float64
}

func (m Dielectric) Scatter(rayIn Ray, hitRec *HitRecord) (bool, Vec3, Ray) {
	outwardNormal := Vec3{}
	refracted := Vec3{}
	scattered := Ray{}
	reflected := rayIn.Direction.Reflect(hitRec.Normal)
	var niOverNt float64
	attenuation := Vec3{1, 1, 1}
	var reflectProb float64
	var cosine float64
	if rayIn.Direction.Dot(hitRec.Normal) > 0 {
		outwardNormal = hitRec.Normal.MultiplyScalar(-1)
		niOverNt = m.RefractiveIndex
		cosine = m.RefractiveIndex * rayIn.Direction.Dot(hitRec.Normal) / rayIn.Direction.Length()

	} else {
		outwardNormal = hitRec.Normal
		niOverNt = 1.0 / m.RefractiveIndex
		cosine = -rayIn.Direction.Dot(hitRec.Normal) / rayIn.Direction.Length()
	}
	isRefracted, refracted := Refract(rayIn.Direction, outwardNormal, niOverNt)
	if isRefracted {
		reflectProb = schlick(cosine, m.RefractiveIndex)
	} else {
		reflectProb = 1.0
		scattered = Ray{hitRec.P, reflected}
	}
	if rand.Float64() < reflectProb {
		scattered = Ray{hitRec.P, reflected}
	} else {
		scattered = Ray{hitRec.P, refracted}
	}
	return true, attenuation, scattered
}

func schlick(cosine, refractiveIndex float64) float64 {
	r0 := (1 - refractiveIndex) / (1 + refractiveIndex)
	r0 = r0 * r0
	return r0 + (1-r0)*math.Pow((1-cosine), 5)
}
