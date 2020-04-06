package entity

import (
	"raytracer/geometry"
)

// Collision represents a Ray colliding with a Renderable in 3D space
type Collision struct {
	T          float64
	Renderable Renderable
	Location   geometry.Vector
	Normal     geometry.Vector
}

// ClosestCollision returns the closest collision to any Entity on Ray r. Closest means the shortest distance from the rays origin.
// If there was no intersection, return nil
func ClosestCollision(ray geometry.Ray, renderables []Renderable) *Collision {
	var closest *Collision = nil
	for _, r := range renderables {
		c := r.Entity.CheckForCollision(ray)
		if c != nil && (closest == nil || c.T < closest.T) {
			closest = c
			closest.Renderable = r
		}
	}
	if closest != nil {
		closest.Location = ray.PointOnRay(closest.T)
		closest.Normal = closest.Renderable.Entity.NormalAtPoint(closest.Location)
	}
	return closest
}

func CollisionCloserThan(r geometry.Ray, renderables []Renderable, t float64) bool {
	c := ClosestCollision(r, renderables)
	if c != nil && r.Origin.VectorTo(r.PointOnRay(c.T)).Length() < t {
		return true
	}
	return false
}
