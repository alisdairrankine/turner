package turner

import (
	"math/rand"
)

func RandomPointInUnitSphere() Vec3 {
	p := Vec3{}
	for {
		p = Vec3{rand.Float64(), rand.Float64(), rand.Float64()}.MultiplyScalar(2).Minus(Vec3{1, 1, 1})
		if p.SquaredLength() < 1 {
			break
		}
	}
	return p
}

func RandomPointInUnitDisc() Vec3 {
	p := Vec3{}
	for {
		p = Vec3{rand.Float64(), rand.Float64(), 0.0}.MultiplyVector(Vec3{1, 1, 0}).MultiplyScalar(2)
		if p.Dot(p) < 1 {
			break
		}
	}
	return p

}
