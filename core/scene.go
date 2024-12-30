package core

import (
	"github.com/SemyonHoyrish/GoPlayEngine/basic"
	"github.com/SemyonHoyrish/GoPlayEngine/primitive"
)

type SceneInterface interface {
	basic.BaseInterface

	AddNode(node BaseNodeInterface)
	RemoveNode(node BaseNodeInterface) bool
	FindNode(id basic.IDType) (BaseNodeInterface, bool)
	GetAllNodes() []BaseNodeInterface

	SetBackgroundColor(color primitive.Color)
	GetBackgroundColor() primitive.Color
}

type Scene struct {
	basic.Base

	nodes []BaseNodeInterface

	bgColor primitive.Color
}

func NewScene() *Scene {
	return &Scene{
		Base:    basic.MakeBase(),
		nodes:   make([]BaseNodeInterface, 0),
		bgColor: primitive.Color{0, 0, 0, 255},
	}
}

func (s *Scene) AddNode(node BaseNodeInterface) {
	s.nodes = append(s.nodes, node)
}
func (s *Scene) RemoveNode(node BaseNodeInterface) bool {
	newNodes := make([]BaseNodeInterface, 0, len(s.nodes))
	found := false
	for _, n := range s.nodes {
		if n != node {
			newNodes = append(newNodes, n)
		} else {
			found = true
		}
	}

	if found {
		copy(s.nodes, newNodes)
	}

	return found
}

func (s *Scene) FindNode(id basic.IDType) (BaseNodeInterface, bool) {
	for _, n := range s.nodes {
		if n.GetID() == id {
			return n, true
		}
	}

	return nil, false
}

func (s *Scene) GetAllNodes() []BaseNodeInterface {
	result := make([]BaseNodeInterface, len(s.nodes))
	copy(result, s.nodes)
	return result
}

func (s *Scene) SetBackgroundColor(color primitive.Color) {
	s.bgColor = color
}

func (s *Scene) GetBackgroundColor() primitive.Color {
	return s.bgColor
}
