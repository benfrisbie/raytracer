package scene

import (
	"image/color"

	"github.com/benfrisbie/raytracer/entity"
	"github.com/benfrisbie/raytracer/geometry"
	"github.com/benfrisbie/raytracer/material"
)

type Scene1 struct {
	Scene
	Renderables         []entity.Renderable
	RenderablesNoLights []entity.Renderable
	Lights              []entity.Renderable
}

func (s Scene1) GetRenderables() []entity.Renderable {
	return s.Renderables
}

func (s Scene1) GetRenderablesNoLights() []entity.Renderable {
	return s.RenderablesNoLights
}

func (s Scene1) GetLights() []entity.Renderable {
	return s.Lights
}

func NewScene() *Scene1 {
	var scene Scene1 = Scene1{}

	var e entity.Entity
	var m material.Material

	// Spheres
	center := geometry.Vector{X: -1, Y: -1, Z: -8}
	radius := 1.0
	e = entity.Sphere{Center: center, Radius: radius}
	m = material.Matte{Color: color.NRGBA{R: 0, G: 0, B: 255, A: 255}}
	renderable := entity.Renderable{Entity: e, Material: m}
	scene.Renderables = append(scene.Renderables, renderable)

	center = geometry.Vector{X: 1, Y: -1, Z: -7}
	radius = 1.0
	e = entity.Sphere{Center: center, Radius: radius}
	m = material.Matte{Color: color.NRGBA{R: 0, G: 255, B: 0, A: 255}}
	renderable = entity.Renderable{Entity: e, Material: m}
	scene.Renderables = append(scene.Renderables, renderable)

	center = geometry.Vector{X: 1, Y: 0.5, Z: -6}
	radius = 0.4
	e = entity.Sphere{Center: center, Radius: radius}
	m = material.Matte{Color: color.NRGBA{R: 255, G: 0, B: 0, A: 255}}
	renderable = entity.Renderable{Entity: e, Material: m}
	scene.Renderables = append(scene.Renderables, renderable)

	// Triangles
	// back wall
	vertices := [3]geometry.Vector{
		geometry.Vector{X: -10, Y: -2, Z: -10},
		geometry.Vector{X: 10, Y: -2, Z: -10},
		geometry.Vector{X: -10, Y: 2, Z: -10}}
	e = entity.Triangle{Vertices: vertices}
	m = material.Matte{Color: color.NRGBA{R: 128, G: 0, B: 128, A: 255}}
	renderable = entity.Renderable{Entity: e, Material: m}
	scene.Renderables = append(scene.Renderables, renderable)

	vertices = [3]geometry.Vector{
		geometry.Vector{X: 2.5, Y: -2, Z: -10},
		geometry.Vector{X: 2.5, Y: 2, Z: -10},
		geometry.Vector{X: -2.5, Y: 2, Z: -10}}
	e = entity.Triangle{Vertices: vertices}
	m = material.Matte{Color: color.NRGBA{R: 128, G: 0, B: 128, A: 255}}
	renderable = entity.Renderable{Entity: e, Material: m}
	scene.Renderables = append(scene.Renderables, renderable)

	// // floor
	// vertices = [3]geometry.Vector{
	// 	geometry.Vector{X: -2.5, Y: -2, Z: -2},
	// 	geometry.Vector{X: 2.5, Y: -2, Z: -2},
	// 	geometry.Vector{X: 2.5, Y: -2, Z: -10}}
	// e = entity.Triangle{Vertices: vertices}
	// m = material.Matte{Color: color.NRGBA{R: 255, G: 255, B: 255, A: 255}}
	// renderable = entity.Renderable{Entity: e, Material: m}
	// scene.Renderables = append(scene.Renderables, renderable)

	// vertices = [3]geometry.Vector{
	// 	geometry.Vector{X: -2.5, Y: -2, Z: -2},
	// 	geometry.Vector{X: 2.5, Y: -2, Z: -10},
	// 	geometry.Vector{X: -2.5, Y: -2, Z: -10}}
	// e = entity.Triangle{Vertices: vertices}
	// m = material.Matte{Color: color.NRGBA{R: 255, G: 255, B: 255, A: 255}}
	// renderable = entity.Renderable{Entity: e, Material: m}
	// scene.Renderables = append(scene.Renderables, renderable)

	// // ceiling
	// vertices = [3]geometry.Vector{
	// 	geometry.Vector{X: -2.5, Y: 2, Z: -2},
	// 	geometry.Vector{X: 2.5, Y: 2, Z: -2},
	// 	geometry.Vector{X: 2.5, Y: 2, Z: -10}}
	// e = entity.Triangle{Vertices: vertices}
	// m = material.Matte{Color: color.NRGBA{R: 255, G: 255, B: 255, A: 255}}
	// renderable = entity.Renderable{Entity: e, Material: m}
	// scene.Renderables = append(scene.Renderables, renderable)

	// vertices = [3]geometry.Vector{
	// 	geometry.Vector{X: -2.5, Y: 2, Z: -2},
	// 	geometry.Vector{X: 2.5, Y: 2, Z: -10},
	// 	geometry.Vector{X: -2.5, Y: 2, Z: -10}}
	// e = entity.Triangle{Vertices: vertices}
	// m = material.Matte{Color: color.NRGBA{R: 255, G: 255, B: 255, A: 255}}
	// renderable = entity.Renderable{Entity: e, Material: m}
	// scene.Renderables = append(scene.Renderables, renderable)

	// // left wall
	// vertices = [3]geometry.Vector{
	// 	geometry.Vector{X: -2.5, Y: -2, Z: -2},
	// 	geometry.Vector{X: -2.5, Y: -2, Z: -10},
	// 	geometry.Vector{X: -2.5, Y: 2, Z: -10}}
	// e = entity.Triangle{Vertices: vertices}
	// m = material.Matte{Color: color.NRGBA{R: 231, G: 76, B: 60, A: 255}}
	// renderable = entity.Renderable{Entity: e, Material: m}
	// scene.Renderables = append(scene.Renderables, renderable)

	// vertices = [3]geometry.Vector{
	// 	geometry.Vector{X: -2.5, Y: -2, Z: -2},
	// 	geometry.Vector{X: -2.5, Y: 2, Z: -10},
	// 	geometry.Vector{X: -2.5, Y: 2, Z: -2}}
	// e = entity.Triangle{Vertices: vertices}
	// m = material.Matte{Color: color.NRGBA{R: 231, G: 76, B: 60, A: 255}}
	// renderable = entity.Renderable{Entity: e, Material: m}
	// scene.Renderables = append(scene.Renderables, renderable)

	// // right wall
	// vertices = [3]geometry.Vector{
	// 	geometry.Vector{X: 2.5, Y: -2, Z: -2},
	// 	geometry.Vector{X: 2.5, Y: -2, Z: -10},
	// 	geometry.Vector{X: 2.5, Y: 2, Z: -10}}
	// e = entity.Triangle{Vertices: vertices}
	// m = material.Matte{Color: color.NRGBA{R: 5, G: 50, B: 255, A: 255}}
	// renderable = entity.Renderable{Entity: e, Material: m}
	// scene.Renderables = append(scene.Renderables, renderable)

	// vertices = [3]geometry.Vector{
	// 	geometry.Vector{X: 2.5, Y: -2, Z: -2},
	// 	geometry.Vector{X: 2.5, Y: 2, Z: -10},
	// 	geometry.Vector{X: 2.5, Y: 2, Z: -2}}
	// e = entity.Triangle{Vertices: vertices}
	// m = material.Matte{Color: color.NRGBA{R: 5, G: 50, B: 255, A: 255}}
	// renderable = entity.Renderable{Entity: e, Material: m}
	// scene.Renderables = append(scene.Renderables, renderable)

	// lights
	center = geometry.Vector{X: 0, Y: 1.5, Z: -5}
	radius = 0.1
	e = entity.Sphere{Center: center, Radius: radius}
	m = material.Light{Color: color.NRGBA{R: 255, G: 255, B: 0, A: 255}}
	renderable = entity.Renderable{Entity: e, Material: m}
	scene.Renderables = append(scene.Renderables, renderable)

	center = geometry.Vector{X: 1.75, Y: 2, Z: -6}
	radius = 0.1
	e = entity.Sphere{Center: center, Radius: radius}
	m = material.Light{Color: color.NRGBA{R: 255, G: 255, B: 0, A: 255}}
	renderable = entity.Renderable{Entity: e, Material: m}
	scene.Renderables = append(scene.Renderables, renderable)

	center = geometry.Vector{X: -1.75, Y: 2, Z: -5}
	radius = 0.1
	e = entity.Sphere{Center: center, Radius: radius}
	m = material.Light{Color: color.NRGBA{R: 255, G: 255, B: 0, A: 255}}
	renderable = entity.Renderable{Entity: e, Material: m}
	scene.Renderables = append(scene.Renderables, renderable)

	for _, r := range scene.Renderables {
		if _, ok := r.Material.(material.Light); ok {
			scene.Lights = append(scene.Lights, r)
		} else {
			scene.RenderablesNoLights = append(scene.RenderablesNoLights, r)
		}
	}

	return &scene
}
