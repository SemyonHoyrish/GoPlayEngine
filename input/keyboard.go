package input

import "github.com/veandco/go-sdl2/sdl"

// Keyboard represent keyboard controller.
// It is required to use one instance of Keyboard, which initialized by Engine (Engine.GetKeyboard).
type Keyboard struct {
	state []uint8

	// 0 - no last event, 1 - last event mouse up, 2 - last event mouse down,
	// 3 - previous last event was mouse up, 4 - previous last event was mouse down
	buttonLastEvent map[Scancode]uint32

	deferredChanges map[Scancode]uint32
}

// NewKeyboard initialize new Keyboard object, should be called only once (done inside Engine)
func NewKeyboard() *Keyboard {
	return &Keyboard{
		state:           sdl.GetKeyboardState(),
		buttonLastEvent: make(map[Scancode]uint32),
		deferredChanges: make(map[Scancode]uint32),
	}
}

// ButtonPressed returns true if provided button is pressed and false otherwise.
func (k *Keyboard) ButtonPressed(btn Scancode) bool {
	return k.state[btn] == 1
}

// ButtonDown returns true if last event for provided button was key down.
func (k *Keyboard) ButtonDown(btn Scancode) bool {
	if k.buttonLastEvent[btn] == 2 {
		k.deferredChanges[btn] = 4
		return true
	}
	return false
}

// ButtonUp returns true if last event for provided button was key up.
func (k *Keyboard) ButtonUp(btn Scancode) bool {
	if k.buttonLastEvent[btn] == 1 {
		k.deferredChanges[btn] = 3
		return true
	}
	return false
}

// SetLastEvent is an internal function that used to keep track of last keyboard button event for each button
func (k *Keyboard) SetLastEvent(e *sdl.KeyboardEvent) {
	// TODO: add some timeout after one state will be cleared, in case some event waw ignored, lost, etc.
	if e.Type == sdl.KEYUP && k.buttonLastEvent[e.Keysym.Scancode] != 3 {
		k.buttonLastEvent[e.Keysym.Scancode] = 1
	} else if e.Type == sdl.KEYDOWN && k.buttonLastEvent[e.Keysym.Scancode] != 4 {
		k.buttonLastEvent[e.Keysym.Scancode] = 2
	}
}

// ApplyDeferred is an internal function.
// Function used to set proper values for key down/up event, does not affect key pressed.
func (k *Keyboard) ApplyDeferred() {
	for btn, val := range k.deferredChanges {
		k.buttonLastEvent[btn] = val
	}
	k.deferredChanges = make(map[Scancode]uint32)
}

// GetButtonName returns human-readable name for provided scancode
func (k *Keyboard) GetButtonName(scancode Scancode) string {
	return sdl.GetScancodeName(scancode)
}

// GetButtonNameNonQWERTY returns human-readable name for provided scancode, for non-QWERTY keyboard (hopefully). (not tested)
func (k *Keyboard) GetButtonNameNonQWERTY(scancode Scancode) string {
	return sdl.GetKeyName(sdl.GetKeyFromScancode(scancode))
}

// Scancode is an alias made for consistency
//
// There are rebindings for sdl scancodes that I found useful in game engine.
// If you need some other key that usb keyboard can provide, look at sdl.SCANCODE_*, input.Keyboard supports all of them.
type Scancode = sdl.Scancode // uint32

