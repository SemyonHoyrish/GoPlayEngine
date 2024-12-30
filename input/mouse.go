package input

import (
	"github.com/SemyonHoyrish/GoPlayEngine/basic"
	"github.com/veandco/go-sdl2/sdl"
)

type Mouse struct {
}

func (m *Mouse) GetPosition() basic.Point {
	x, y, _ := sdl.GetMouseState()

	return basic.Point{X: float32(x), Y: float32(y)}
}

type MouseButtonType uint32

const (
	MouseButtonLeft   MouseButtonType = iota
	MouseButtonMiddle MouseButtonType = iota
	MouseButtonRight  MouseButtonType = iota
)

func (m *Mouse) MouseButtonDown(btn MouseButtonType) bool {
	_, _, state := sdl.GetMouseState()
	return (state>>uint32(btn))&1 == 1
}
