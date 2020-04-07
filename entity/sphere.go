package entity

import (
	"fmt"
	"math"
	"github.com/benfrisbie/raytracer/geometry"
)

type Sphere struct {
	Entity
	Center geometry.Vector
	Radius float64
}

func (s Sphere) String() string {
	return fmt.Sprintf("Sphere(%v, %v)", s.Center, s.Radius)
}

func (s Sphere) CheckForCollision(r geometry.Ray) *Collision {
	cToE := s.Center.VectorTo(r.Origin)
	a := r.Direction.Dot(cToE)
	b := r.Direction.Dot(r.Direction)
	discrim := a*a - b*(cToE.Dot(cToE)-s.Radius*s.Radius)

	if discrim == 0 {
		return &Collision{T: -a / b}
	} else if discrim > 0 {
		root := math.Sqrt(discrim)
		t := (-a - root) / b
		if t >= 0 {
			return &Collision{T: t}
		}

		t = (-a + root) / b
		if t >= 0 {
			return &Collision{T: t}
		}
	}

	return nil
}

func (s Sphere) NormalAtPoint(p geometry.Vector) geometry.Vector {
	return s.Center.VectorTo(p).Normalize()
}
