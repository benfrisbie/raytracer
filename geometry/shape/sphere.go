package shape

import (
	"fmt"
	"math"

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

func (sphere Sphere) CheckForCollision(ray geometry.Ray) *float64 {
	cToE := sphere.Center.VectorTo(ray.Origin)
	a := ray.Direction.Dot(cToE)
	b := ray.Direction.Dot(ray.Direction)
	discrim := a*a - b*(cToE.Dot(cToE)-sphere.Radius*sphere.Radius)

	if discrim == 0 {
		t := -a / b
		return &t
	} else if discrim > 0 {
		root := math.Sqrt(discrim)
		t := (-a - root) / b
		if t >= 0 {
			return &t
		}

		t = (-a + root) / b
		if t >= 0 {
			return &t
		}
	}

	return nil
}

func (s Sphere) NormalAtLocation(loc geometry.Vector) geometry.Vector {
	return s.Center.VectorTo(loc).Normalize()
}
