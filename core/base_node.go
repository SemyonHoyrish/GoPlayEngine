package core

import "github.com/SemyonHoyrish/GoPlayEngine/basic"

// LayerType determine order of rendering of objects (objects that renders later, renders on top of previous ones)
// 0 used for objects that should be rendered first, e.g. background object.
// Further increasing value (1, 2, 3 ...) allows to set expected order of rendering
type LayerType = uint32

// NodeType is an internal type
type NodeType uint32

// NodeType represents type of specific node, used in polymorph functions.
const (
	BaseNodeType NodeType = iota
	ObjectNodeType
	TextNodeType
	InputNodeType
)

// BaseNodeInterface is the highest abstraction level for any node that exists in game engine
type BaseNodeInterface interface {
	basic.BaseInterface

	GetNodeType() NodeType

	GetPosition() basic.Point
	SetPosition(basic.Point)

	GetAbsolutePosition() basic.Point

	GetSize() basic.Size
	SetSize(basic.Size)

	GetLayer() LayerType
	SetLayer(LayerType)

	GetParent() BaseNodeInterface
	SetParent(BaseNodeInterface)
}

// BaseNode is an implementation of BaseNodeInterface, which is embedded by any other node in engine.
// BaseNode have to be initialized with NewBaseNode()
type BaseNode struct {
	basic.Base

	position basic.Point
	size     basic.Size
	layer    LayerType

	parent BaseNodeInterface
}

func NewBaseNode() *BaseNode {
	return &BaseNode{
		Base: basic.MakeBase(),
	}
}

// GetNodeType used internally to determine which node passed to polymorph function
func (b *BaseNode) GetNodeType() NodeType { return BaseNodeType }

// GetPosition returns position relative to parent node, if that is set, to window space otherwise
func (b *BaseNode) GetPosition() basic.Point { return b.position }

// SetPosition sets position relative to parent center point, if that is set, to window space otherwise
func (b *BaseNode) SetPosition(p basic.Point) { b.position = p }

// GetAbsolutePosition returns position always relative to window space,
// invokes calculation of coordinates of all parents recursively
func (b *BaseNode) GetAbsolutePosition() basic.Point {
	if b.GetParent() != nil {
		return basic.Point{
			X: b.position.X + b.GetParent().GetAbsolutePosition().X,
			Y: b.position.Y + b.GetParent().GetAbsolutePosition().Y,
		}
	}
	return b.position
}

// GetSize returns size of node set by user
func (b *BaseNode) GetSize() basic.Size { return b.size }

// SetSize used to override size of node, e.g. size of texture for ObjectNode and size of rendered text for TextNode.
// Should be kept as zero, if other is not intended
func (b *BaseNode) SetSize(s basic.Size) { b.size = s }

// GetLayer returns layer of rendering
func (b *BaseNode) GetLayer() LayerType { return b.layer }

// SetLayer sets layer of rendering
func (b *BaseNode) SetLayer(l LayerType) { b.layer = l }

// GetParent return parent node of current node if parent is set, nil otherwise
func (b *BaseNode) GetParent() BaseNodeInterface { return b.parent }

// SetParent used internally, sets parent of node. User invocation of this function can lead to unexpected behavior.
func (b *BaseNode) SetParent(p BaseNodeInterface) { b.parent = p }
