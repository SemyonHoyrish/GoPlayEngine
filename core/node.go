package core

import (
	"fmt"
	"github.com/SemyonHoyrish/GoPlayEngine/basic"
	"github.com/SemyonHoyrish/GoPlayEngine/data_structures"
	"github.com/SemyonHoyrish/GoPlayEngine/primitive"
	"github.com/SemyonHoyrish/GoPlayEngine/resource"
)

// NodeType is an internal type
type NodeType uint32

// LayerType determine order of rendering of objects (objects that renders later, renders on top of previous ones)
// 0 used for objects that should be rendered first, e.g. background object.
// Further increasing value (1, 2, 3 ...) allows to set expected order of rendering
type LayerType uint32

// NodeType represents type of specific node, used in polymorph functions.
const (
	NodeTypeBase = iota
	NodeTypeObject
	NodeTypeText
)

type Node struct {
	basic.Base

	nodeType NodeType

	position basic.Point
	size     basic.Size
	layer    LayerType

	parent   *Node
	children data_structures.Set[*Node]

	// object node
	texture *Texture

	// text node
	textInfo *NodeTextInfo

	overlap OverlapInterface

	// fields for auto overlap functionality
	autoOverlap        *ComposedOverlap
	autoOverlapEnabled bool
	autoOverlapBuilt   bool
	autoOverlapChild   bool
}

type NodeTextInfo struct {
	Text     string
	TextSize int
	Font     *resource.Font
	Color    primitive.Color
}

func NewNode() *Node {
	return &Node{
		Base:     basic.Base{},
		nodeType: NodeTypeBase,
		position: basic.Point{},
		size:     basic.Size{},
		parent:   nil,
		children: data_structures.CreateSet[*Node](),
		texture:  nil,
		textInfo: nil,
	}
}

func NewObjectNode(texture *Texture /* nillable */) *Node {
	return &Node{
		Base:     basic.MakeBase(),
		nodeType: NodeTypeObject,
		position: basic.Point{},
		size:     basic.Size{},
		parent:   nil,
		children: data_structures.CreateSet[*Node](),
		texture:  texture,
		textInfo: nil,
	}
}

func NewTextNode(textInfo *NodeTextInfo) *Node {
	return &Node{
		Base:     basic.MakeBase(),
		nodeType: NodeTypeText,
		position: basic.Point{},
		size:     basic.Size{},
		parent:   nil,
		children: data_structures.CreateSet[*Node](),
		texture:  nil,
		textInfo: textInfo,
	}
}

func (n *Node) GetType() NodeType {
	return n.nodeType
}

func (n *Node) GetPosition() basic.Point {
	return n.position
}

func (n *Node) SetPosition(position basic.Point) {
	n.position = position
}

func (n *Node) GetAbsolutePosition() basic.Point {
	parent := n.parent

	if parent != nil {
		return basic.Point{
			X: n.position.X + parent.GetAbsolutePosition().X,
			Y: n.position.Y + parent.GetAbsolutePosition().Y,
		}
	}

	return n.position
}

func (n *Node) GetOverrideSize() basic.Size {
	return n.size
}

func (n *Node) SetOverrideSize(size basic.Size) {
	n.size = size
}

func (n *Node) GetCalculatedSize() basic.Size {
	switch n.nodeType {
	case NodeTypeBase:
		return n.size
	case NodeTypeObject:
		size := n.size
		if size.Width == 0 && size.Height == 0 && n.texture != nil {
			size = n.texture.GetSize()
			if size.Width == 0 && size.Height == 0 {
				err := fmt.Errorf("size of node (id=%d) is zero still after resolution", n.GetID())
				fmt.Println(err)
			}
		}
		return size
	case NodeTypeText:
		textInfo := n.textInfo
		font := textInfo.Font
		ttfFont := font.GetTTFFont(textInfo.TextSize)
		size := n.size
		if size.Width == 0 && size.Height == 0 {
			w, h, _ := ttfFont.SizeUTF8(textInfo.Text)
			if w == 0 && h == 0 {
				err := fmt.Errorf("size of node (id=%d) is zero still after resolution", n.GetID())
				fmt.Println(err)
			} else {
				size.Width = float32(w)
				size.Height = float32(h)
			}
		}
		return size
	}

	panic("memory corruption???")
}

func (n *Node) GetLayer() LayerType {
	return n.layer
}

func (n *Node) SetLayer(layer LayerType) {
	n.layer = layer
}

func (n *Node) GetParent() *Node {
	return n.parent
}

func (n *Node) setParent(parent *Node) {
	n.parent = parent
}

func (n *Node) AddChild(child *Node) {
	n.children.Add(child)
	child.setParent(n)
}

