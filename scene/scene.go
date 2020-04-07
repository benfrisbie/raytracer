package scene

import (
	"github.com/benfrisbie/raytracer/entity"
)

type Scene interface {
	SetupScene() []entity.Renderable
}
