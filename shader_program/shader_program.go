package shaderprogram

import (
	"fmt"
	"os"
	"strings"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type IShaderProgram interface {
	Start()
	Stop()
	CleanUp()
	AttribVariable(string) uint32
	LoadFloat(string, float32)
	LoadMatrix(string, mgl32.Mat4)
}

type shaderProgram struct {
	program          uint32
	vertexShader     uint32
	fragmentShader   uint32
	attribVariables  map[string]uint32
	uniformVariables map[string]int32
}

func NewShaderProgram(
	vertexShaderFile,
	fragmentShaderFile string,
	attribVariables map[string]uint32,
	uniformVariables []string,
) IShaderProgram {
	vertexShader, err := loadShader(vertexShaderFile, gl.VERTEX_SHADER)
	if err != nil {
		panic(err)
	}

	fragmentShader, err := loadShader(fragmentShaderFile, gl.FRAGMENT_SHADER)
	if err != nil {
		panic(err)
	}

	program := gl.CreateProgram()

	gl.AttachShader(program, vertexShader)
	gl.AttachShader(program, fragmentShader)

	for k, v := range attribVariables {
		gl.BindAttribLocation(program, v, gl.Str(k+"\x00"))
	}

	gl.LinkProgram(program)
	gl.ValidateProgram(program)

	uniformVariablesMap := map[string]int32{}
	for _, v := range uniformVariables {
		uniformVariablesMap[v] = gl.GetUniformLocation(program, gl.Str(v+"\x00"))
	}

	return &shaderProgram{
		program:          program,
		vertexShader:     vertexShader,
		fragmentShader:   fragmentShader,
		attribVariables:  attribVariables,
		uniformVariables: uniformVariablesMap,
	}
}

func (sp *shaderProgram) Start() {
	gl.UseProgram(sp.program)
}

func (sp *shaderProgram) Stop() {
	gl.UseProgram(0)
}

func (sp *shaderProgram) CleanUp() {
	sp.Stop()
	gl.DetachShader(sp.program, sp.vertexShader)
	gl.DetachShader(sp.program, sp.fragmentShader)
	gl.DeleteShader(sp.vertexShader)
	gl.DeleteShader(sp.fragmentShader)
	gl.DeleteProgram(sp.program)
}

func (sp *shaderProgram) AttribVariable(variableName string) uint32 {
	return sp.attribVariables[variableName]
}

func (sp *shaderProgram) LoadFloat(uniformVarName string, value float32) {
	gl.Uniform1f(sp.uniformVariables[uniformVarName], value)
}

func (sp *shaderProgram) LoadMatrix(uniformVarName string, value mgl32.Mat4) {
	gl.UniformMatrix4fv(sp.uniformVariables[uniformVarName], 1, false, &value[0])
}

func loadShader(shaderFile string, shaderType uint32) (uint32, error) {
	shaderSourceBytes, err := os.ReadFile(shaderFile)
	if err != nil {
		return 0, fmt.Errorf("failed to open source file %v", shaderFile)
	}
	shaderSource := string(shaderSourceBytes)

	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(shaderSource)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", shaderSource, log)
	}

	return shader, nil
}
