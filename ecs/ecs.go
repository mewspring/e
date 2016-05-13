// Package ecs provides interfaces for the Entity Component System (ECS)
// paradigm.
//
// This package is heavily inspired by engo.io/ecs.
//
// The ECS paradigm aims to decouple distinct domains (e.g. rendering, input
// handling, AI) from one another, through a composition of independent
// components. The core concepts of ECS are described below.
//
//
// Entities
//
// An entity is simply a set of components with a unique ID attached to it,
// nothing more. In particular, an entity has no logic attached to it and stores
// no data explicitly (except for the ID).
//
// Each entity corresponds to a specific entity within the game, such as a
// character, an item, or a spell.
//
//
// Components
//
// A component stores the raw data related to a specific aspect of an entity,
// nothing more. In particular, a component has no logic attached to it.
//
// Different aspects may include the position, animation graphics, or input
// actions of an entity.
//
//
// Systems
//
// A system implements logic for processing entities possessing components of
// the same aspects as the system.
//
// For instance, an animation system may render entities possessing animation
// components.
package ecs

import (
	"sync/atomic"
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

// A System implements logic for processing entities possessing components of
// the same aspects as the system.
//
// By convention, systems provide an Add method for adding entities and their
// associated components to the system; e.g.
//
//    Add(id ID, pos *Position)
type System interface {
	// Remove removes the given entity from the system.
	Remove(id ID)
}

// An UpdateSystem periodically updates relevant aspects of entities.
type UpdateSystem interface {
	System
	// Update updates the system. It is invoked by the engine once every frame,
	// with dt being the duration since the previous update.
	Update(dt time.Duration) error
}

// An Entity is simply a set of components with a unique ID attached to it,
// nothing more.
type Entity interface {
	// ID returns the unique identifier of the entity.
	ID() ID
}

// An ID represents the unique identifier of an entity.
type ID struct {
	// Entity ID.
	id uint64
}

// unique provides access to unique entity IDs, starting at 1.
var unique uint64

// NewID returns a new unique ID. It is safe for concurrent use.
func NewID() ID {
	return ID{id: atomic.AddUint64(&unique, 1)}
}
