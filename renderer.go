package turner

import (
	"fmt"
	"image"
	"image/png"
	"math"
	"math/rand"
	"os"
	"sync"
	"time"
)

type Renderer struct {
	world    Hitable
	camera   Camera
	width    int
	height   int
	filename string
	samples  int
	depth    int
}

func NewRenderer(world Hitable, camera Camera, width, height int, filename string, samples, depth int) *Renderer {
	return &Renderer{
		world:    world,
		camera:   camera,
		width:    width,
		height:   height,
		filename: filename,
		samples:  samples,
		depth:    depth,
	}
}

func (r *Renderer) Render() {
	rand.Seed(time.Now().Unix())

	photonCount := r.width * r.height * r.samples

	img := image.NewRGBA(image.Rect(
		0, 0, r.width, r.height,
	))

	fmt.Printf("starting raytrace: %d Photons ready\n", photonCount)
	i := 0
	fmt.Println()
	for y := r.height - 1; y >= 0; y-- {
		for x := 0; x < r.width; x++ {
			if i%1000 == 0 {
				fmt.Printf("%d/%d\r", i, photonCount)
			}
			colour := Vec3{0, 0, 0}
			m := &sync.Mutex{}
			wg := &sync.WaitGroup{}
			for s := 0; s < r.samples; s++ {
				u := (float64(x) + rand.Float64()) / float64(r.width)
				v := (float64(r.height-y) + rand.Float64()) / float64(r.height)
				ray := r.camera.Ray(u, v)
				wg.Add(1)
				func(w *sync.WaitGroup, m *sync.Mutex) {
					defer w.Done()
					c := r.rayColour(ray, r.world, 0)
					m.Lock()
					defer m.Unlock()
					colour = colour.Add(c)
				}(wg, m)

			}
			wg.Wait()
			colour = colour.DivideScalar(float64(r.samples))
			img.Set(x, y, colour.Colour())
			i += r.samples
		}
	}
	fmt.Printf("%d/%d\n", i, r.width*r.height)
	fmt.Println("Saving file")

	file, _ := os.Create(r.filename)
	defer file.Close()
	png.Encode(file, img)
}

func (r *Renderer) rayColour(ray Ray, world Hitable, depth int) Vec3 {
	if hitRec, hit := world.Hit(&ray, 0.01, math.MaxFloat64); hit {
		if depth < r.depth {
			if hit, attenuation, scatteredRay := hitRec.Mat.Scatter(ray, hitRec); hit {
				return attenuation.MultiplyVector(r.rayColour(scatteredRay, world, depth+1))
			}
		}
		return Vec3{}
	}
	unitDir := ray.Direction.UnitVector()
	t := 0.5 * (unitDir.Y + 1)
	return Vec3{1, 1, 1}.MultiplyScalar(1.0 - t).Add(Vec3{0.5, 0.7, 1.0}.MultiplyScalar(t))
}
