## systems

systems contains ECS implementations of commonly used game systems for use with
[engo](htts://engo.io). These systems aim to be minimal and useful for any game
that uses engo. Improvments, bug reports, and PRs are always welcome! Currently,
this package is a reflection of what's available in engo's `common` package, but
that package will be removed and replaced with this in v1.2

Check out the demos at engo.io/engo/demos to see how they work.

The current package contains ECS systems for:

 * audio
 * rendering
   * animation
   * camera control
   * fonts
   * sprite sheets
   * TMX maps
 * physics
   * mouse detection (clicks, hover, dragging, etc)
   * collision detection

Package basic just contains interfaces for ecs.BasicEntity

#### Breaking Changes since v1.0
