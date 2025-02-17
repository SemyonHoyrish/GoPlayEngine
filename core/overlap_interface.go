package core

import (
	"github.com/SemyonHoyrish/GoPlayEngine/basic"
	"github.com/SemyonHoyrish/GoPlayEngine/input"
)

type OverlapInterface interface {
	basic.BaseInterface

	OverlapsWith(OverlapInterface) bool
	MouseOver(*input.Mouse) bool
	SetNode(*Node) bool
}
