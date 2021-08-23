package model

type IModel interface {
	Vao() uint32
	VerticesCount() int32
	AttrListId() uint32
}

type model struct {
	attrListId    uint32
	vao           uint32
	verticesCount int32
}

func NewModel(attrListId, vao uint32, verticesCount int32) IModel {
	return &model{
		attrListId:    attrListId,
		vao:           vao,
		verticesCount: verticesCount,
	}
}

func (m *model) AttrListId() uint32 {
	return m.attrListId
}

func (m *model) Vao() uint32 {
	return m.vao
}

func (m *model) VerticesCount() int32 {
	return m.verticesCount
}
