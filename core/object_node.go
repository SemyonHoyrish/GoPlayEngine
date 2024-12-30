package core

type ObjectNodeInterface interface {
	BaseNodeInterface

	GetChildren() []BaseNodeInterface
	AddChild(child BaseNodeInterface)
	RemoveChild(child BaseNodeInterface) bool

	GetTexture() *Texture
	SetTexture(texture *Texture)
}

type ObjectNode struct {
	BaseNode

	texture  *Texture
	children []BaseNodeInterface
}

func NewObjectNode() *ObjectNode {
	return &ObjectNode{
		BaseNode: *NewBaseNode(),
		texture:  nil,
		children: make([]BaseNodeInterface, 0),
	}
}

func (n *ObjectNode) GetNodeType() NodeType { return ObjectNodeType }
func (n *ObjectNode) GetChildren() []BaseNodeInterface {
	return n.children
}

func (n *ObjectNode) AddChild(child BaseNodeInterface) {
	n.children = append(n.children, child)
	child.SetParent(n)
}

func (n *ObjectNode) RemoveChild(child BaseNodeInterface) bool {
	newNodes := make([]BaseNodeInterface, 0, len(n.children))
	found := false
	for _, c := range n.children {
		if c != child {
			newNodes = append(newNodes, c)
		} else {
			found = true
		}
	}

	if found {
		copy(n.children, newNodes)
	}

	return found
}

func (n *ObjectNode) GetTexture() *Texture {
	return n.texture
}

func (n *ObjectNode) SetTexture(texture *Texture) {
	n.texture = texture
}