const (
	SCANCODE_UNKNOWN = sdl.SCANCODE_UNKNOWN // "" (no name, empty string)

	SCANCODE_A = sdl.SCANCODE_A // "A"
	SCANCODE_B = sdl.SCANCODE_B // "B"
	SCANCODE_C = sdl.SCANCODE_C // "C"
	SCANCODE_D = sdl.SCANCODE_D // "D"
	SCANCODE_E = sdl.SCANCODE_E // "E"
	SCANCODE_F = sdl.SCANCODE_F // "F"
	SCANCODE_G = sdl.SCANCODE_G // "G"
	SCANCODE_H = sdl.SCANCODE_H // "H"
	SCANCODE_I = sdl.SCANCODE_I // "I"
	SCANCODE_J = sdl.SCANCODE_J // "J"
	SCANCODE_K = sdl.SCANCODE_K // "K"
	SCANCODE_L = sdl.SCANCODE_L // "L"
	SCANCODE_M = sdl.SCANCODE_M // "M"
	SCANCODE_N = sdl.SCANCODE_N // "N"
	SCANCODE_O = sdl.SCANCODE_O // "O"
	SCANCODE_P = sdl.SCANCODE_P // "P"
	SCANCODE_Q = sdl.SCANCODE_Q // "Q"
	SCANCODE_R = sdl.SCANCODE_R // "R"
	SCANCODE_S = sdl.SCANCODE_S // "S"
	SCANCODE_T = sdl.SCANCODE_T // "T"
	SCANCODE_U = sdl.SCANCODE_U // "U"
	SCANCODE_V = sdl.SCANCODE_V // "V"
	SCANCODE_W = sdl.SCANCODE_W // "W"
	SCANCODE_X = sdl.SCANCODE_X // "X"
	SCANCODE_Y = sdl.SCANCODE_Y // "Y"
	SCANCODE_Z = sdl.SCANCODE_Z // "Z"

	SCANCODE_1 = sdl.SCANCODE_1 // "1"
	SCANCODE_2 = sdl.SCANCODE_2 // "2"
	SCANCODE_3 = sdl.SCANCODE_3 // "3"
	SCANCODE_4 = sdl.SCANCODE_4 // "4"
	SCANCODE_5 = sdl.SCANCODE_5 // "5"
	SCANCODE_6 = sdl.SCANCODE_6 // "6"
	SCANCODE_7 = sdl.SCANCODE_7 // "7"
	SCANCODE_8 = sdl.SCANCODE_8 // "8"
	SCANCODE_9 = sdl.SCANCODE_9 // "9"
	SCANCODE_0 = sdl.SCANCODE_0 // "0"

	SCANCODE_RETURN    = sdl.SCANCODE_RETURN    // "Return"
	SCANCODE_ESCAPE    = sdl.SCANCODE_ESCAPE    // "Escape" (the Esc key)
	SCANCODE_BACKSPACE = sdl.SCANCODE_BACKSPACE // "Backspace"
	SCANCODE_TAB       = sdl.SCANCODE_TAB       // "Tab" (the Tab key)
	SCANCODE_SPACE     = sdl.SCANCODE_SPACE     // "Space" (the Space Bar key(s))

	SCANCODE_MINUS        = sdl.SCANCODE_MINUS        // "-"
	SCANCODE_EQUALS       = sdl.SCANCODE_EQUALS       // "="
	SCANCODE_LEFTBRACKET  = sdl.SCANCODE_LEFTBRACKET  // "["
	SCANCODE_RIGHTBRACKET = sdl.SCANCODE_RIGHTBRACKET // "]"
	SCANCODE_BACKSLASH    = sdl.SCANCODE_BACKSLASH    // "\"
	SCANCODE_NONUSHASH    = sdl.SCANCODE_NONUSHASH    // "#" (ISO USB keyboards actually use this code instead of 49 for the same key, but all OSes I've seen treat the two codes identically. So, as an implementor, unless your keyboard generates both of those codes and your OS treats them differently, you should generate SDL_SCANCODE_BACKSLASH instead of this code. As a user, you should not rely on this code because SDL will never generate it with most (all?) keyboards.)
	SCANCODE_SEMICOLON    = sdl.SCANCODE_SEMICOLON    // ";"
	SCANCODE_APOSTROPHE   = sdl.SCANCODE_APOSTROPHE   // "'"
	SCANCODE_GRAVE        = sdl.SCANCODE_GRAVE        // "`"
	SCANCODE_COMMA        = sdl.SCANCODE_COMMA        // ","
	SCANCODE_PERIOD       = sdl.SCANCODE_PERIOD       // "."
	SCANCODE_SLASH        = sdl.SCANCODE_SLASH        // "/"
	SCANCODE_CAPSLOCK     = sdl.SCANCODE_CAPSLOCK     // "CapsLock"
	SCANCODE_F1           = sdl.SCANCODE_F1           // "F1"
	SCANCODE_F2           = sdl.SCANCODE_F2           // "F2"
	SCANCODE_F3           = sdl.SCANCODE_F3           // "F3"
	SCANCODE_F4           = sdl.SCANCODE_F4           // "F4"
	SCANCODE_F5           = sdl.SCANCODE_F5           // "F5"
	SCANCODE_F6           = sdl.SCANCODE_F6           // "F6"
	SCANCODE_F7           = sdl.SCANCODE_F7           // "F7"
	SCANCODE_F8           = sdl.SCANCODE_F8           // "F8"
	SCANCODE_F9           = sdl.SCANCODE_F9           // "F9"
	SCANCODE_F10          = sdl.SCANCODE_F10          // "F10"
	SCANCODE_F11          = sdl.SCANCODE_F11          // "F11"
	SCANCODE_F12          = sdl.SCANCODE_F12          // "F12"
	SCANCODE_PRINTSCREEN  = sdl.SCANCODE_PRINTSCREEN  // "PrintScreen"
	SCANCODE_SCROLLLOCK   = sdl.SCANCODE_SCROLLLOCK   // "ScrollLock"
	SCANCODE_PAUSE        = sdl.SCANCODE_PAUSE        // "Pause" (the Pause / Break key)
	SCANCODE_INSERT       = sdl.SCANCODE_INSERT       // "Insert" (insert on PC, help on some Mac keyboards (but does send code 73, not 117))
	SCANCODE_HOME         = sdl.SCANCODE_HOME         // "Home"
	SCANCODE_PAGEUP       = sdl.SCANCODE_PAGEUP       // "PageUp"
	SCANCODE_DELETE       = sdl.SCANCODE_DELETE       // "Delete"
	SCANCODE_END          = sdl.SCANCODE_END          // "End"
	SCANCODE_PAGEDOWN     = sdl.SCANCODE_PAGEDOWN     // "PageDown"
	SCANCODE_RIGHT        = sdl.SCANCODE_RIGHT        // "Right" (the Right arrow key (navigation keypad))
	SCANCODE_LEFT         = sdl.SCANCODE_LEFT         // "Left" (the Left arrow key (navigation keypad))
	SCANCODE_DOWN         = sdl.SCANCODE_DOWN         // "Down" (the Down arrow key (navigation keypad))
	SCANCODE_UP           = sdl.SCANCODE_UP           // "Up" (the Up arrow key (navigation keypad))

	SCANCODE_NUMLOCKCLEAR = sdl.SCANCODE_NUMLOCKCLEAR // "Numlock" (the Num Lock key (PC) / the Clear key (Mac))
	SCANCODE_KP_DIVIDE    = sdl.SCANCODE_KP_DIVIDE    // "Keypad /" (the / key (numeric keypad))
	SCANCODE_KP_MULTIPLY  = sdl.SCANCODE_KP_MULTIPLY  // "Keypad *" (the * key (numeric keypad))
	SCANCODE_KP_MINUS     = sdl.SCANCODE_KP_MINUS     // "Keypad -" (the - key (numeric keypad))
	SCANCODE_KP_PLUS      = sdl.SCANCODE_KP_PLUS      // "Keypad +" (the + key (numeric keypad))
	SCANCODE_KP_ENTER     = sdl.SCANCODE_KP_ENTER     // "Keypad Enter" (the Enter key (numeric keypad))
	SCANCODE_KP_1         = sdl.SCANCODE_KP_1         // "Keypad 1" (the 1 key (numeric keypad))
	SCANCODE_KP_2         = sdl.SCANCODE_KP_2         // "Keypad 2" (the 2 key (numeric keypad))
	SCANCODE_KP_3         = sdl.SCANCODE_KP_3         // "Keypad 3" (the 3 key (numeric keypad))
	SCANCODE_KP_4         = sdl.SCANCODE_KP_4         // "Keypad 4" (the 4 key (numeric keypad))
	SCANCODE_KP_5         = sdl.SCANCODE_KP_5         // "Keypad 5" (the 5 key (numeric keypad))
	SCANCODE_KP_6         = sdl.SCANCODE_KP_6         // "Keypad 6" (the 6 key (numeric keypad))
	SCANCODE_KP_7         = sdl.SCANCODE_KP_7         // "Keypad 7" (the 7 key (numeric keypad))
	SCANCODE_KP_8         = sdl.SCANCODE_KP_8         // "Keypad 8" (the 8 key (numeric keypad))
	SCANCODE_KP_9         = sdl.SCANCODE_KP_9         // "Keypad 9" (the 9 key (numeric keypad))
	SCANCODE_KP_0         = sdl.SCANCODE_KP_0         // "Keypad 0" (the 0 key (numeric keypad))
	SCANCODE_KP_PERIOD    = sdl.SCANCODE_KP_PERIOD    // "Keypad ." (the . key (numeric keypad))

	SCANCODE_KP_EQUALS = sdl.SCANCODE_KP_EQUALS // "Keypad =" (the = key (numeric keypad))
	SCANCODE_KP_COMMA  = sdl.SCANCODE_KP_COMMA  // "Keypad ," (the Comma key (numeric keypad))

	SCANCODE_KP_00          = sdl.SCANCODE_KP_00          // "Keypad 00" (the 00 key (numeric keypad))
	SCANCODE_KP_000         = sdl.SCANCODE_KP_000         // "Keypad 000" (the 000 key (numeric keypad))
	SCANCODE_KP_LEFTPAREN   = sdl.SCANCODE_KP_LEFTPAREN   // "Keypad (" (the Left Parenthesis key (numeric keypad))
	SCANCODE_KP_RIGHTPAREN  = sdl.SCANCODE_KP_RIGHTPAREN  // "Keypad )" (the Right Parenthesis key (numeric keypad))
	SCANCODE_KP_LEFTBRACE   = sdl.SCANCODE_KP_LEFTBRACE   // "Keypad {" (the Left Brace key (numeric keypad))
	SCANCODE_KP_RIGHTBRACE  = sdl.SCANCODE_KP_RIGHTBRACE  // "Keypad }" (the Right Brace key (numeric keypad))
	SCANCODE_KP_TAB         = sdl.SCANCODE_KP_TAB         // "Keypad Tab" (the Tab key (numeric keypad))
	SCANCODE_KP_BACKSPACE   = sdl.SCANCODE_KP_BACKSPACE   // "Keypad Backspace" (the Backspace key (numeric keypad))
	SCANCODE_KP_POWER       = sdl.SCANCODE_KP_POWER       // "Keypad ^" (the Power key (numeric keypad))
	SCANCODE_KP_PERCENT     = sdl.SCANCODE_KP_PERCENT     // "Keypad %" (the Percent key (numeric keypad))
	SCANCODE_KP_LESS        = sdl.SCANCODE_KP_LESS        // "Keypad <" (the Less key (numeric keypad))
	SCANCODE_KP_GREATER     = sdl.SCANCODE_KP_GREATER     // "Keypad >" (the Greater key (numeric keypad))
	SCANCODE_KP_AMPERSAND   = sdl.SCANCODE_KP_AMPERSAND   // "Keypad &" (the & key (numeric keypad))
	SCANCODE_KP_VERTICALBAR = sdl.SCANCODE_KP_VERTICALBAR // "Keypad |" (the | key (numeric keypad))
	SCANCODE_KP_COLON       = sdl.SCANCODE_KP_COLON       // "Keypad :" (the : key (numeric keypad))
	SCANCODE_KP_HASH        = sdl.SCANCODE_KP_HASH        // "Keypad #" (the # key (numeric keypad))
	SCANCODE_KP_SPACE       = sdl.SCANCODE_KP_SPACE       // "Keypad Space" (the Space key (numeric keypad))
	SCANCODE_KP_AT          = sdl.SCANCODE_KP_AT          // "Keypad @" (the @ key (numeric keypad))
	SCANCODE_KP_EXCLAM      = sdl.SCANCODE_KP_EXCLAM      // "Keypad !" (the ! key (numeric keypad))

	SCANCODE_LCTRL  = sdl.SCANCODE_LCTRL  // "Left Ctrl"
	SCANCODE_LSHIFT = sdl.SCANCODE_LSHIFT // "Left Shift"
	SCANCODE_LALT   = sdl.SCANCODE_LALT   // "Left Alt" (alt, option)
	SCANCODE_LGUI   = sdl.SCANCODE_LGUI   // "Left GUI" (windows, command (apple), meta)
	SCANCODE_RCTRL  = sdl.SCANCODE_RCTRL  // "Right Ctrl"
	SCANCODE_RSHIFT = sdl.SCANCODE_RSHIFT // "Right Shift"
	SCANCODE_RALT   = sdl.SCANCODE_RALT   // "Right Alt" (alt gr, option)
	SCANCODE_RGUI   = sdl.SCANCODE_RGUI   // "Right GUI" (windows, command (apple), meta)
)
