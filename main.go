package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/ideaspaper/learn-opengl/entities"
	"github.com/ideaspaper/learn-opengl/loaders"
	"github.com/ideaspaper/learn-opengl/renderer"
	shaderprogram "github.com/ideaspaper/learn-opengl/shader_program"
)

const (
	windowWidth  int = 640
	windowHeight int = 480
)

// square vertices
var vertices = []float32{
	-0.5, 0.5, 0,
	-0.5, -0.5, 0,
	0.5, -0.5, 0,
	0.5, 0.5, 0,
}

// square indices
var indicides = []uint32{0, 1, 3, 3, 1, 2}

func main() {
	runtime.LockOSThread()

	// initialize window
	window := initGlfw()

	// create a new loader
	loader := loaders.NewLoader()
	defer loader.CleanUp()

	// create a new renderer
	renderer := renderer.NewRenderer(70, float32(windowWidth/windowHeight), 0.1, 1000.0)

	// create a new shader program
	ex, err := os.Executable()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	shaderPath := filepath.Join(filepath.Dir(ex), "shaders")
	vertexShaderFile := filepath.Join(shaderPath, "vertex_shader.vert")
	fragmentShaderFile := filepath.Join(shaderPath, "fragment_shader.frag")
	attribVariables := map[string]uint32{
		"position": 0,
	}
	uniformVariables := []string{
		"transformationMatrix",
		"projectionMatrix",
	}
	shaderProgram := shaderprogram.NewShaderProgram(
		vertexShaderFile,
		fragmentShaderFile,
		attribVariables,
		uniformVariables,
	)
	defer shaderProgram.CleanUp()

	// store vertices and indices to VAO, then create a model
	model := loader.LoadToVao(shaderProgram.AttribVariable("position"), vertices, indicides)

	// create an entity based on model
	coordinate := []float32{0.0, 0.0, -2.0}
	rotation := []float32{0.0, 0.0, 0.0}
	scale := []float32{1.0, 1.0, 1.0}
	entity := entities.NewEntity(model, coordinate, rotation, scale)

	togglePositions := false

	for !window.ShouldClose() {
		renderer.Prepare()
		shaderProgram.Start()

		if entity.Coordinate()[2] > -1 || entity.Coordinate()[2] < -3 {
			togglePositions = !togglePositions
		}

		if !togglePositions {
			entity.Coordinate()[2] += 0.00025
		} else {
			entity.Coordinate()[2] -= 0.00025
		}

		renderer.Render(entity, shaderProgram)
		shaderProgram.Stop()
		glfw.PollEvents()
		window.SwapBuffers()
	}
}

func initGlfw() *glfw.Window {
	if err := glfw.Init(); err != nil {
		panic(err)
	}
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 6)
	window, err := glfw.CreateWindow(windowWidth, windowHeight, "First Window", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()
	return window
}
