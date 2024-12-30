package core

import (
	"github.com/SemyonHoyrish/GoPlayEngine/basic"
	"github.com/SemyonHoyrish/GoPlayEngine/primitive"
)

// SceneInterface describes scene of game engine, that used to provide nodes to render, and update function.
// update function called before render of every frame
type SceneInterface interface {
	basic.BaseInterface

	AddNode(node BaseNodeInterface)
	RemoveNode(node BaseNodeInterface) bool
	FindNode(id basic.IDType) (BaseNodeInterface, bool)
	GetAllNodes() []BaseNodeInterface

	SetBackgroundColor(color primitive.Color)
	GetBackgroundColor() primitive.Color

	SetUpdateFunction(func())
	GetUpdateFunction() func()
}

// Scene implements scene of game engine.
// Scene have to be initialized with NewScene
type Scene struct {
	basic.Base

	nodes []BaseNodeInterface

	bgColor primitive.Color

	updateFunction func()
}

func NewScene() *Scene {
	return &Scene{
		Base:    basic.MakeBase(),
		nodes:   make([]BaseNodeInterface, 0),
		bgColor: primitive.Color{0, 0, 0, 255},
	}
}

// AddNode adds node to scene nodes
func (s *Scene) AddNode(node BaseNodeInterface) {
	s.nodes = append(s.nodes, node)
}

// RemoveNode tries to find and remove node based on pointer equality, return true on success, false on fail
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

// FindNode tries to find node based on node id (any node embeds Base, so have GetID() function),
// returns (node, true) on success, (nil, false) on fail
func (s *Scene) FindNode(id basic.IDType) (BaseNodeInterface, bool) {
	for _, n := range s.nodes {
		if n.GetID() == id {
			return n, true
		}
	}

	return nil, false
}

// GetAllNodes returns all nodes attached to scene
func (s *Scene) GetAllNodes() []BaseNodeInterface {
	result := make([]BaseNodeInterface, len(s.nodes))
	copy(result, s.nodes)
	return result
}

// SetBackgroundColor sets color which will be used to clear screen before render of every frame,
// Color{0, 0, 0, 255} (black) by default
func (s *Scene) SetBackgroundColor(color primitive.Color) {
	s.bgColor = color
}

// GetBackgroundColor returns color which is used to clear screen before render of every frame
func (s *Scene) GetBackgroundColor() primitive.Color {
	return s.bgColor
}

// SetUpdateFunction sets function that called before rendering every frame of this scene.
func (s *Scene) SetUpdateFunction(updateFunction func()) {
	s.updateFunction = updateFunction
}

// GetUpdateFunction returns function that called before rendering every frame of this scene.
// Used internally
func (s *Scene) GetUpdateFunction() func() {
	return s.updateFunction
}
