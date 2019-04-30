package texture

import "github.com/alisdairrankine/turner"

type colour struct {
	rgb turner.Vec3
}

func Colour(rgb turner.Vec3) Texture {
	return &colour{
		rgb: rgb,
	}
}

func (c colour) Value(u, v float64, p *turner.Vec3) turner.Vec3 {
	return c.rgb
}
