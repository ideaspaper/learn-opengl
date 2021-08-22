package renderer

import (
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/ideaspaper/learn-opengl/model"
)

type IRenderer interface {
	Prepare()
	Render(model.IModel)
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

func (r *renderer) Render(model model.IModel) {
	gl.BindVertexArray(model.Vao())
	gl.EnableVertexAttribArray(model.AttribListId()) // enable the attribute list where the data stored
	gl.DrawElementsWithOffset(gl.TRIANGLES, model.VertexCount(), gl.UNSIGNED_INT, 0)
	gl.DisableVertexAttribArray(model.AttribListId())
	gl.BindVertexArray(0)
}
