package renderer

import (
	"math"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/ideaspaper/learn-opengl/entities"
	"github.com/ideaspaper/learn-opengl/helpers"
	shaderprogram "github.com/ideaspaper/learn-opengl/shader_program"
)

type IRenderer interface {
	Prepare()
	Render(entities.IEntity, shaderprogram.IShaderProgram)
}

type renderer struct {
	fovy       float32
	aspect     float32
	near       float32
	far        float32
	projection mgl32.Mat4
}

func NewRenderer(fovy, aspect, near, far float32) IRenderer {
	err := gl.Init()
	if err != nil {
		panic(err)
	}
	return &renderer{
		fovy:   fovy,
		aspect: aspect,
		near:   near,
		far:    far,
	}
}

func (r *renderer) Prepare() {
	gl.ClearColor(1, 0, 0, 0)
	gl.Clear(gl.COLOR_BUFFER_BIT)
}

func (r *renderer) Render(entity entities.IEntity, shaderProgram shaderprogram.IShaderProgram) {
	model := entity.Model()

	gl.BindVertexArray(model.Vao())
	gl.EnableVertexAttribArray(model.AttrListId()) // enable the attribute list where the data stored

	transformationMatrix := helpers.CreateTransformationMatrix(
		entity.Coordinate(),
		entity.Rotation(),
		entity.Scale(),
	)

	r.createProjectionMatrix()
	shaderProgram.LoadMatrix("transformationMatrix", transformationMatrix)
	shaderProgram.LoadMatrix("projectionMatrix", r.projection)

	gl.DrawElementsWithOffset(gl.TRIANGLES, model.VerticesCount(), gl.UNSIGNED_INT, 0)
	gl.DisableVertexAttribArray(model.AttrListId())
	gl.BindVertexArray(0)
}

func (r *renderer) createProjectionMatrix() {
	// create perspective using frustum
	top := r.near * float32(math.Tan(float64(r.fovy)*math.Pi/180/2))
	bottom := -top
	right := top * r.aspect
	left := -right
	r.projection = mgl32.Frustum(left, right, bottom, top, r.near, r.far)
}
