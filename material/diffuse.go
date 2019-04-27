package material

import (
	"github.com/alisdairrankine/turner"
)

type lambertian struct {
	albedo turner.Vec3
}

func Diffuse(albedo turner.Vec3) turner.Material {
	return lambertian{albedo: albedo}
}

func (l lambertian) Scatter(rayIn turner.Ray, hitRec *turner.HitRecord) (bool, turner.Vec3, turner.Ray) {

	target := hitRec.P.Add(hitRec.Normal).Add(turner.RandomPointInUnitSphere())
	bounce := turner.Ray{hitRec.P, target.Minus(hitRec.P), rayIn.Time}
	return true, l.albedo, bounce
}
