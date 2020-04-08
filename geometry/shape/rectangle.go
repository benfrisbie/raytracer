package shape

import (
	"fmt"

	"github.com/benfrisbie/raytracer/geometry"
)

type Rectangle struct {
	Shape
	Vertices  [4]geometry.Vector
	triangles [2]Triangle
}

func NewRectangle(vertices [4]geometry.Vector) Rectangle {
	rect := Rectangle{Vertices: vertices}
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

func (rect Rectangle) String() string {
	return fmt.Sprintf("Rectangle(%v, %v, %v, %v)", rect.Vertices[0], rect.Vertices[1], rect.Vertices[2], rect.Vertices[3])
}

func (rect Rectangle) CheckForCollision(ray geometry.Ray) *float64 {
	for _, triangle := range rect.triangles {
		t := triangle.CheckForCollision(ray)
		if t != nil {
			return t
		}
	}
	return nil
}

func (rect Rectangle) NormalAtLocation(loc geometry.Vector) geometry.Vector {
	return rect.triangles[0].NormalAtLocation(loc)
}
