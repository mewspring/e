// Package ecs provides interfaces for the ECS paradigm.
//
// This package is heavily inspired by engo.io/ecs.
package ecs

// TODO: Extend top-level doc comment.
//
// ECS stands for Entity Component System...

import (
	"sync"
	"time"
)

// TODO: Implement support for prioritizing systems.

// An Engine manages a set of systems, and ensures that each system is updated
// once every frame.
type Engine interface {
	System
	// AddSystem adds the given system to the engine, sorted by priority.
	AddSystem(system System)
	// Systems returns the list of systems managed by the engine.
	Systems() []System
}

// A System represents a system within the ECS paradigm. Systems implement logic
// for dealing with entities possessing specific aspects (i.e. containing
// specific components). For instance, a collision detection system may operate
// on entities containing a "PositionComponent", which provides access to the
// positional information of an entity.
type System interface {
	// Update updates the system. It is invoked once every frame, with dt being
	// the duration since the previous update.
	Update(dt time.Duration) error
}

// A Component represents a component within the ECS paradigm. Components store
// data related to a specific aspect of an entity, such as its position in the
// Cartesian coordinate system.
type Component interface {
	// Type returns the canonical string representation of the component's type.
	// Conventionally this is equivalent to the name of the component's Go type;
	// e.g. "PositionComponent".
	Type() string
}

// An ID represents the unique identifier of an entity within the ECS paradigm.
// An entity is nothing more than a set of components with an ID attached to it.
// Each entity corresponds to a specific entity within the game, such as a
// warrior, a healing potion or a door.
type ID struct {
	// Entity ID.
	id uint64
}

// unique provides access to unique entity IDs.
type unique struct {
	ID
	sync.Mutex
}

// gen generates unique entity IDs.
var gen unique

// NewID returns a new unique ID. It is safe for concurrent use.
func NewID() ID {
	gen.Lock()
	id := gen.ID
	gen.id++
	gen.Unlock()
	return id
}
