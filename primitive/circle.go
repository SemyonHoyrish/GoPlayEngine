package primitive

// Circle implements circle primitive
type Circle struct {
	Radius float32
	Color  Color
}

// GetPrimitiveType used internally, to differ primitives in polymorph functions.
func (c Circle) GetPrimitiveType() PrimitiveType {
	return CirclePrimitive
}

// GetColor returns color of the circle.
func (c Circle) GetColor() Color {
	return c.Color
}
