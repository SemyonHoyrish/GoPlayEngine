package core

import (
	"github.com/SemyonHoyrish/GoPlayEngine/basic"
	"github.com/SemyonHoyrish/GoPlayEngine/primitive"
	"github.com/SemyonHoyrish/GoPlayEngine/resource"
)

// Texture represent source to create actual texture, which will be rendered.
type Texture struct {
	primitive primitive.PrimitiveInterface
	image     *resource.Image
}

// NewTextureFromImage creates texture source based on Image resource
func NewTextureFromImage(i *resource.Image) *Texture {
	return &Texture{image: i}
}

// NewTextureFromPrimitive creates texture source based on primitive
func NewTextureFromPrimitive(p primitive.PrimitiveInterface) *Texture {
	return &Texture{primitive: p}
}

// GetSize calculates size based on source provided on creation,
// returned value describes size needed to render texture as it was intended,
// can be overridden by size of node.
func (t *Texture) GetSize() basic.Size {
	if t.primitive != nil {
		// TODO: implement all cases
		switch t.primitive.GetPrimitiveType() {
		case primitive.RectanglePrimitive:
			return basic.Size{
				Width:  t.primitive.(primitive.Rectangle).Width,
				Height: t.primitive.(primitive.Rectangle).Height,
			}
		case primitive.CirclePrimitive:
			return basic.Size{
				Width:  t.primitive.(primitive.Circle).Radius,
				Height: t.primitive.(primitive.Circle).Radius,
			}
		case primitive.EllipsePrimitive:
		case primitive.LinePrimitive:
			//vec := t.primitive.(primitive.Line).Definition
			//return basic.Size{
			//	Width:  float32(math.Abs(float64(vec.From.X - vec.To.X))),
			//	Height: float32(math.Abs(float64(vec.From.Y - vec.To.Y))),
			//}

			to := t.primitive.(primitive.Line).To
			return basic.Size{
				Width:  to.X,
				Height: to.Y,
			}
		}

		panic("to implement")
	} else {
		if surf := t.image.GetSurface(); surf != nil {
			return basic.Size{Height: float32(surf.H), Width: float32(surf.W)}
		} else {
			return basic.Size{}
		}
	}
}

// GetPrimitive returns primitive on which this texture was created, nil if was created on Image.
func (t *Texture) GetPrimitive() primitive.PrimitiveInterface { return t.primitive }

// GetImage returns Image in which this texture was created, nil if was created on primitive.
func (t *Texture) GetImage() *resource.Image { return t.image }

// TODO: ?Move creating Texture to core.Texture instead of Engine.render
