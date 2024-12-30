package resource

import (
	"fmt"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

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

func (i *Image) Reload() {
	surf, err := img.Load(i.path)
	if err != nil {
		fmt.Println(fmt.Errorf("cannot reload image (%s): %v", i.path, err))
	} else {
		i.surface = surf
		i.loaded = true
	}
}

func (i *Image) GetSurface() *sdl.Surface {
	if !i.loaded {
		i.Reload()
	}

	return i.surface
}
