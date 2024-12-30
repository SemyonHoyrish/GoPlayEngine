package core

import (
	"github.com/SemyonHoyrish/GoPlayEngine/primitive"
	"github.com/SemyonHoyrish/GoPlayEngine/resource"
)

// TextNodeInterface describes node, that renders text
type TextNodeInterface interface {
	BaseNodeInterface

	SetText(text string)
	GetText() string

	SetTextSize(size int)
	GetTextSize() int

	SetFont(font *resource.Font)
	GetFont() *resource.Font

	SetColor(color primitive.Color)
	GetColor() primitive.Color
}

// TextNode implements TextNodeInterface, cannot have child nodes.
// TextNode have to be initialized with NewTextNode()
type TextNode struct {
	BaseNode

	text     string
	textSize int
	font     *resource.Font
	color    primitive.Color
}

func NewTextNode() *TextNode {
	return &TextNode{
		BaseNode: *NewBaseNode(),
		text:     "",
		textSize: 0,
		font:     nil,
		color:    primitive.Color{0, 0, 0, 255},
	}
}

// GetNodeType used internally to determine which node passed to polymorph function
func (t *TextNode) GetNodeType() NodeType { return TextNodeType }

// SetText set text rendered by node
func (t *TextNode) SetText(text string) {
	t.text = text
}

// GetText returns text rendered by node
func (t *TextNode) GetText() string {
	return t.text
}

// SetTextSize sets size of rendering text
func (t *TextNode) SetTextSize(size int) {
	t.textSize = size
}

// GetTextSize returns size of rendering text
func (t *TextNode) GetTextSize() int {
	return t.textSize
}

// SetFont sets font to render text with
func (t *TextNode) SetFont(font *resource.Font) {
	t.font = font
}

// GetFont returns font set to render text, nil if none was set
func (t *TextNode) GetFont() *resource.Font {
	return t.font
}

// SetColor sets color of text
func (t *TextNode) SetColor(color primitive.Color) {
	t.color = color
}

// GetColor returns color of text
func (t *TextNode) GetColor() primitive.Color {
	return t.color
}
