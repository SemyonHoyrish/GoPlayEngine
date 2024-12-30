package primitive

type Rectangle struct {
	Width  float32
	Height float32
	Color  Color
}

func (r Rectangle) GetPrimitiveType() PrimitiveType {
	return RectanglePrimitive
}

func (r Rectangle) GetColor() Color {
	return r.Color
}
