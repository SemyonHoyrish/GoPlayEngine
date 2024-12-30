package primitive

import "github.com/SemyonHoyrish/GoPlayEngine/basic"

// Line implements line primitive
//
// This primitive has *unique aspect* of how node should be constructed.
// Generally, position of node determines center point of the node, but in case node texture created
// from the line, position of the node is start of the line, so it is required that position of node and
// `Line.Definition.From` is the same, otherwise you will get an error. End of the line will be `Line.Definition.To`.
type Line struct {
	Definition basic.Vector
	Color      Color
}

// GetPrimitiveType used internally, to differ primitives in polymorph functions.
func (r Line) GetPrimitiveType() PrimitiveType {
	return LinePrimitive
}

// GetColor returns color of the line.
func (r Line) GetColor() Color {
	return r.Color
}
