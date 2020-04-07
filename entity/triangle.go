package entity

import (
	"fmt"
	"github.com/benfrisbie/raytracer/geometry"
	"github.com/benfrisbie/raytracer/material"
)

const EPSILON = 0.000000001

type Triangle struct {
	Entity
	Vertices [3]geometry.Vector
	Material material.Material

	normal *geometry.Vector
}

func (t Triangle) String() string {
	return fmt.Sprintf("Triangle(%v, %v, %v)", t.Vertices[0], t.Vertices[1], t.Vertices[2])
}

func (tri Triangle) CheckForCollision(r geometry.Ray) *Collision {
	// compute plane's normal
	edge1 := tri.Vertices[0].VectorTo(tri.Vertices[1])
	edge2 := tri.Vertices[0].VectorTo(tri.Vertices[2])
	h := r.Direction.Cross(edge2)
	a := edge1.Dot(h)
	if a > -EPSILON && a < EPSILON {
		// the ray is parallel to the triangle
		return nil
	}
	f := 1.0 / a
	s := tri.Vertices[0].VectorTo(r.Origin) // rayOrigin - vertex0
	u := f * s.Dot(h)
	if u < 0.0 || u > 1.0 {
		return nil
	}
	q := s.Cross(edge1)
	v := f * r.Direction.Dot(q)
	if v < 0.0 || u+v > 1.0 {
		return nil
	}
	// At this stage we can compute t to find out where the intersection point is on the line.
	t := f * edge2.Dot(q)
	if t > EPSILON && t < 1/EPSILON {
		// ray intersection
		// outIntersectionPoint = rayOrigin + rayVector*t
		return &Collision{T: t}
	} else {
		// This means that there is a line intersection but not a ray intersection.
		return nil
	}
}

func (t Triangle) NormalAtPoint(p geometry.Vector) geometry.Vector {
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
