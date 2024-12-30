package primitive

// PrimitiveType represents existing types of primitive objects of game engine
type PrimitiveType uint32

const (
	RectanglePrimitive PrimitiveType = iota
	CirclePrimitive    PrimitiveType = iota
	EllipsePrimitive   PrimitiveType = iota
	LinePrimitive      PrimitiveType = iota
)

// PrimitiveInterface describes primitive object
type PrimitiveInterface interface {
	GetPrimitiveType() PrimitiveType
	GetColor() Color
}
