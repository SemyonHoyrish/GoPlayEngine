package resource

import (
	"fmt"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

// Image used to link path to file on disk and its loaded content,
// have to be initialized with NewFont
type Image struct {
	path string

	loaded  bool
	surface *sdl.Surface
}

func NewImage(path string) *Image {
	return &Image{
		path:    path,
		loaded:  false,
		surface: nil,
	}
}

// Reload load content of file and stores it as pixel data in memory.
// Reload called automatically if it was not called before.
func (i *Image) Reload() {
	surf, err := img.Load(i.path)
	if err != nil {
		fmt.Println(fmt.Errorf("cannot reload image (%s): %v", i.path, err))
	} else {
		i.surface = surf
		i.loaded = true
	}
}

// GetSurface is an internal function, returns image representation used to render it.
func (i *Image) GetSurface() *sdl.Surface {
	if !i.loaded {
		i.Reload()
	}

	return i.surface
}
