package camera

import (
	"image"
	"image/color"
	"math"
	"math/rand"
	"sync"

	"github.com/benfrisbie/raytracer/entity"
	"github.com/benfrisbie/raytracer/geometry"
	"github.com/benfrisbie/raytracer/material"
	"github.com/benfrisbie/raytracer/scene"
	"github.com/cheggaaa/pb"
)

type Camera struct {
	image       *image.NRGBA
	width       int
	widthFloat  float64
	height      int
	heightFloat float64
	aspectRatio float64

	fov       float64
	scale     float64
	antialias int
	location  geometry.Vector

	scene scene.Scene
}

func NewCamera(width int, height int, antialias int, fov float64) *Camera {
	c := Camera{}

	c.location = geometry.Vector{X: 0, Y: 0, Z: 0}
	c.width = width
	c.height = height
	c.widthFloat = float64(c.width)
	c.heightFloat = float64(c.height)
	c.aspectRatio = c.widthFloat / c.heightFloat
	c.image = image.NewNRGBA(image.Rect(0, 0, c.width, c.height))
	c.antialias = antialias
	c.fov = fov
	c.scale = math.Tan(0.5 * c.fov * math.Pi / 180)

	return &c
}

func (c *Camera) RenderSceneToImage(scene scene.Scene) *image.NRGBA {
	c.scene = scene

	var wg sync.WaitGroup
	wg.Add(c.width * c.height)
	progressBar := pb.StartNew(c.width * c.height)

	// Calculate the color for each pixel of the image
	for x := 0; x < c.width; x++ {
		for y := 0; y < c.height; y++ {
			go c.calculatePixel(x, y, &wg, progressBar)
		}
	}

	wg.Wait()
	progressBar.Finish()

	return c.image
}

func (c Camera) calculatePixel(x int, y int, wg *sync.WaitGroup, progressBar *pb.ProgressBar) {
	var cm material.ColorMixer

	for i := 0; i < c.antialias; i++ {
		pixelPos := geometry.Vector{}
		pixelPos.Z = -1
		pixelPos.X = (2*(float64(x)+rand.Float64())/c.widthFloat - 1) * c.aspectRatio * c.scale
		pixelPos.Y = (1 - 2*(float64(y)+rand.Float64())/c.heightFloat) * c.scale
		col := c.shootRay(geometry.Ray{Origin: c.location, Direction: c.location.VectorTo(pixelPos).Normalize()})
		cm.Add(col)
	}

	c.image.Set(x, y, cm.ToNRGBA())
	progressBar.Increment()
	wg.Done()
}

func (c Camera) shootRay(ray geometry.Ray) color.NRGBA {
	closest := entity.ClosestCollision(ray, c.scene.GetRenderables())
	if closest != nil {
		// if closest.Entity.GetMaterial().Reflective {
		// 	// shoot a ray in the reflected direction
		// 	return shootRay(ray.ReflectRay(closest.Point, closest.Normal))
		// }
		if _, ok := closest.Renderable.Material.(material.Light); ok {
			// light - return the color of the light
			return closest.Renderable.Material.GetColor()
		} else {
			// Shoot rays to light sources to see if we're in shadow
			var totalDiffuse float64 = 0
			for _, light := range c.scene.GetLights() {
				// TODO: all lights are being treated as point lights right now
				rayToLight := geometry.Ray{Origin: closest.Location, Direction: closest.Location.VectorTo(light.Entity.(entity.Sphere).Center)}.OffsetOrigin()
				if !entity.CollisionCloserThan(rayToLight, c.scene.GetRenderablesNoLights(), rayToLight.Direction.Length()) {
					totalDiffuse += math.Abs(rayToLight.Direction.Dot(closest.Normal))
				}
			}
			totalDiffuse /= float64(len(c.scene.GetLights()))
			totalDiffuse = math.Max(0.2, math.Min(1, totalDiffuse))

			c := closest.Renderable.Material.GetColor()
			c.R = uint8(totalDiffuse * float64(c.R))
			c.G = uint8(totalDiffuse * float64(c.G))
			c.B = uint8(totalDiffuse * float64(c.B))
			return c
		}
	}
	return color.NRGBA{0, 0, 0, 255} // No hit, return black
}
