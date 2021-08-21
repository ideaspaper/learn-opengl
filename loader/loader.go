package loader

import (
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/ideaspaper/learn-opengl/model"
)

type ILoader interface {
	LoadToVao([]float32) model.IModel
	createVao() uint32
	storeDataInAttrList(uint32, []float32)
	unbindVao()
	CleanUp()
}

type loader struct {
	Vaos []uint32
	Vbos []uint32
}

func NewLoader() ILoader {
	return &loader{}
}

func (l *loader) LoadToVao(data []float32) model.IModel {
	vao := l.createVao()
	l.storeDataInAttrList(0, data)
	l.unbindVao()
	return model.NewModel(vao, int32(len(data)/3))
}

func (l *loader) createVao() uint32 {
	// create VAO
	var vao uint32
	gl.GenVertexArrays(1, &vao)
	l.Vaos = append(l.Vaos, vao)

	// if we want to do anything with this VAO we have to bind it
	gl.BindVertexArray(vao)

	return vao
}

func (l *loader) storeDataInAttrList(attrListNumber uint32, data []float32) {
	// create VBO
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	l.Vbos = append(l.Vbos, vbo)

	// if we want to do anything with this VBO we have to bind it
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)

	// store data into VBO
	// gl.STATIC_DRAW data that will never be editted again
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(data), gl.Ptr(data), gl.STATIC_DRAW)

	// store VBO into one of the attribute lists of the VAO
	// attrListNumber = attributeList index
	// 3              = length of each vertex, 3 dimension
	// gl.Float       = the type of data
	// false          = data is not normalized
	// 0              = distance between each of your vertices
	// nil            = offset
	gl.VertexAttribPointer(attrListNumber, 3, gl.FLOAT, false, 0, nil)

	// unbind VBO
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
}

func (l *loader) unbindVao() {
	// unbind VAO
	gl.BindVertexArray(0)
}

func (l *loader) CleanUp() {
	for _, v := range l.Vaos {
		gl.DeleteVertexArrays(1, &v)
	}
	for _, v := range l.Vbos {
		gl.DeleteBuffers(1, &v)
	}
}
