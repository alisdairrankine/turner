package texture

import "github.com/alisdairrankine/turner"

type Texture interface {
	Value(u, v float64, p *turner.Vec3) turner.Vec3
}
