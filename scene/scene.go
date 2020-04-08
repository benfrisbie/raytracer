package scene

import "github.com/benfrisbie/raytracer/geometry/renderable"

type Scene interface {
	NewScene() *Scene
	GetRenderables() []renderable.Renderable
	GetRenderablesNoLights() []renderable.Renderable
	GetLights() []renderable.Renderable
}
