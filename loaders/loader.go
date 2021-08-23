package loaders

import (
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/ideaspaper/learn-opengl/models"
)

type ILoader interface {
	LoadToVao(uint32, []float32, []uint32) model.IModel
	CleanUp()
}

type loader struct {
	vaos []uint32
	vbos []uint32
}

func NewLoader() ILoader {
	return &loader{}
}

func (l *loader) LoadToVao(attrListId uint32, vertices []float32, indicides []uint32) model.IModel {
	vao := l.createVao()

	// store vertices to VBO
	l.createVbo(gl.ARRAY_BUFFER)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(vertices), gl.Ptr(vertices), gl.STATIC_DRAW)

	// store indices to VBO
	l.createVbo(gl.ELEMENT_ARRAY_BUFFER)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, 4*len(indicides), gl.Ptr(indicides), gl.STATIC_DRAW)

	// store VBO to VAO
	gl.VertexAttribPointer(attrListId, 3, gl.FLOAT, false, 0, nil)

	gl.BindVertexArray(0)                     // unbind VAO
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)         // unbind the vertices VBO
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, 0) // unbind the indices VBO
	return model.NewModel(attrListId, vao, int32(len(indicides)))
}

func (l *loader) CleanUp() {
	for _, v := range l.vaos {
		gl.DeleteVertexArrays(1, &v)
	}
	for _, v := range l.vbos {
		gl.DeleteBuffers(1, &v)
	}
}

func (l *loader) createVbo(target uint32) {
	var vbo uint32 // create VBO
	gl.GenBuffers(1, &vbo)
	l.vbos = append(l.vbos, vbo)
	gl.BindBuffer(target, vbo) // bind the VBO
}

func (l *loader) createVao() uint32 {
	var vao uint32 // create VAO
	gl.GenVertexArrays(1, &vao)
	l.vaos = append(l.vaos, vao)
	gl.BindVertexArray(vao) // bind the VAO
	return vao
}
