// Package e provides an experimenal game engine, and an exploration into the
// ECS paradigm.
//
// This project is heavily inspired by the engo and ebiten game engines.
package e

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/mewkiz/pkg/errutil"
	"github.com/mewmew/e/ecs"
)

// Main runs the game, and ensures that each system is updated 60 times per
// second, even if a rendering frame is skipped.
//
// NOTE: This function must be called from the main thread.
func Main(width, height int, title string) error {
	if err := ebiten.Run(e.update, width, height, 1, title); err != nil {
		return errutil.Err(err)
	}
	return nil
}

// AddSystem adds the given system to the game engine, sorted by priority.
func AddSystem(system ecs.System) {
	e.AddSystem(system)
}

// Systems returns the list of systems managed by the game engine.
func Systems() []ecs.System {
	return e.Systems()
}
