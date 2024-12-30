# [v0.2.0] Documentation
### NEW
Everything (close to that) now documented in code, so package now has documentation on [pkg.go.dev](https://pkg.go.dev)

[Link to docs](https://pkg.go.dev/github.com/SemyonHoyrish/GoPlayEngine)

### CHANGES
Now each `Scene` has own `updateFunction` which will be called before rendering every frame of that scene.
Before we had one `updateFunction` for entire game, which was not ideal.



# [v0.1.0] Basic structure of the game engine

### NEW

#### engine

#### basic
  - Point
  - Size
  - Vector

#### core
  - BaseNode
  - ObjectNode
  - TextNode
  - Scene
  - Texture

#### input
  - Mouse

#### primitive
  - Color
  - PrimitiveInterface
  - Rectangle primitive

#### resource
  - Font
  - Image
