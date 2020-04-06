package geometry

import (
	"fmt"
	"math"
)

// Vector represents a vector or point in 3d space
type Vector struct {
	X, Y, Z float64
}

func (v Vector) String() string {
	return fmt.Sprintf("<%v, %v, %v>", v.X, v.Y, v.Z)
}

// Add computes the vector addition of v1 + v2
func (v1 Vector) Add(v2 Vector) Vector {
	return Vector{v1.X + v2.X, v1.Y + v2.Y, v1.Z + v2.Z}
}

// Sub computes the vector subtraction of v1 - v2
func (v1 Vector) Sub(v2 Vector) Vector {
	return Vector{v1.X - v2.X, v1.Y - v2.Y, v1.Z - v2.Z}
}

// VectorTo computes the vector from point v1 to point v2, which is simply v2 - v1
func (p1 Vector) VectorTo(p2 Vector) Vector {
	return p2.Sub(p1)
}

// Dot computes the vector dot product of v1 · v2
func (v1 Vector) Dot(v2 Vector) float64 {
	return v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z
}

// Cross computes the vector cross product of v1 x v2
func (v1 Vector) Cross(v2 Vector) Vector {
	return Vector{v1.Y*v2.Z - v1.Z*v2.Y, v1.Z*v2.X - v1.X*v2.Z, v1.X*v2.Y - v1.Y*v2.X}
}

// Scale computes the vector scalar of v1 * s
func (v Vector) Scale(s float64) Vector {
	return Vector{s * v.X, s * v.Y, s * v.Z}
}

// Length computes the vector length or magnitude of v1
func (v Vector) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

// Normalize computes the normalized vector of v1
func (v Vector) Normalize() Vector {
	norm := 1.0 / v.Length()
	return Vector{v.X * norm, v.Y * norm, v.Z * norm}
}

// Reflect computes the reflected vector of v across a normal vector
func (v Vector) Reflect(normal Vector) Vector {
	return v.Sub(normal.Scale(2 * v.Dot(normal)))
}
