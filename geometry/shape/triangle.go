package shape

import (
	"fmt"

	"github.com/benfrisbie/raytracer/geometry"
)

const EPSILON = 0.000000001

type Triangle struct {
	Shape
	Vertices [3]geometry.Vector

	normal *geometry.Vector
}

func (triangle Triangle) String() string {
	return fmt.Sprintf("Triangle(%v, %v, %v)", triangle.Vertices[0], triangle.Vertices[1], triangle.Vertices[2])
}

func (triangle Triangle) CheckForCollision(ray geometry.Ray) (Shape, float64) {
	// compute plane's normal
	edge1 := triangle.Vertices[0].VectorTo(triangle.Vertices[1])
	edge2 := triangle.Vertices[0].VectorTo(triangle.Vertices[2])
	h := ray.Direction.Cross(edge2)
	a := edge1.Dot(h)
	if a > -EPSILON && a < EPSILON {
		// the ray is parallel to the triangle
		return nil, 0
	}
	f := 1.0 / a
	s := triangle.Vertices[0].VectorTo(ray.Origin) // rayOrigin - vertex0
	u := f * s.Dot(h)
	if u < 0.0 || u > 1.0 {
		return nil, 0
	}
	q := s.Cross(edge1)
	v := f * ray.Direction.Dot(q)
	if v < 0.0 || u+v > 1.0 {
		return nil, 0
	}
	// At this stage we can compute t to find out where the intersection point is on the line.
	t := f * edge2.Dot(q)
	if t > EPSILON && t < 1/EPSILON {
		return triangle, t
	}
	return nil, 0
}

func (t Triangle) NormalAtLocation(loc geometry.Vector) geometry.Vector {
	// Normal is the same for every location on the triangle. calculate and cache for future calls
	if t.normal == nil {
		// Normal is the cross product of two of the edges
		edge1 := t.Vertices[0].VectorTo(t.Vertices[1])
		edge2 := t.Vertices[0].VectorTo(t.Vertices[2])
		n := edge1.Cross(edge2)
		t.normal = &n
	}
	return *t.normal
}
