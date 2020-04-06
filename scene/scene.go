package scene

import (
	"raytracer/entity"
)

type Scene interface {
	SetupScene() []entity.Renderable
}
