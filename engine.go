package e

import (
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/mewkiz/pkg/errutil"
	"github.com/mewmew/e/ecs"
)

// e represents the game engine.
var e = &engine{}

// An engine manages a set of systems, and ensures that each system is updated
// once every frame.
type engine struct {
	// prev tracks the time of the previous update.
	prev time.Time
	// Active systems.
	systems []ecs.System
}

// AddSystem adds the given system to the engine, sorted by priority.
func (e *engine) AddSystem(system ecs.System) {
	// TODO: Implement handling of system priority.
	e.systems = append(e.systems, system)
}

// Systems returns the list of systems managed by the engine.
func (e *engine) Systems() []ecs.System {
	return e.systems
}

// Update updates each system of the engine.
//
// NOTE: The underlying ebiten engine is responsible for updating the game
// engine, do not call this method directly.
func (e *engine) Update(dt time.Duration) error {
	panic("e.engine.Update: invalid call to Update. The underlying ebiten engine is responsible for updating the game engine, do not call this method directly.")
}

// update updates each system, and renders with each rendering system of the
// engine. It is invoked once every frame from the underlying ebiten engine,
// which is guaranteed to run at 60 FPS.
func (e *engine) update(screen *ebiten.Image) error {
	// Calculate the duration since the previous update.
	dt := delta(e.prev)
	e.prev = time.Now()

	// Update active systems.
	for _, system := range e.systems {
		if err := system.Update(dt); err != nil {
			return errutil.Err(err)
		}
	}

	// TODO: Skip rendering if close to frame time limit. Dropping a frame is
	// preferable to introducing screen tearing. Therefore, either skip or run
	// all rendering systems.

	// Render with active rendering systems.
	for _, system := range e.systems {
		if system, ok := system.(RenderingSystem); ok {
			if err := system.Render(screen); err != nil {
				return errutil.Err(err)
			}
		}
	}

	// TODO: Add support for input handling systems, audio processing systems.

	return nil
}

// delta returns the duration since the previous update. A duration of 16.6 ms
// (i.e. the duration of 1 frame at 60 FPS) is returned for the first update.
func delta(prev time.Time) time.Duration {
	if prev.IsZero() {
		const fps = 60
		return time.Second / fps
	}
	return time.Since(e.prev)
}

// Ensure that engine implements the ecs.Engine interface.
var _ ecs.Engine = &engine{}
