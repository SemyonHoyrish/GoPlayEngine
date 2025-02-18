# [v0.6.0]
### NEW
Sublayers system for nodes to make sure they render in expected order. No changes for user level interactions.



# [v0.5.1] Line primitive change
Line primitive now has new definition and nodes with line primitive as texture are ignored when auto overlap is being built.



# [v0.5.0] Auto Overlaps & Node system rewrite
### CHANGES
Node system has been rewritten to make them much simpler. Now, there is only one struct `Node` that represents
any node in Game Engine, based in `NodeType` field. At the same time, user interactions with nodes a pretty much the same, with very small changes, if any.

### NEW
Auto Overlaps for `Node`.
Auto Overlap system will build a `ComposedOverlap` for a `Node` and all of its children by creating an `Overlap` with `CalculatedSize` of the node and position in center of that node for every node in selected node hierarchy.

There is new functions for `Node`:
- `AutoOverlap(bool)` - to enable or disable auto overlap.
- `AutoOverlapEnabled() bool` - to get if auto overlap is enabled for this node.



# [v0.4.0] Input handling & major fix
### NEW
#### input handling
- Implemented input.Keyboard
- For input.Mouse & input.Keyboard now available functions `ButtonPressed`, `ButtonDown`, `ButtonUp`. For input.Keyboard are introduced some rebindings of sdl scancode for easier access, however, direct usage of sdl.SCANCODE_* constants are not limited in any way.
#### engine time
- New functions for engine: `GetTicks`, `GetTimeDelta` (see docs).

### FIX
Fixed issue that prevented any render to happen, which was caused by not complying with thread safety instructions of the SDL library. Now application should stably run on any architectures and configurations supported by SDL Library.



# [v0.3.0] Overlaps
### NEW
#### Overlaps system
This system allows to attach `Overlap` or `ComposedOverlap` to a nodes, and then determine whether if they overlaps overlapping each other or not. Additionally, it provides `MouseOver` function to check whether mouse is currently hovering an overlap.
- Overlap
- ComposedOverlap
- BaseNode.{SetOverlap, GetOverlap, RemoveOverlap}

#### Primitives
- circle
- line (has special requirements for node position, see docs)



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
