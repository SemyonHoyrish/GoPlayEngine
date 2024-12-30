package core

// ObjectNodeInterface describes object node, most common node of game engine,
// which used to render primitives or textures, or to group multiple nodes.
type ObjectNodeInterface interface {
	BaseNodeInterface

	GetChildren() []BaseNodeInterface
	AddChild(child BaseNodeInterface)
	RemoveChild(child BaseNodeInterface) bool

	GetTexture() *Texture
	SetTexture(texture *Texture)
}

// ObjectNode is an implementation of ObjectNodeInterface
// ObjectNode have to be initialized with NewObjectNode()
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

// GetNodeType used internally to determine which node passed to polymorph function
func (n *ObjectNode) GetNodeType() NodeType { return ObjectNodeType }

// GetChildren returns slice of all children attached to this node
func (n *ObjectNode) GetChildren() []BaseNodeInterface {
	return n.children
}

// AddChild adds child to this node
func (n *ObjectNode) AddChild(child BaseNodeInterface) {
	n.children = append(n.children, child)
	child.SetParent(n)
}

// RemoveChild tries to find and remove child based on pointer equality,
// returns true in success and false if child was not found
func (n *ObjectNode) RemoveChild(child BaseNodeInterface) bool {
	newNodes := make([]BaseNodeInterface, 0, len(n.children))
	found := false
	for _, c := range n.children {
		if c != child {
			newNodes = append(newNodes, c)
		} else {
			found = true
			c.SetParent(nil)
		}
	}

	if found {
		copy(n.children, newNodes)
	}

	return found
}

// GetTexture returns set texture or nil if no texture wa provided to node
func (n *ObjectNode) GetTexture() *Texture {
	return n.texture
}

// SetTexture sets texture of the node
func (n *ObjectNode) SetTexture(texture *Texture) {
	n.texture = texture
}
