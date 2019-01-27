package material

import (
	"math"
	"math/rand"

	"github.com/alisdairrankine/turner"
)

type dielectric struct {
	refractiveIndex float64
}

func Dielectric(refractiveIndex float64) turner.Material {
	return dielectric{refractiveIndex: refractiveIndex}
}

func (d dielectric) Scatter(rayIn turner.Ray, hitRec *turner.HitRecord) (bool, turner.Vec3, turner.Ray) {
	outwardNormal := turner.Vec3{}
	refracted := turner.Vec3{}
	scattered := turner.Ray{}
	reflected := rayIn.Direction.Reflect(hitRec.Normal)
	var niOverNt float64
	attenuation := turner.Vec3{1, 1, 1}
	var reflectProb float64
	var cosine float64
	if rayIn.Direction.Dot(hitRec.Normal) > 0 {
		outwardNormal = hitRec.Normal.MultiplyScalar(-1)
		niOverNt = d.refractiveIndex
		cosine = d.refractiveIndex * rayIn.Direction.Dot(hitRec.Normal) / rayIn.Direction.Length()

	} else {
		outwardNormal = hitRec.Normal
		niOverNt = 1.0 / d.refractiveIndex
		cosine = -rayIn.Direction.Dot(hitRec.Normal) / rayIn.Direction.Length()
	}
	isRefracted, refracted := rayIn.Direction.Refract(outwardNormal, niOverNt)
	if isRefracted {
		reflectProb = schlick(cosine, d.refractiveIndex)
	} else {
		reflectProb = 1.0
		scattered = turner.Ray{hitRec.P, reflected}
	}
	if rand.Float64() < reflectProb {
		scattered = turner.Ray{hitRec.P, reflected}
	} else {
		scattered = turner.Ray{hitRec.P, refracted}
	}
	return true, attenuation, scattered
}

func schlick(cosine, refractiveIndex float64) float64 {
	r0 := (1 - refractiveIndex) / (1 + refractiveIndex)
	r0 = r0 * r0
	return r0 + (1-r0)*math.Pow((1-cosine), 5)
}
