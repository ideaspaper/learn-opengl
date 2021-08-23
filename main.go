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

// square
var vertices = []float32{
	-0.5, 0.5, 0,
	-0.5, -0.5, 0,
	0.5, -0.5, 0,
	0.5, 0.5, 0,
}

var indicides = []uint32{0, 1, 3, 3, 1, 2}

func main() {
	runtime.LockOSThread()
	window := initGlfw()

	loader := loaders.NewLoader()
	defer loader.CleanUp()

	renderer := renderer.NewRenderer()

	ex, err := os.Executable()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	shaderPath := filepath.Join(filepath.Dir(ex), "shaders")
	vertexShaderFile := filepath.Join(shaderPath, "vertex_shader.vert")
	fragmentShaderFile := filepath.Join(shaderPath, "fragment_shader.frag")
	shaderProgram := shaderprogram.NewShaderProgram(
		vertexShaderFile,
		fragmentShaderFile,
		map[string]uint32{
			"position": 0,
		},
		[]string{
			"transformationMatrix",
			"projectionMatrix",
		},
	)
	defer shaderProgram.CleanUp()

	model := loader.LoadToVao(shaderProgram.AttribVariable("position"), vertices, indicides)
	entity := entities.NewEntity(model, []float32{0.0, 0.0, -5.0}, []float32{0.0, 0.0, 0.0}, []float32{1.0, 1.0, 1.0})

	togglePositions := false
	toggleScales := false

	for !window.ShouldClose() {
		renderer.Prepare()
		shaderProgram.Start()

		if entity.Coordinate()[0] > 1 || entity.Coordinate()[0] < -1 {
			togglePositions = !togglePositions
		}

		if !togglePositions {
			entity.Coordinate()[0] += 0.0001
		} else {
			entity.Coordinate()[0] -= 0.0001
		}

		if entity.Scale()[0] > 1.2 || entity.Coordinate()[0] < 0.8 {
			toggleScales = !toggleScales
		}

		if !toggleScales {
			entity.Scale()[0] += 0.001
		} else {
			entity.Scale()[0] -= 0.001
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
	window, err := glfw.CreateWindow(640, 480, "First Window", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()
	return window
}
