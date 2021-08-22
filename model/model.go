package model

type IModel interface {
	Vao() uint32
	VertexCount() int32
	AttribListId() uint32
}

type model struct {
	attribListId uint32
	vao          uint32
	vertexCount  int32
}

func NewModel(attribListId, vao uint32, vertexCount int32) IModel {
	return &model{
		attribListId: attribListId,
		vao:          vao,
		vertexCount:  vertexCount,
	}
}

func (m *model) AttribListId() uint32 {
	return m.attribListId
}

func (m *model) Vao() uint32 {
	return m.vao
}

func (m *model) VertexCount() int32 {
	return m.vertexCount
}
