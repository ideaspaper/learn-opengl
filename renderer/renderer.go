package renderer

import (
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/ideaspaper/learn-opengl/entities"
	shaderprogram "github.com/ideaspaper/learn-opengl/shader_program"
)

type IRenderer interface {
	Prepare()
	Render(entities.IEntity, shaderprogram.IShaderProgram)
}

type renderer struct{}

func NewRenderer() IRenderer {
	err := gl.Init()
	if err != nil {
		panic(err)
	}
	return &renderer{}
}

func (r *renderer) Prepare() {
	gl.ClearColor(1, 0, 0, 0)
	gl.Clear(gl.COLOR_BUFFER_BIT)
}

func (r *renderer) Render(inputEntity entities.IEntity, inputShaderProgram shaderprogram.IShaderProgram) {
	model := inputEntity.Model()

	gl.BindVertexArray(model.Vao())
	gl.EnableVertexAttribArray(model.AttrListId()) // enable the attribute list where the data stored

	transformationMatrix := createTransformationMatrix(
		inputEntity.Coordinate(),
		inputEntity.Rotation(),
		inputEntity.Scale(),
	)

	projectionMatrix := createProjectionMatrix()

	inputShaderProgram.LoadMatrix("transformationMatrix", transformationMatrix)
	inputShaderProgram.LoadMatrix("projectionMatrix", projectionMatrix)

	gl.DrawElementsWithOffset(gl.TRIANGLES, model.VerticesCount(), gl.UNSIGNED_INT, 0)
	gl.DisableVertexAttribArray(model.AttrListId())
	gl.BindVertexArray(0)
}

func createTransformationMatrix(position, angles, scales []float32) mgl32.Mat4 {
	result := mgl32.Ident4()
	result = result.Mul4(mgl32.Translate3D(position[0], position[1], position[2]))
	result = result.Mul4(mgl32.HomogRotate3DX(angles[0]))
	result = result.Mul4(mgl32.HomogRotate3DX(angles[1]))
	result = result.Mul4(mgl32.HomogRotate3DX(angles[2]))
	result = result.Mul4(mgl32.Scale3D(scales[0], scales[1], scales[2]))
	return result
}

func createProjectionMatrix() mgl32.Mat4 {
	result := mgl32.Frustum(-2.0, 2.0, -2.0, 2.0, 1.0, 10.0)
	return result
}
