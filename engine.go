package goplayengine

import (
	"fmt"
	"github.com/SemyonHoyrish/GoPlayEngine/core"
	"github.com/SemyonHoyrish/GoPlayEngine/input"
	"github.com/SemyonHoyrish/GoPlayEngine/primitive"
	"github.com/veandco/go-sdl2/gfx"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	"os"
	"sort"
)

// Engine is main structure, which links everything together, and behaves like entry point for your game.
type Engine struct {
	activeScene                   *core.Scene
	activeSceneNoFunctionReported bool
	running                       bool
	started                       bool
	exitCode                      int

	window   *sdl.Window
	renderer *sdl.Renderer

	mouse    *input.Mouse
	keyboard *input.Keyboard

	previousTicks uint64
	deltaTime     uint64

	// TODO: move to engine configuration
	maxEventsPolledPerRender int

	cleanUp func()
}

func NewEngine() *Engine {
	engine := &Engine{
		activeScene:                   nil,
		activeSceneNoFunctionReported: false,
		running:                       false,
		started:                       false,
		exitCode:                      0,
		window:                        nil,
		renderer:                      nil,
		mouse:                         input.NewMouse(),
		keyboard:                      input.NewKeyboard(),
		previousTicks:                 0,
		deltaTime:                     0,

		maxEventsPolledPerRender: 10,
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

	r, err := sdl.CreateRenderer(w, -1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC)
	if err != nil {
		panic(err)
	}

	engine.window = w
	engine.renderer = r

	engine.previousTicks = engine.GetTicks()

	engine.cleanUp = func() {
		engine.renderer.Destroy()
		engine.window.Destroy()
		ttf.Quit()
		sdl.Quit()
	}

	return engine
}

// SetActiveScene sets active scene in the engine, which will be used
// to render text frame
func (e *Engine) SetActiveScene(scene *core.Scene) {
	e.activeScene = scene
	e.activeSceneNoFunctionReported = false
}

// GetActiveScene returns current scene of the Engine
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

			if t != nil {
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
						if size.Width != size.Height {
							fmt.Println(fmt.Errorf("circle width and height differ, possibly trying to override with node size (node id = %d)", objNode.GetID()))
							break
						}
						gfx.FilledCircleColor(
							e.renderer,
							int32(objNode.GetAbsolutePosition().X-size.Width/2),
							int32(objNode.GetAbsolutePosition().Y-size.Height/2),
							int32(size.Width),
							c,
						)
					case primitive.EllipsePrimitive:
						panic("TODO implement")
					case primitive.LinePrimitive:
						npos := objNode.GetPosition()
						vec := prim.(primitive.Line).Definition
						if npos != vec.From {
							fmt.Println(fmt.Errorf("line start and node position differ (node id=%d). For more info see docs for primitive.Line", objNode.GetID()))
							break
						}
						e.renderer.DrawLineF(vec.From.X, vec.From.Y, vec.To.X, vec.To.Y)
					}
				} else if t.GetImage() != nil {
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
				} else {
					fmt.Println(fmt.Errorf("node has empty texture (node id = %d)", node.GetID()))
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

// GetMouse returns Engine instance of input.Mouse, the only initialized instance you should use
func (e *Engine) GetMouse() *input.Mouse {
	return e.mouse
}

// GetKeyboard returns Engine instance of input.Keyboard, the only initialized instance you should use
func (e *Engine) GetKeyboard() *input.Keyboard {
	return e.keyboard
}

// GetTicks returns number of milliseconds since SDL was initialized in NewEngine function
func (e *Engine) GetTicks() uint64 {
	return sdl.GetTicks64()
}

// GetDeltaTime returns difference between two frames in milliseconds.
func (e *Engine) GetDeltaTime() uint64 {
	return e.deltaTime
}

// Run creates window and start rendering activeScene.
//
// It is required to call Run in main thread
func (e *Engine) Run() {
	if e.started {
		return
	}

	e.running = true
	e.started = true

	for e.running {
		// Handle events
		{
			iters := 0
			for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
				iters += 1

				//TODO:
				switch event.(type) {
				case *sdl.QuitEvent:
					e.running = false

				case *sdl.MouseButtonEvent:
					e.GetMouse().SetLastEvent(event.(*sdl.MouseButtonEvent))

				case *sdl.KeyboardEvent:
					e.GetKeyboard().SetLastEvent(event.(*sdl.KeyboardEvent))

				}

				if iters >= e.maxEventsPolledPerRender {
					break
				}
			}
		}

		// Render
		{
			curTicks := e.GetTicks()
			e.deltaTime = curTicks - e.previousTicks
			e.previousTicks = curTicks

			if e.activeScene.GetUpdateFunction() != nil {
				e.activeScene.GetUpdateFunction()()
			} else if !e.activeSceneNoFunctionReported {
				fmt.Println(fmt.Errorf("no update function on scene ID=(%d)", e.activeScene.GetID()))
				e.activeSceneNoFunctionReported = true
			}

			e.GetMouse().ApplyDeferred()
			e.GetKeyboard().ApplyDeferred()

			nodes := e.activeScene.GetAllNodes()

			color := e.GetActiveScene().GetBackgroundColor()
			e.renderer.SetDrawColor(color.R, color.G, color.B, color.A)
			e.renderer.Clear()
			e.render(nodes)
			e.renderer.Present()
		}
	}

	e.cleanUp()
	os.Exit(e.exitCode)
}

// Exit tries to gracefully shutdown game engine and exit with provided code.
func (e *Engine) Exit(code int) {
	e.exitCode = code
	e.running = false
}
