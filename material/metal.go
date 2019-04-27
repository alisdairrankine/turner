package material

import "github.com/alisdairrankine/turner"

type metal struct {
	albedo turner.Vec3
	fuzz   float64
}

func Metal(albedo turner.Vec3, fuzz float64) turner.Material {
	return metal{
		albedo: albedo,
		fuzz:   fuzz,
	}
}

func (m metal) Scatter(rayIn turner.Ray, hitRec *turner.HitRecord) (bool, turner.Vec3, turner.Ray) {
	reflected := rayIn.Direction.UnitVector().Reflect(hitRec.Normal)
	scattered := turner.Ray{hitRec.P, reflected.Add(turner.RandomPointInUnitSphere().MultiplyScalar(m.fuzz)), rayIn.Time}

	return (scattered.Direction.Dot(hitRec.Normal) > 0), m.albedo, scattered
}
