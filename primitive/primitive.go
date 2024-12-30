package primitive

type PrimitiveType = uint32

const (
	RectanglePrimitive PrimitiveType = iota
	CirclePrimitive    PrimitiveType = iota
	EllipsePrimitive   PrimitiveType = iota
	LinePrimitive      PrimitiveType = iota
)

type PrimitiveInterface interface {
	GetPrimitiveType() PrimitiveType
	GetColor() Color
}
