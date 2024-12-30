package goplayengine

import (
	"fmt"
	"github.com/SemyonHoyrish/GoPlayEngine/core"
	"github.com/SemyonHoyrish/GoPlayEngine/input"
	"github.com/SemyonHoyrish/GoPlayEngine/primitive"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	"os"
	"sort"
	"sync"
)

type Engine struct {
	activeScene *core.Scene
	running     bool
	started     bool
	exitCode    int

	wg *sync.WaitGroup

	window   *sdl.Window
	renderer *sdl.Renderer

	mouse *input.Mouse

	cleanUp func()
}

func NewEngine() *Engine {
	engine := &Engine{
		activeScene: nil,
		running:     false,
		started:     false,
		exitCode:    0,
		wg:          &sync.WaitGroup{},
		window:      nil,
		renderer:    nil,
		mouse:       nil,
	}

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}

	if err := ttf.Init(); err != nil {
		panic(err)
	}

	w, err := sdl.CreateWindow("GAME", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 720, 480, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}

	r, err := sdl.CreateRenderer(w, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}

	engine.window = w
	engine.renderer = r

	engine.wg = &sync.WaitGroup{}

	engine.cleanUp = func() {
		engine.renderer.Destroy()
		engine.window.Destroy()
		ttf.Quit()
		sdl.Quit()
	}

	return engine
}

func (e *Engine) SetActiveScene(scene *core.Scene) {
	e.activeScene = scene
}

func (e *Engine) GetActiveScene() *core.Scene {
	return e.activeScene
}

func (e *Engine) render(nodes []core.BaseNodeInterface) {
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].GetLayer() < nodes[j].GetLayer()
	})

	for _, node := range nodes {
		switch node.GetNodeType() {
		case core.ObjectNodeType:
			objNode := node.(core.ObjectNodeInterface)

			t := objNode.GetTexture()

			size := node.GetSize()
			if size.Width == 0 && size.Height == 0 {
				size = t.GetSize()
				if size.Width == 0 && size.Height == 0 {
					err := fmt.Errorf("size of node (id=%d) is zero still after resolution", objNode.GetID())
					fmt.Println(err)
				}
			}

			if t.GetPrimitive() != nil {
				// TODO: handle errors
				prim := t.GetPrimitive()

				c := prim.GetColor()
				e.renderer.SetDrawColor(c.R, c.G, c.B, c.A)

				switch prim.GetPrimitiveType() {
				case primitive.RectanglePrimitive:
					e.renderer.FillRectF(&sdl.FRect{
						X: objNode.GetAbsolutePosition().X - size.Width/2,
						Y: objNode.GetAbsolutePosition().Y - size.Height/2,
						W: size.Width,
						H: size.Height,
					})
				case primitive.CirclePrimitive:
					panic("TODO implement")
				case primitive.EllipsePrimitive:
					panic("TODO implement")
				case primitive.LinePrimitive:
					panic("TODO implement")
				}
			} else {
				image := t.GetImage()
				if image.GetSurface() == nil {
					fmt.Println(fmt.Errorf("no image for node id %d", objNode.GetID()))
				} else {
					tx, err := e.renderer.CreateTextureFromSurface(image.GetSurface())
					_ = err // TODO:

					e.renderer.CopyF(tx, nil, &sdl.FRect{
						X: objNode.GetAbsolutePosition().X - size.Width/2,
						Y: objNode.GetAbsolutePosition().Y - size.Height/2,
						W: size.Width,
						H: size.Height,
					})

					tx.Destroy()
				}
			}

			childNodes := objNode.GetChildren()
			baseChildNodes := make([]core.BaseNodeInterface, len(childNodes))
			for i, chn := range childNodes {
				baseChildNodes[i] = chn
			}
			e.render(baseChildNodes)

		case core.TextNodeType:
			textNode := node.(core.TextNodeInterface)
			font := textNode.GetFont()
			ttfFont := font.GetTTFFont(textNode.GetTextSize())
			// TODO: handle error
			textContent := textNode.GetText()
			//surf, _ := ttfFont.RenderUTF8Solid(textContent, textNode.GetColor())
			surf, _ := ttfFont.RenderUTF8Blended(textNode.GetText(), textNode.GetColor())

			size := node.GetSize()
			if size.Width == 0 && size.Height == 0 {
				w, h, _ := ttfFont.SizeUTF8(textContent)
				if w == 0 && h == 0 {
					err := fmt.Errorf("size of node (id=%d) is zero still after resolution", textNode.GetID())
					fmt.Println(err)
				} else {
					size.Width = float32(w)
					size.Height = float32(h)
				}
			}

			tx, _ := e.renderer.CreateTextureFromSurface(surf)
			e.renderer.CopyF(tx, nil, &sdl.FRect{
				X: textNode.GetAbsolutePosition().X,
				Y: textNode.GetAbsolutePosition().Y,
				W: size.Width,
				H: size.Height,
			})

			tx.Destroy()

		case core.InputNodeType:
			panic("TODO implement")

		case core.BaseNodeType:
			err := fmt.Errorf("BaseNodeType cannot be rendered, Node ID: %d", node.GetID())
			fmt.Println(err)

		}
	}
}

func (e *Engine) GetMouse() *input.Mouse {
	if e.mouse == nil {
		e.mouse = &input.Mouse{}
	}

	return e.mouse
}

// Run creates window and start rendering activeScene.
// updateFunc is called before render of every frame
//
// It is required to call Run in main thread
func (e *Engine) Run(updateFunc func()) {
	if e.started {
		return
	}

	e.running = true
	e.started = true

	e.wg.Add(1)
	go func() {
		for e.running {
			updateFunc()

			nodes := e.activeScene.GetAllNodes()

			color := e.GetActiveScene().GetBackgroundColor()
			e.renderer.SetDrawColor(color.R, color.G, color.B, color.A)
			e.renderer.Clear()
			e.render(nodes)
			e.renderer.Present()
		}
		e.wg.Done()
	}()

	for e.running {
		event := sdl.PollEvent()
		if event != nil {
			switch event.(type) {
			case *sdl.QuitEvent:
				e.running = false

				//TODO:
			}
		}
	}

	e.wg.Wait()

	e.cleanUp()
	os.Exit(e.exitCode)
}

func (e *Engine) Exit(code int) {
	e.exitCode = code
	e.running = false
}