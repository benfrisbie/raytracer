package main

import (
	"flag"
	"fmt"
	"image/png"
	"log"
	"os"
	"time"

	"github.com/benfrisbie/raytracer/camera"
	"github.com/benfrisbie/raytracer/scene"
)

func main() {
	var width int
	var height int
	var antialias int
	var fov float64
	var output string
	flag.IntVar(&width, "width", 1920, "width of the image")
	flag.IntVar(&height, "height", 1080, "height of the image")
	flag.Float64Var(&fov, "fov", 90, "field of view of the camera")
	flag.IntVar(&antialias, "antialias", 4, "antialias sets how many rays per pixel should be shot into the scene. increasing this improves quality around edges and the expense of performance")
	flag.StringVar(&output, "output", "test.png", "name of the output png file")
	flag.Parse()

	// Setup camera
	fmt.Println("Setting up camera...")
	cam := camera.NewCamera(width, height, antialias, fov)

	// Setup scene
	fmt.Println("Setting up scene...")
	sce := scene.NewScene()

	// Have the camera render the scene
	fmt.Println(fmt.Sprintf("Rendering image to '%v'...", output))
	startTime := time.Now()
	data := cam.RenderSceneToImage(sce)

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
