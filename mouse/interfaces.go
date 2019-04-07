// interfaces.go is intended to provide a simple means of adding components to
// each system
//
// Getters
//
// These are added functions to each class to allow them to meet the interfaces
// we use with AddByInterface methods on each system.
//
// Faces
//
// The interfaces that end in "Face" are all met by a specific component, which
// can be composed into an Entity. The word Get is used because, otherwise it
// would collide with the name of the object, when stored anonymously in a
// parent entity.
//
// Ables
//
// The interfaces that end in "able" are those required by a specific system,
// and if an an object meets this interface it can be added to that system.
//
// Note: *able* is used not *er* because they don't really do thing anything.
//
// Note: The names have not been contracted for consistency, the interface is
// *Collisionable* not *Collidable*.
//
// Not-Ables
//
// The Not-Ables are interfaces of components used to flag entities to not add to the system,
// for use with the ecs.World.AddSystemInterface

package mouse

import (
	"github.com/EngoEngine/systems/basic"
	"github.com/EngoEngine/systems/physics"
	"github.com/EngoEngine/systems/render"
)

// Getters

// GetMouseComponent Provides container classes ability to fulfil the interface and be accessed more simply by systems, eg in AddByInterface Methods
func (c *MouseComponent) GetMouseComponent() *MouseComponent {
	return c
}

// Faces

// MouseFace allows typesafe access to an anonymouse child MouseComponent
type MouseFace interface {
	GetMouseComponent() *MouseComponent
}

// Mouseable is the required interface for the MouseSystem.AddByInterface method
type Mouseable interface {
	basic.BasicFace
	MouseFace
	physics.SpaceFace
	render.RenderFace
}

// Not-Ables

// NotMouseComponent is used to flag an entity as not in the AudioSystem even if
// it has the proper components
type NotMouseComponent struct{}

// GetNotMouseComponent implements the NotMousable interface
func (n *NotMouseComponent) GetNotMouseComponent() *NotMouseComponent {
	return n
}

// NotMouseable is an interface used to flag an entity as not in the MouseSystem
// even if it has the proper components
type NotMouseable interface {
	GetNotMouseComponent() *NotMouseComponent
}
