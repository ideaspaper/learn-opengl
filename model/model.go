package model

type IModel interface {
	Vao() uint32
	VertexCount() int32
}

type model struct {
	vao         uint32
	vertexCount int32
}

func NewModel(vao uint32, vertexCount int32) IModel {
	return &model{
		vao:         vao,
		vertexCount: vertexCount,
	}
}

func (m *model) Vao() uint32 {
	return m.vao
}

func (m *model) VertexCount() int32 {
	return m.vertexCount
}
