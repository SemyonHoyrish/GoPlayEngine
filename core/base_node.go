package core

import "github.com/SemyonHoyrish/GoPlayEngine/basic"

type LayerType = uint32

type NodeType = uint32

const (
	BaseNodeType NodeType = iota
	ObjectNodeType
	TextNodeType
	InputNodeType
)

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

func (b *BaseNode) GetNodeType() NodeType { return BaseNodeType }

func (b *BaseNode) GetPosition() basic.Point  { return b.position }
func (b *BaseNode) SetPosition(p basic.Point) { b.position = p }

func (b *BaseNode) GetAbsolutePosition() basic.Point {
	if b.GetParent() != nil {
		return basic.Point{
			X: b.position.X + b.GetParent().GetAbsolutePosition().X,
			Y: b.position.Y + b.GetParent().GetAbsolutePosition().Y,
		}
	}
	return b.position
}

func (b *BaseNode) GetSize() basic.Size  { return b.size }
func (b *BaseNode) SetSize(s basic.Size) { b.size = s }

func (b *BaseNode) GetLayer() LayerType  { return b.layer }
func (b *BaseNode) SetLayer(l LayerType) { b.layer = l }

func (b *BaseNode) GetParent() BaseNodeInterface { return b.parent }

func (b *BaseNode) SetParent(p BaseNodeInterface) { b.parent = p }
