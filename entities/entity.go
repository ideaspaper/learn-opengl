package entities

import "github.com/ideaspaper/learn-opengl/models"

type IEntity interface {
	Model() model.IModel
	Coordinate() []float32
	Rotation() []float32
	Scale() []float32
}

type entity struct {
	model      model.IModel
	coordinate []float32
	rotation   []float32
	scale      []float32
}

func NewEntity(model model.IModel, coordinate, rotation, scale []float32) IEntity {
	return &entity{
		model:      model,
		coordinate: coordinate,
		rotation:   rotation,
		scale:      scale,
	}
}

func (e *entity) Model() model.IModel {
	return e.model
}

func (e *entity) Coordinate() []float32 {
	return e.coordinate
}

func (e *entity) SetCoordinate(coordinate []float32) {
	e.coordinate = coordinate
}

func (e *entity) Rotation() []float32 {
	return e.rotation
}

func (e *entity) SetRotation(rotation []float32) {
	e.rotation = rotation
}

func (e *entity) Scale() []float32 {
	return e.scale
}

func (e *entity) SetScale(newScales []float32) {
	e.scale = newScales
}
