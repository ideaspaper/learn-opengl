package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/ideaspaper/learn-opengl/loader"
	"github.com/ideaspaper/learn-opengl/renderer"
	shaderprogram "github.com/ideaspaper/learn-opengl/shader_program"
)

var square = []float32{
	-0.5, 0.5, 0,
	-0.5, -0.5, 0,
	0.5, -0.5, 0,
	0.5, 0.5, 0,
}

var indicides = []uint32{0, 1, 3, 3, 1, 2}

func main() {
	runtime.LockOSThread()
	window := initGlfw()

	loader := loader.NewLoader()
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
		map[uint32]string{
			1: "position\x00",
		},
	)
	defer shaderProgram.CleanUp()

	model := loader.LoadToVao(1, square, indicides)

	for !window.ShouldClose() {
		renderer.Prepare()
		shaderProgram.Start()
		renderer.Render(model)
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
