package renderer

import (
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/ideaspaper/learn-opengl/model"
)

type IRenderer interface {
	Prepare()
	Render(model.IModel)
}

type renderer struct {
}

func NewRenderer() IRenderer {
	if err := gl.Init(); err != nil {
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
	gl.EnableVertexAttribArray(0) // enable the attribute list where the data stored
	gl.DrawElementsWithOffset(gl.TRIANGLES, model.VertexCount(), gl.UNSIGNED_INT, 0)
	gl.DisableVertexAttribArray(0)
	gl.BindVertexArray(0)
}
