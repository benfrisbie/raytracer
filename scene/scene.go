package scene

import "github.com/benfrisbie/raytracer/entity"

type Scene interface {
	NewScene() *Scene
	GetRenderables() []entity.Renderable
	GetRenderablesNoLights() []entity.Renderable
	GetLights() []entity.Renderable
}
