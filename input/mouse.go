package input

import (
	"github.com/SemyonHoyrish/GoPlayEngine/basic"
	"github.com/veandco/go-sdl2/sdl"
)

// Mouse represent mouse controller.
// It is required to use one instance of Mouse, which initialized by Engine (Engine.GetMouse).
type Mouse struct {
	// 0 - no last event, 1 - last event mouse up, 2 - last event mouse down
	// uint32 used instead of MouseButtonType because sdl mouse event provides some index which starts from 1 for some reason
	buttonLastEvent map[uint32]uint32
}

// NewMouse initialize new Mouse object, should be called only once (done inside Engine)
func NewMouse() *Mouse {
	return &Mouse{
		buttonLastEvent: make(map[uint32]uint32),
	}
}

// GetPosition return position of mouse relative to window space,
// position related to current position in event queue, may be different from os mouse position in case
// not all events have been processed up to the current time.
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

// ButtonPressed returns true if provided button is pressed and false otherwise.
func (m *Mouse) ButtonPressed(btn MouseButtonType) bool {
	_, _, state := sdl.GetMouseState()
	return (state>>uint32(btn))&1 == 1
}

// ButtonDown returns true if last event for provided button was mouse button down.
func (m *Mouse) ButtonDown(btn MouseButtonType) bool {
	btnInd := uint32(btn) + 1
	if m.buttonLastEvent[btnInd] == 2 {
		m.buttonLastEvent[btnInd] = 0
		return true
	}
	return false
}

// ButtonUp returns true if last event for provided button was mouse button up.
func (m *Mouse) ButtonUp(btn MouseButtonType) bool {
	btnInd := uint32(btn) + 1
	if m.buttonLastEvent[btnInd] == 1 {
		m.buttonLastEvent[btnInd] = 0
		return true
	}
	return false
}

// SetLastEvent is an internal function that used to keep track of last mouse button event
func (m *Mouse) SetLastEvent(e *sdl.MouseButtonEvent) {
	if e.Type == sdl.MOUSEBUTTONUP {
		m.buttonLastEvent[uint32(e.Button)] = 1
	} else if e.Type == sdl.MOUSEBUTTONDOWN {
		m.buttonLastEvent[uint32(e.Button)] = 2
	}
}
