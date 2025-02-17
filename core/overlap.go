package core

import (
	"fmt"
	"github.com/SemyonHoyrish/GoPlayEngine/basic"
	"github.com/SemyonHoyrish/GoPlayEngine/input"
)

// Overlap rectangle representation of area that can be overlapped with other OverlapInterfaces, or hovered by Mouse.
// Overlap can be used as a completed thing to attach to BaseNode, or be a part of ComposedOverlap.
// Be considered, that overlaps are not affected by layers.
//
// In case this overlap attached to a node, position will be relative to that node position.
// In case this overlap is not attached to a node, but is part of ComposedOverlap, position will be calculated based on
// composeOverlap node position
// In other case this overlap will have coordinates (-1, -1) (-1, -1), and error will be reported (TODO)
type Overlap struct {
	basic.Base

	x1, x2, y1, y2 float32

	node            *Node
	composedOverlap *ComposedOverlap
}

// NewOverlap creates a new Overlap with coordinates relative to the further attached node position.
func NewOverlap(leftTop basic.Point, rightBottom basic.Point) *Overlap {
	return &Overlap{
		Base: basic.MakeBase(),

		x1: leftTop.X,
		x2: rightBottom.X,
		y1: leftTop.Y,
		y2: rightBottom.Y,
	}
}

// SetNode internal function that links node and overlap. Should NOT be called by user.
func (over *Overlap) SetNode(node *Node) bool {
	if node == nil {
		if over.node == nil {
			fmt.Println(fmt.Errorf("trying to detach overlap, that already not attached (overlap id = %v)", over.GetID()))
		}
		over.node = nil
		return true
	}
	if over.node != nil {
		fmt.Println(fmt.Errorf("trying to attach overlap to node, but overlap already attached to node (overlap id = %d) (overlap attached to node id = %d) (new node id = %d)", over.GetID(), over.node.GetID(), node.GetID()))
		return false
	}
	over.node = node
	return true
}

// SetComposedOverlap is an internal function. Should NOT be called by user.
func (over *Overlap) SetComposedOverlap(compOver *ComposedOverlap) {
	if compOver != nil && over.composedOverlap != nil {
		fmt.Println(fmt.Errorf("overriding composed overlap for node (node id = %d) (old compover id = %d) (new compover id = %d)", over.GetID(), over.composedOverlap.GetID(), compOver.GetID()))
	}
	over.composedOverlap = compOver
}

// GetAbsoluteValues is an internal function, which returns coordinates relative to world space, instead of node, Overlap attached to.
func (over *Overlap) GetAbsoluteValues() (float32, float32, float32, float32) {
	var pos basic.Point
	if over.node == nil && over.composedOverlap == nil {
		// TODO: print error
		return -1, -1, -1, -1
	} else if over.node == nil && over.composedOverlap != nil {
		pos = over.composedOverlap.GetAbsolutePosition()
	} else {
		pos = over.node.GetAbsolutePosition()
	}

	return over.x1 + pos.X, over.x2 + pos.X, over.y1 + pos.Y, over.y2 + pos.Y
}

// OverlapsWith returns true if this overlap has an intersection with `other`.
func (over *Overlap) OverlapsWith(other OverlapInterface) bool {
	if other == nil {
		fmt.Println(fmt.Errorf("OverlapsWith called on nil pointer (overlap_id=%d)", over.GetID()))
		return false
	}

	if compOver, ok := other.(*ComposedOverlap); ok {
		return compOver.OverlapsWith(over)
	}

	otherOver, ok := other.(*Overlap)
	if !ok {
		panic("Undefined overlap type")
	}

	x1, x2, y1, y2 := over.GetAbsoluteValues()
	otherX1, otherX2, otherY1, otherY2 := otherOver.GetAbsoluteValues()

	cond := (x1 < otherX2) &&
		(x2 > otherX1) &&
		(y1 < otherY2) &&
		(y2 > otherY1)

	return cond
}

// MouseOver returns true of mouse is over this Overlap
func (over *Overlap) MouseOver(m *input.Mouse) bool {
	x1, x2, y1, y2 := over.GetAbsoluteValues()
	mpos := m.GetPosition()

	return mpos.X >= x1 && mpos.X <= x2 && mpos.Y >= y1 && mpos.Y <= y2
}
