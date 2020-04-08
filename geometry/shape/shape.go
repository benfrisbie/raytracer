package shape

import (
	"github.com/benfrisbie/raytracer/geometry"
)

// Shape represents a shape in 3D space
//
// CheckForCollision determines if a Ray collides with this shape
//
// NormalAtLocation calculates the normal vector at a given location on this shape
type Shape interface {
	CheckForCollision(ray geometry.Ray) *float64
	NormalAtLocation(loc geometry.Vector) geometry.Vector
}
