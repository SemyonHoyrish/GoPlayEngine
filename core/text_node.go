package core

import (
	"github.com/SemyonHoyrish/GoPlayEngine/primitive"
	"github.com/SemyonHoyrish/GoPlayEngine/resource"
)

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

func (t *TextNode) GetNodeType() NodeType { return TextNodeType }

func (t *TextNode) SetText(text string) {
	t.text = text
}

func (t *TextNode) GetText() string {
	return t.text
}

func (t *TextNode) SetTextSize(size int) {
	t.textSize = size
}

func (t *TextNode) GetTextSize() int {
	return t.textSize
}

func (t *TextNode) SetFont(font *resource.Font) {
	t.font = font
}

func (t *TextNode) GetFont() *resource.Font {
	return t.font
}

func (t *TextNode) SetColor(color primitive.Color) {
	t.color = color
}

func (t *TextNode) GetColor() primitive.Color {
	return t.color
}
