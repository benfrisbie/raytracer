package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"math/rand"
	"os"
	"raytracer/entity"
	"raytracer/geometry"
	"raytracer/material"
	"raytracer/scene"
	"sync"
	"time"

	"github.com/cheggaaa/pb/v3"
)

const WidthHeightDefault int = 512
const AntialiasDefault int = 4

var width int
var height int
var antialias int
var camPos geometry.Vector

var entities []entity.Renderable
var entitiesNoLights []entity.Renderable
var lights []entity.Renderable

func main() {
	var output string
	flag.IntVar(&width, "width", WidthHeightDefault, "width of the image")
	flag.IntVar(&height, "height", WidthHeightDefault, "height of the image")
	flag.IntVar(&antialias, "antialias", AntialiasDefault, "antialias sets how many rays per pixel should be shot into the scene. increasing this improves quality around edges and the expense of performance")
	flag.StringVar(&output, "output", "test.png", "name of the output file")
	flag.Parse()

	// Setup scene
	camPos = geometry.Vector{X: 0, Y: 0, Z: 0}
	fmt.Println("Setting up scene...")
	setupScene()

	fmt.Println(fmt.Sprintf("Rendering image to '%v'...", output))
	startTime := time.Now()
	data := image.NewNRGBA(image.Rect(0, 0, width, height))

	// Determine color of each pixel
	var wg sync.WaitGroup
	wg.Add(width * height)
	progressBar := pb.StartNew(width * height)
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			go calculatePixel(data, x, y, antialias, &wg, progressBar)
		}
	}
	wg.Wait() // Wait for all routines to finish
	progressBar.Finish()

	// Open the output file and write data to it
	img, err := os.OpenFile(output, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer img.Close()
	png.Encode(img, data)

	// Print off render time
	renderTime := time.Since(startTime).Seconds()
	fmt.Println("Render complete")
	fmt.Println(fmt.Sprintf("Elapsed time: %v seconds", renderTime))
}

func setupScene() {
	s := scene.Scene1{}
	entities = s.SetupScene()

	// Get lights and non lights
	for _, o := range entities {
		// if o.GetMaterial().Light {
		if _, ok := o.Material.(material.Light); ok {
			lights = append(lights, o)
		} else {
			entitiesNoLights = append(entitiesNoLights, o)
		}
	}

}

func calculatePixel(data *image.NRGBA, x int, y int, antialias int, wg *sync.WaitGroup, progressBar *pb.ProgressBar) {
	var cm material.ColorMixer

	// var c color.RGBA
	width := float64(width)
	height := float64(height)

	for i := 0; i < antialias; i++ {
		pixelPos := geometry.Vector{}
		pixelPos.Z = -2
		pixelPos.X = (float64(x)+rand.Float64())/width*2 - 1
		pixelPos.Y = -((float64(y)+rand.Float64())/height*2 - 1)
		c := shootRay(geometry.Ray{Origin: camPos, Direction: camPos.VectorTo(pixelPos).Normalize()})
		cm.Add(c)
	}

	data.Set(x, y, cm.ToNRGBA())
	progressBar.Increment()
	wg.Done()
}

func shootRay(ray geometry.Ray) color.NRGBA {
	closest := entity.ClosestCollision(ray, entities)
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
			var total_diffuse float64 = 0
			for _, light := range lights {
				// TODO: all lights are being treated as point lights right now
				ray_to_light := geometry.Ray{Origin: closest.Location, Direction: closest.Location.VectorTo(light.Entity.(entity.Sphere).Center)}.OffsetOrigin()
				if !entity.CollisionCloserThan(ray_to_light, entitiesNoLights, ray_to_light.Direction.Length()) {
					total_diffuse += math.Abs(ray_to_light.Direction.Dot(closest.Normal))
				}
			}
			total_diffuse /= float64(len(lights))
			total_diffuse = math.Max(0.2, math.Min(1, total_diffuse))

			c := closest.Renderable.Material.GetColor()
			c.R = uint8(total_diffuse * float64(c.R))
			c.G = uint8(total_diffuse * float64(c.G))
			c.B = uint8(total_diffuse * float64(c.B))
			return c
		}
	}
	return color.NRGBA{0, 0, 0, 255} // No hit, return black
}
