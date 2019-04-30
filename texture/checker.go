package texture

import (
	"github.com/alisdairrankine/turner"
	"math"
)

type checker struct {
	a Texture
	b Texture
}

func Checker(a, b Texture) Texture {
	return &checker{
		a: a,
		b: b,
	}
}

func (c checker) Value(u, v float64, p *turner.Vec3) turner.Vec3 {
	sine := math.Sin(10*p.X) * math.Sin(10*p.Y) * math.Sin(10*p.Z)
	if sine < 0 {
		return c.a.Value(u, v, p)
	}
	return c.b.Value(u, v, p)
}
