package core

import (
	"github.com/SemyonHoyrish/GoPlayEngine/basic"
	"github.com/SemyonHoyrish/GoPlayEngine/primitive"
	"github.com/SemyonHoyrish/GoPlayEngine/resource"
)

type Texture struct {
	primitive primitive.PrimitiveInterface
	image     *resource.Image
}

func NewTextureFromImage(i *resource.Image) *Texture {
	return &Texture{image: i}
}
func NewTextureFromPrimitive(p primitive.PrimitiveInterface) *Texture {
	return &Texture{primitive: p}
}

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
		case primitive.EllipsePrimitive:
		case primitive.LinePrimitive:
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

func (t *Texture) GetPrimitive() primitive.PrimitiveInterface { return t.primitive }
func (t *Texture) GetImage() *resource.Image                  { return t.image }

// TODO: ?Move creating Texture to core.Texture instead of Engine.render
