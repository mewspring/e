// Package e provides an experimenal game engine, and an exploration into the
// ECS paradigm.
//
// This project is heavily inspired by the engo and ebiten game engines.
package e

// TODO: Add to top-level doc comment.
//
// ECS stands for Entity Component System...

import (
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/mewkiz/pkg/errutil"
)

// TODO: Define minimal API for
//
//    1. Window handling
//    2. Rendering
//    3. Input handling
//    4. Message handling
//    5. An Entity Component System

// Main runs the game, and ensures that each system is updated 60 times per
// second, even if a rendering frame is skipped.
//
// NOTE: This function must be called from the main thread.
func Main(width, height int, title string) error {
	world := &world{}
	if err := ebiten.Run(world.update, width, height, 1, title); err != nil {
		return errutil.Err(err)
	}
	return nil
}

// world manages the active systems, and ensures that each system is updated
// once every frame.
type world struct {
	// prev tracks the time of the previous call to update.
	prev time.Time
	// Active systems.
	systems []System
}

// update is invoked once every frame from the underlying ebiten engine, which
// is guaranteed to run at 60 FPS. It calls the update function of each system,
// which handle game logic, input handling, rendering and audio.
func (world *world) update(screen *ebiten.Image) error {
	// Calculate the duration since the previous call to update.
	const fps = 60
	dt := time.Second / fps
	var zero time.Time
	if world.prev != zero {
		dt = time.Since(world.prev)
	}
	world.prev = time.Now()

	// Update the active systems.
	for _, system := range world.systems {
		system.Update(dt)
	}

	return nil
}

// An ID represents the unique identifier of an entity within the ECS paradigm.
// An entity is nothing more than a set of components with an ID attached to it.
// Each entity corresponds to a specific entity within the game, such as a
// warrior, a healing potion or a door.
type ID uint64

// A Component represents a component within the ECS paradigm. Components store
// data related to a specific aspect of an entity, such as its position in the
// Cartesian coordinate system.
type Component interface {
	// Type returns the canonical string representation of the component's type.
	// Conventionally this is equivalent to the name of the component's Go type;
	// e.g. "PositionComponent".
	Type() string
}

// A System represents a system within the ECS paradigm. Systems implement logic
// for dealing with entities possessing specific aspects (i.e. containing
// specific components). For instance, a collision detection system may operate
// on entities containing a "PositionComponent", which provides access to the
// positional information of an entity.
type System interface {
	// Update is invoked once every frame, with dt being the duration since the
	// previous call to update.
	Update(dt time.Duration) error
	// RemoveEntity removes the given entity from the system.
	RemoveEntity(id ID)
}
