package texture

//Work in progress. not ready.

import (
	"github.com/alisdairrankine/turner"
	"math/rand"
)

type perlin struct {
	rndPool [256]float64
	permX   [256]int
	permY   [256]int
	permZ   [256]int
}

func rndPool() [256]float64 {
	p := [256]float64{}
	for i := 0; i < 256; i++ {
		p[i] = rand.Float64()
	}
	return p
}

func permute() {
	// p := int[256]
}

func (pn perlin) Noise(p turner.Vec3) float64 {
	// u := p.X - math.Floor(p.X)
	// v := p.Y - math.Floor(p.Y)
	// w := p.Z - math.Floor(p.Z)

	i := int(4*p.X) & 255
	j := int(4*p.Y) & 255
	k := int(4*p.Z) & 255

	return pn.rndPool[pn.permX[i]^pn.permX[j]^pn.permX[k]]

}
