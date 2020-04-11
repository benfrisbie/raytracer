package shape

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/benfrisbie/raytracer/geometry"
)

type Quadrilateral struct {
	Shape
	Vertices  [4]geometry.Vector
	triangles [2]Triangle
}

func NewQuadrilateral(vertices [4]geometry.Vector) Quadrilateral {
	quad := Quadrilateral{Vertices: vertices}
	quad.triangles = [2]Triangle{
		Triangle{Vertices: [3]geometry.Vector{
			quad.Vertices[0],
			quad.Vertices[1],
			quad.Vertices[2]}},
		Triangle{Vertices: [3]geometry.Vector{
			quad.Vertices[0],
			quad.Vertices[2],
			quad.Vertices[3]}}}

	return quad
}

func (quad Quadrilateral) String() string {
	return fmt.Sprintf("Quadrilateral(%v, %v, %v, %v)", quad.Vertices[0], quad.Vertices[1], quad.Vertices[2], quad.Vertices[3])
}

func (quad Quadrilateral) CheckForCollision(ray geometry.Ray) (Shape, float64) {
	var tMin float64 = math.MaxFloat64
	var shapeMin Shape = nil

	for _, triangle := range quad.triangles {
		s, t := triangle.CheckForCollision(ray)
		if s != nil && (shapeMin == nil || t < tMin) {
			tMin = t
			shapeMin = s
		}
	}
	return shapeMin, tMin
}

func (quad Quadrilateral) RandomLocationOn() geometry.Vector {
	// Choose a triangle based on area
	area := quad.triangles[0].Area() + quad.triangles[1].Area()
	if rand.Float64()*area < quad.triangles[0].Area() {
		return quad.triangles[0].RandomLocationOn()
	}
	return quad.triangles[1].RandomLocationOn()
}
