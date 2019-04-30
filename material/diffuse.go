package material

import (
	"github.com/alisdairrankine/turner"
	"github.com/alisdairrankine/turner/texture"
)

type lambertian struct {
	albedo texture.Texture
}

func Diffuse(albedo texture.Texture) turner.Material {
	return lambertian{albedo: albedo}
}

func (l lambertian) Scatter(rayIn turner.Ray, hitRec *turner.HitRecord) (bool, turner.Vec3, turner.Ray) {

	target := hitRec.P.Add(hitRec.Normal).Add(turner.RandomPointInUnitSphere())
	bounce := turner.Ray{hitRec.P, target.Minus(hitRec.P), rayIn.Time}
	return true, l.albedo.Value(0, 0, &hitRec.P), bounce
}
