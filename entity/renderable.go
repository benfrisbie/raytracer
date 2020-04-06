package entity

import (
	"raytracer/material"
)

type Renderable struct {
	Entity   Entity
	Material material.Material
}
