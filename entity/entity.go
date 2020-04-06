package entity

import (
	"raytracer/geometry"
)

// Entity represents an object in 3D space
//
// CheckForCollision determines if a Ray collides with this entity
//
// NormalAtPoint calculates the normal vector at a given point on the entity.
type Entity interface {
	CheckForCollision(r geometry.Ray) *Collision
	NormalAtPoint(p geometry.Vector) geometry.Vector
}
