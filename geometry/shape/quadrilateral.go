package shape

import (
	"fmt"
	"math"

	"github.com/benfrisbie/raytracer/geometry"
)

type Quadrilateral struct {
	Shape
	Vertices  [4]geometry.Vector
	triangles [2]Triangle
}

func NewQuadrilateral(vertices [4]geometry.Vector) Quadrilateral {
	rect := Quadrilateral{Vertices: vertices}
	rect.triangles = [2]Triangle{
		Triangle{Vertices: [3]geometry.Vector{
			rect.Vertices[0],
			rect.Vertices[1],
			rect.Vertices[2]}},
		Triangle{Vertices: [3]geometry.Vector{
			rect.Vertices[0],
			rect.Vertices[2],
			rect.Vertices[3]}}}

	return rect
}

func (rect Quadrilateral) String() string {
	return fmt.Sprintf("Quadrilateral(%v, %v, %v, %v)", rect.Vertices[0], rect.Vertices[1], rect.Vertices[2], rect.Vertices[3])
}

func (rect Quadrilateral) CheckForCollision(ray geometry.Ray) (Shape, float64) {
	var tMin float64 = math.MaxFloat64
	var shapeMin Shape = nil

	for _, triangle := range rect.triangles {
		s, t := triangle.CheckForCollision(ray)
		if s != nil && (shapeMin == nil || t < tMin) {
			tMin = t
			shapeMin = s
		}
	}
	return shapeMin, tMin
}
