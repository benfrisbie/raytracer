package scene

import (
	"image/color"

	"github.com/benfrisbie/raytracer/geometry"
	"github.com/benfrisbie/raytracer/geometry/renderable"
	"github.com/benfrisbie/raytracer/geometry/shape"
	"github.com/benfrisbie/raytracer/material"
)

type Scene1 struct {
	Scene
	Renderables         []renderable.Renderable
	RenderablesNoLights []renderable.Renderable
	Lights              []renderable.Renderable
}

func (s Scene1) GetRenderables() []renderable.Renderable {
	return s.Renderables
}

func (s Scene1) GetRenderablesNoLights() []renderable.Renderable {
	return s.RenderablesNoLights
}

func (s Scene1) GetLights() []renderable.Renderable {
	return s.Lights
}

func NewScene() *Scene1 {
	var scene Scene1 = Scene1{}

	var s shape.Shape
	var m material.Material

	// Spheres
	center := geometry.Vector{X: 0, Y: -7, Z: -15}
	radius := 3.0
	s = shape.Sphere{Center: center, Radius: radius}
	m = material.Matte{Color: color.NRGBA{R: 0, G: 0, B: 255, A: 255}}
	r := renderable.Renderable{Shape: s, Material: m}
	scene.Renderables = append(scene.Renderables, r)

	// center = geometry.Vector{X: 1, Y: -1, Z: -7}
	// radius = 1.0
	// s = shape.Sphere{Center: center, Radius: radius}
	// m = material.Matte{Color: color.NRGBA{R: 0, G: 255, B: 0, A: 255}}
	// r = renderable.Renderable{Shape: s, Material: m}
	// scene.Renderables = append(scene.Renderables, r)

	// center = geometry.Vector{X: 1, Y: 0.5, Z: -6}
	// radius = 0.4
	// s = shape.Sphere{Center: center, Radius: radius}
	// m = material.Matte{Color: color.NRGBA{R: 255, G: 0, B: 0, A: 255}}
	// r = renderable.Renderable{Shape: s, Material: m}
	// scene.Renderables = append(scene.Renderables, r)

	// white back wall
	vertices4 := [4]geometry.Vector{
		geometry.Vector{X: -10, Y: 10, Z: -20},
		geometry.Vector{X: 10, Y: 10, Z: -20},
		geometry.Vector{X: 10, Y: -10, Z: -20},
		geometry.Vector{X: -10, Y: -10, Z: -20}}
	s = shape.NewQuadrilateral(vertices4)
	m = material.Matte{Color: color.NRGBA{R: 220, G: 220, B: 220, A: 255}}
	r = renderable.Renderable{Shape: s, Material: m}
	scene.Renderables = append(scene.Renderables, r)

	// white floor
	vertices4 = [4]geometry.Vector{
		geometry.Vector{X: -10, Y: -10, Z: -20},
		geometry.Vector{X: 10, Y: -10, Z: -20},
		geometry.Vector{X: 10, Y: -10, Z: -10},
		geometry.Vector{X: -10, Y: -10, Z: -10}}
	s = shape.NewQuadrilateral(vertices4)
	r = renderable.Renderable{Shape: s, Material: m}
	scene.Renderables = append(scene.Renderables, r)

	// white ceiling
	vertices4 = [4]geometry.Vector{
		geometry.Vector{X: -10, Y: 10, Z: -20},
		geometry.Vector{X: 10, Y: 10, Z: -20},
		geometry.Vector{X: 10, Y: 10, Z: -10},
		geometry.Vector{X: -10, Y: 10, Z: -10}}
	s = shape.NewQuadrilateral(vertices4)
	r = renderable.Renderable{Shape: s, Material: m}
	scene.Renderables = append(scene.Renderables, r)

	// red left wall
	vertices4 = [4]geometry.Vector{
		geometry.Vector{X: -10, Y: 10, Z: -10},
		geometry.Vector{X: -10, Y: 10, Z: -20},
		geometry.Vector{X: -10, Y: -10, Z: -20},
		geometry.Vector{X: -10, Y: -10, Z: -10}}
	s = shape.NewQuadrilateral(vertices4)
	m = material.Matte{Color: color.NRGBA{R: 230, G: 60, B: 30, A: 255}}
	r = renderable.Renderable{Shape: s, Material: m}
	scene.Renderables = append(scene.Renderables, r)

	// green right wall
	vertices4 = [4]geometry.Vector{
		geometry.Vector{X: 10, Y: 10, Z: -10},
		geometry.Vector{X: 10, Y: 10, Z: -20},
		geometry.Vector{X: 10, Y: -10, Z: -20},
		geometry.Vector{X: 10, Y: -10, Z: -10}}
	s = shape.NewQuadrilateral(vertices4)
	m = material.Matte{Color: color.NRGBA{R: 0, G: 120, B: 40, A: 255}}
	r = renderable.Renderable{Shape: s, Material: m}
	scene.Renderables = append(scene.Renderables, r)

	// yellow ceiling light
	// vertices4 = [4]geometry.Vector{
	// 	geometry.Vector{X: -2, Y: 9.9, Z: -17},
	// 	geometry.Vector{X: 2, Y: 9.9, Z: -17},
	// 	geometry.Vector{X: 2, Y: 9.9, Z: -13},
	// 	geometry.Vector{X: -2, Y: 9.9, Z: -13}}
	// s = shape.NewQuadrilateral(vertices4)
	// m = material.Light{Color: color.NRGBA{R: 255, G: 255, B: 0, A: 255}}
	// r = renderable.Renderable{Shape: s, Material: m}
	// scene.Renderables = append(scene.Renderables, r)

	center = geometry.Vector{X: 0, Y: 9, Z: -15}
	radius = 1
	s = shape.Sphere{Center: center, Radius: radius}
	m = material.Light{Color: color.NRGBA{R: 255, G: 255, B: 0, A: 255}}
	r = renderable.Renderable{Shape: s, Material: m}
	scene.Renderables = append(scene.Renderables, r)

	for _, r := range scene.Renderables {
		if _, ok := r.Material.(material.Light); ok {
			scene.Lights = append(scene.Lights, r)
		} else {
			scene.RenderablesNoLights = append(scene.RenderablesNoLights, r)
		}
	}

	return &scene
}
