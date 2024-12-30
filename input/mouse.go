package input

import (
	"github.com/SemyonHoyrish/GoPlayEngine/basic"
	"github.com/veandco/go-sdl2/sdl"
)

// Mouse represent mouse controller
type Mouse struct {
}

// GetPosition return position of mouse relative to window space,
// position related to current position in event queue, may be different from os mouse position in case
// not all events processed to current time.
func (m *Mouse) GetPosition() basic.Point {
	x, y, _ := sdl.GetMouseState()

	return basic.Point{X: float32(x), Y: float32(y)}
}

// MouseButtonType describes mouse buttons
type MouseButtonType uint32

// describes mouse buttons
const (
	MouseButtonLeft   MouseButtonType = iota
	MouseButtonMiddle MouseButtonType = iota
	MouseButtonRight  MouseButtonType = iota
)

// MouseButtonDown returns true if provided button is down and false otherwise.
func (m *Mouse) MouseButtonDown(btn MouseButtonType) bool {
	_, _, state := sdl.GetMouseState()
	return (state>>uint32(btn))&1 == 1
}
