package shape

import (
	"github.com/benfrisbie/raytracer/geometry"
)

// Shape represents a shape in 3D space
type Shape interface {
	// CheckForCollision determines if a Ray collides with this shape
	CheckForCollision(ray geometry.Ray) (Shape, float64)

	// NormalAtLocation calculates the normal vector at a given location on this shape
	NormalAtLocation(loc geometry.Vector) geometry.Vector

	// RandomLocationOn return a random location on the surface of the shape
	RandomLocationOn() geometry.Vector
}
