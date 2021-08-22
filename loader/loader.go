package loader

import (
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/ideaspaper/learn-opengl/model"
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

func (l *loader) LoadToVao(attribListId uint32, positions []float32, indicides []uint32) model.IModel {
	vao := l.createVao()
	l.bindIndicesBuffer(indicides)
	l.storeDataInAttrList(attribListId, positions)
	l.unbindVao()
	return model.NewModel(attribListId, vao, int32(len(indicides)))
}

func (l *loader) CleanUp() {
	for _, v := range l.vaos {
		gl.DeleteVertexArrays(1, &v)
	}
	for _, v := range l.vbos {
		gl.DeleteBuffers(1, &v)
	}
}

func (l *loader) createVao() uint32 {
	// create VAO
	var vao uint32
	gl.GenVertexArrays(1, &vao)
	l.vaos = append(l.vaos, vao)

	// if we want to do anything with this VAO we have to bind it
	gl.BindVertexArray(vao)

	return vao
}

func (l *loader) storeDataInAttrList(attrListNumber uint32, data []float32) {
	// create VBO
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	l.vbos = append(l.vbos, vbo)

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

func (l *loader) bindIndicesBuffer(indicides []uint32) {
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	l.vbos = append(l.vbos, vbo)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, 4*len(indicides), gl.Ptr(indicides), gl.STATIC_DRAW)
}

func (l *loader) unbindVao() {
	// unbind VAO
	gl.BindVertexArray(0)
}
