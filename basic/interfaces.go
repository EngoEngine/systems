package basic // import "github.com/EngoEngine/systems/basic"

import "github.com/EngoEngine/ecs"

// BasicFace is the means of accessing the ecs.BasicEntity class , it also has the ID method, to simplify, finding an item within a system
type BasicFace interface {
	ID() uint64
	GetBasicEntity() *ecs.BasicEntity
}
