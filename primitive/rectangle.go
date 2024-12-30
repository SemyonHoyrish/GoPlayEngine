package primitive

// Rectangle implements rectangle primitive
type Rectangle struct {
	Width  float32
	Height float32
	Color  Color
}

// GetPrimitiveType used internally, to differ primitives in polymorph functions.
func (r Rectangle) GetPrimitiveType() PrimitiveType {
	return RectanglePrimitive
}

// GetColor returns color of rectangle.
func (r Rectangle) GetColor() Color {
	return r.Color
}
