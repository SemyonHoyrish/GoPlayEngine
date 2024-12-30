package core

import (
	"fmt"
	"github.com/SemyonHoyrish/GoPlayEngine/basic"
	"github.com/SemyonHoyrish/GoPlayEngine/input"
)

// ComposedOverlap represent area, consisting of areas defined by other OverlapInterfaces, that can be overlapped with
// other OverlapInterfaces, or hovered by mouse.
// Consists of multiple Overlaps.
// Be considered, that overlaps are not affected by layers.
type ComposedOverlap struct {
	basic.Base

	node     *BaseNode
	overlaps []*Overlap
}

func NewComposedOverlap() *ComposedOverlap {
	return &ComposedOverlap{
		Base:     basic.MakeBase(),
		overlaps: make([]*Overlap, 0),
	}
}

// SetNode internal function that links node and overlap. Should NOT be called by user.
func (co *ComposedOverlap) SetNode(node *BaseNode) bool {
	if node == nil {
		if co.node == nil {
			fmt.Println(fmt.Errorf("trying to detach overlap, that already not attached (composed overlap id = %v)", co.GetID()))
		}
		co.node = nil
		return true
	}
	if co.node != nil {
		fmt.Println(fmt.Errorf("trying to attach overlap to node, but overlap already attached to node (composed overlap id = %d) (overlap attached to node id = %d) (new node id = %d)", co.GetID(), co.node.GetID(), node.GetID()))
		return false
	}
	co.node = node
	return true
}

// GetAbsolutePosition is an internal function.
func (co *ComposedOverlap) GetAbsolutePosition() basic.Point {
	// TODO: error if node is nil
	return co.node.GetAbsolutePosition()
}

// Add adds Overlap to composition.
func (co *ComposedOverlap) Add(overlap *Overlap) {
	co.overlaps = append(co.overlaps, overlap)
	overlap.SetComposedOverlap(co)
}

// OverlapsWith returns true if any of underlying Overlaps overlapping with a target.
// In case `other` also a ComposedOverlap, every pair is checked.
func (co *ComposedOverlap) OverlapsWith(other OverlapInterface) bool {
	if over, ok := other.(*Overlap); ok {
		for _, overlap := range co.overlaps {
			if over.OverlapsWith(overlap) {
				return true
			}
		}
		return false
	} else if compOver, ok := other.(*ComposedOverlap); ok {
		for _, curOver := range co.overlaps {
			for _, otherOver := range compOver.overlaps {
				if curOver.OverlapsWith(otherOver) {
					return true
				}
			}
		}
		return false
	} else {
		panic("Undefined overlap type (in composed)")
	}
}

// MouseOver return true if any of underlying Overlaps is hovered by Mouse.
func (co *ComposedOverlap) MouseOver(m *input.Mouse) bool {
	for _, over := range co.overlaps {
		if over.MouseOver(m) {
			return true
		}
	}
	return false
}