func (n *Node) AddChildMany(child ...*Node) {
	for _, c := range child {
		n.children.Add(c)
		c.setParent(n)
	}
}

func (n *Node) RemoveChild(child *Node) bool {
	removed := n.children.Remove(child)
	if removed {
		child.setParent(nil)
	}

	return removed
}

func (n *Node) GetChildren() []*Node {
	val := make([]*Node, 0, n.children.Len())

	for child := range n.children.Values() {
		val = append(val, child)
	}

	return val
}

func (n *Node) SetOverlap(o OverlapInterface) {
	if n.autoOverlapChild == true {
		fmt.Println(fmt.Errorf("cannot set overlap when auto overlap enabled (node_id=%d)", n.GetID()))
		return
	}
	n.overlap = o
	n.overlap.SetNode(n)
}

func (n *Node) GetOverlap() OverlapInterface {
	if n.autoOverlap != nil {
		return n.autoOverlap
	}

	return n.overlap
}

func (n *Node) RemoveOverlap() {
	if n.autoOverlapChild == true {
		fmt.Println(fmt.Errorf("cannot remove overlap when overlap enabled (node_id=%d)", n.GetID()))
		return
	}
	if n.overlap != nil {
		n.overlap.SetNode(nil)
		n.overlap = nil
	}
}

func (n *Node) AutoOverlap(enabled bool) {
	n.autoOverlapEnabled = enabled
}

func (n *Node) AutoOverlapEnabled() bool {
	return n.autoOverlapEnabled
}

func (n *Node) GetAutoOverlap() *ComposedOverlap {
	return n.autoOverlap
}

func (n *Node) DestroyAutoOverlap() {
	for ch := range n.children.Values() {
		if ch.autoOverlapChild {
			ch.DestroyAutoOverlap()
		}
	}

	n.autoOverlapChild = false
	n.RemoveOverlap()
	n.autoOverlapBuilt = false
}

func (n *Node) BuildAutoOverlap(ignoreAutoOverlapSet bool, root *Node, rootOverlap *ComposedOverlap) *ComposedOverlap {
	if !n.autoOverlapEnabled && !ignoreAutoOverlapSet {
		fmt.Println(fmt.Errorf("cannot build auto overlap for node with auto overlap disabled (node_id=%d)", n.GetID()))
		return nil
	}

	if root == nil && rootOverlap != nil || root != nil && rootOverlap == nil {
		// TODO: report error
	}

	if rootOverlap == nil && root == nil && n.autoOverlapBuilt {
		n.DestroyAutoOverlap()
	}
	if n.autoOverlapChild && root == nil {
		fmt.Println(fmt.Errorf("cannot build auto overlap for root node (id=%d) because it is marked as a child of another auto overlap", n.GetID()))
		return nil
	} else if n.autoOverlapChild {
		fmt.Println(fmt.Errorf("cannot build auto overlap for node (id=%d) because node (id=%d) is a child of another auto overlap", root.GetID(), n.GetID()))
		return nil
	}

	isRootNode := false
	if rootOverlap == nil && root == nil {
		rootOverlap = NewComposedOverlap()
		root = n
		n.autoOverlapBuilt = true
		isRootNode = true
	}

	size := n.GetCalculatedSize()

	ov := NewOverlap(
		basic.Point{-size.Width / 2, -size.Height / 2},
		basic.Point{+size.Width / 2, +size.Height / 2},
	)

	if size != (basic.Size{0, 0}) {
		n.SetOverlap(ov)
		rootOverlap.Add(ov)
	}
	n.autoOverlapChild = true

	if isRootNode {
		n.autoOverlap = rootOverlap
		rootOverlap.SetNode(n)
	}

	for ch := range n.children.Values() {

		// for child nodes we ignore returned value
		_ = ch.BuildAutoOverlap(true, root, rootOverlap)
	}

	return rootOverlap
}

// --- object node ---

func (n *Node) GetTexture() *Texture {
	return n.texture
}

func (n *Node) SetTexture(texture *Texture) error {
	if n.nodeType != NodeTypeObject {
		return fmt.Errorf("cannot set texture on non-object node (id=%d)", n.GetID())
	}

	n.texture = texture
	return nil
}

// ---------------

// --- text node ---

func (n *Node) GetTextInfo() *NodeTextInfo {
	return n.textInfo
}

func (n *Node) SetTextInfo(textInfo *NodeTextInfo) error {
	if n.nodeType != NodeTypeText {
		return fmt.Errorf("cannot set texture on non-text node (id=%d)", n.GetID())
	}

	n.textInfo = textInfo
	return nil
}

// ---------------
