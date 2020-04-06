package geometry

import (
	"fmt"
)

// Ray represents a ray in 3D space. A ray has an origin and direction
type Ray struct {
	Origin    Vector
	Direction Vector
}

func (r Ray) String() string {
	return fmt.Sprintf("%v -> %v", r.Origin, r.Direction)
}

// PointOnRay computes the point on Ray r that is a of distance t from its origin
func (r Ray) PointOnRay(t float64) Vector {
	return r.Origin.Add(r.Direction.Scale(t))
}

// OffsetOrigin computes a new Ray similar to r with the only difference being its origin is slightly offset in the direction of r.
// Argument is optional and the default value offset is 0.0001
func (r Ray) OffsetOrigin(offset ...float64) Ray {
	// Default value
	o := 0.0001
	if len(offset) > 0 {
		o = offset[0]
	}

	return Ray{
		Origin:    r.PointOnRay(o),
		Direction: r.Direction}
}

// ReflectRay calculates the reflected ray at a collison point over a normal vector
func (r Ray) ReflectRay(collision Vector, normal Vector) Ray {
	reflect := Ray{Origin: collision, Direction: r.Direction.Reflect(normal)}
	// TODO: might need this?
	// .OffsetOrigin()
	return reflect
}
