package entity

import (
	"github.com/benfrisbie/raytracer/material"
)

type Renderable struct {
	Entity   Entity
	Material material.Material
}
