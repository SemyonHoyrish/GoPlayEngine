package core

import (
	"github.com/SemyonHoyrish/GoPlayEngine/basic"
	"github.com/SemyonHoyrish/GoPlayEngine/data_structures"
	"github.com/SemyonHoyrish/GoPlayEngine/primitive"
)

// Scene implements scene of game engine.
// Scene have to be initialized with NewScene
type Scene struct {
	basic.Base

	nodes data_structures.Set[*Node]

	bgColor primitive.Color

	updateFunction func()
}

func NewScene() *Scene {
	return &Scene{
		Base:    basic.MakeBase(),
		nodes:   data_structures.CreateSet[*Node](),
		bgColor: primitive.Color{0, 0, 0, 255},
	}
}

// AddNode adds node to scene nodes
func (s *Scene) AddNode(node *Node) {
	s.nodes.Add(node)
}

// RemoveNode tries to find and remove node based on pointer equality, return true on success, false on fail
func (s *Scene) RemoveNode(node *Node) bool {
	return s.nodes.Remove(node)
}

// FindNode tries to find node based on node id (any node embeds Base, so have GetID() function),
// returns a pointer to the node on success, nil on fail
func (s *Scene) FindNode(id basic.IDType) *Node {
	for n := range s.nodes.Values() {
		if n.GetID() == id {
			return n
		}
	}

	return nil
}

// GetAllNodes returns all nodes attached to scene
func (s *Scene) GetAllNodes() []*Node {
	result := make([]*Node, 0, s.nodes.Len())
	for n := range s.nodes.Values() {
		result = append(result, n)
	}
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
