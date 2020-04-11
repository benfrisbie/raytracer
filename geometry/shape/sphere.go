package shape

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/benfrisbie/raytracer/geometry"
)

type Sphere struct {
	Shape
	Center geometry.Vector
	Radius float64
}

func (sphere Sphere) String() string {
	return fmt.Sprintf("Sphere(%v, %v)", sphere.Center, sphere.Radius)
}

func (sphere Sphere) CheckForCollision(ray geometry.Ray) (Shape, float64) {
	cToE := sphere.Center.VectorTo(ray.Origin)
	a := ray.Direction.Dot(cToE)
	b := ray.Direction.Dot(ray.Direction)
	discrim := a*a - b*(cToE.Dot(cToE)-sphere.Radius*sphere.Radius)

	if discrim == 0 {
		t := -a / b
		return sphere, t
	} else if discrim > 0 {
		root := math.Sqrt(discrim)
		t := (-a - root) / b
		if t >= 0 {
			return sphere, t
		}

		t = (-a + root) / b
		if t >= 0 {
			return sphere, t
		}
	}

	return nil, 0
}

func (s Sphere) NormalAtLocation(loc geometry.Vector) geometry.Vector {
	return s.Center.VectorTo(loc).Normalize()
}

func (s Sphere) RandomLocationOn() geometry.Vector {
	x := rand.NormFloat64()
	y := rand.NormFloat64()
	z := rand.NormFloat64()
	return geometry.Vector{X: x, Y: y, Z: z}.Normalize().Scale(s.Radius)
}
