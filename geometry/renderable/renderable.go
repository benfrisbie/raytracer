package renderable

import (
	"github.com/benfrisbie/raytracer/geometry"
	"github.com/benfrisbie/raytracer/geometry/shape"
	"github.com/benfrisbie/raytracer/material"
)

type Renderable struct {
	Shape    shape.Shape
	Material material.Material
}

// Collision represents a Ray colliding with a Shape in 3D space
type Collision struct {
	T          float64
	shape      shape.Shape
	Renderable Renderable
	Location   geometry.Vector
	Normal     geometry.Vector
}

// ClosestCollision returns the closest collision to any Shape on Ray r. Closest means the shortest distance from the rays origin.
// If there was no collision, return nil
func ClosestCollision(ray geometry.Ray, renderables []Renderable) *Collision {
	var closest *Collision = nil
	for _, rend := range renderables {
		t := rend.Shape.CheckForCollision(ray)
		if t != nil && (closest == nil || *t < closest.T) {
			closest = &Collision{}
			closest.T = *t
			closest.Renderable = rend
			closest.shape = rend.Shape
		}
	}
	if closest != nil {
		closest.Location = ray.PointOnRay(closest.T)
		closest.Normal = closest.shape.NormalAtLocation(closest.Location)
	}
	return closest
}

// CollisionCloserThan checks if there was a collision with any shape on a ray closer than location t
func CollisionCloserThan(r geometry.Ray, renderables []Renderable, t float64) bool {
	c := ClosestCollision(r, renderables)
	if c != nil && r.Origin.VectorTo(r.PointOnRay(c.T)).Length() < t {
		return true
	}
	return false
}
