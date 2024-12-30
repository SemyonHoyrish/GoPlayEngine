package resource

import (
	"fmt"
	"github.com/veandco/go-sdl2/ttf"
)

type Font struct {
	path string

	loaded   map[int]bool
	ttfFonts map[int]*ttf.Font
}

func NewFont(path string) *Font {
	return &Font{
		path:     path,
		loaded:   make(map[int]bool),
		ttfFonts: make(map[int]*ttf.Font),
	}
}

func (f *Font) Reload(fontSize int) {
	font, err := ttf.OpenFont(f.path, fontSize)
	if err != nil {
		fmt.Println(fmt.Errorf("cannot reload font (%s): %v", f.path, err))
	} else {
		f.ttfFonts[fontSize] = font
		f.loaded[fontSize] = true
	}
}

func (f *Font) GetTTFFont(fontSize int) *ttf.Font {
	loaded, ok := f.loaded[fontSize]
	if !ok || !loaded {
		f.Reload(fontSize)
	}

	return f.ttfFonts[fontSize]
}
