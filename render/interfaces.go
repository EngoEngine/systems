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

package render

import (
	"github.com/EngoEngine/systems/basic"
	"github.com/EngoEngine/systems/physics"
)

// Getters

// GetAnimationComponent Provides container classes ability to fulfil the interface and be accessed more simply by systems, eg in AddByInterface Methods
func (c *AnimationComponent) GetAnimationComponent() *AnimationComponent {
	return c
}

// GetRenderComponent Provides container classes ability to fulfil the interface and be accessed more simply by systems, eg in AddByInterface Methods
func (c *RenderComponent) GetRenderComponent() *RenderComponent {
	return c
}

// Faces

// AnimationFace allows typesafe Access to an Annonymous child AnimationComponent
type AnimationFace interface {
	GetAnimationComponent() *AnimationComponent
}

// RenderFace allows typesafe access to an anonymous RenderComponent
type RenderFace interface {
	GetRenderComponent() *RenderComponent
}

// Combined for systems

// Animationable is the required interface for AnimationSystem.AddByInterface method
type Animationable interface {
	basic.BasicFace
	AnimationFace
	RenderFace
}

// Renderable is the required interface for the RenderSystem.AddByInterface method
type Renderable interface {
	basic.BasicFace
	RenderFace
	physics.SpaceFace
}

// Not-Ables

// NotAnimationComponent is used to flag an entity as not in the AnimationSystem
// even if it has the proper components
type NotAnimationComponent struct{}

// GetNotAnimationComponent implements the NotAnimationable interface
func (n *NotAnimationComponent) GetNotAnimationComponent() *NotAnimationComponent {
	return n
}

// NotAnimationable is an interface used to flag an entity as not in the
// AnimationSystem even if it has the proper components
type NotAnimationable interface {
	GetNotAnimationComponent() *NotAnimationComponent
}

// NotRenderComponent is used to flag an entity as not in the RenderSystem even
// if it has the proper components
type NotRenderComponent struct{}

// GetNotRenderComponent implements the NotRenderable interface
func (n *NotRenderComponent) GetNotRenderComponent() *NotRenderComponent {
	return n
}

// NotRenderable is an interface used to flag an entity as not in the
// Rendersystem even if it has the proper components
type NotRenderable interface {
	GetNotRenderComponent() *NotRenderComponent
}
