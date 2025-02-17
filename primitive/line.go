package primitive

import "github.com/SemyonHoyrish/GoPlayEngine/basic"

// Line implements line primitive
//
// `Line.To` is the point line will be drawn to, from the node position, for which line primitive were used to create a texture.
// Nodes with line primitive as a texture will be skipped when creating an auto overlap.
type Line struct {
	To    basic.Point
	Color Color
}

// GetPrimitiveType used internally, to differ primitives in polymorph functions.
func (r Line) GetPrimitiveType() PrimitiveType {
	return LinePrimitive
}

// GetColor returns color of the line.
func (r Line) GetColor() Color {
	return r.Color
}
