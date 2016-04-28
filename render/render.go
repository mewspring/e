// Package render implements a simple rendering system.
package render

import (
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/mewmew/e"
	"github.com/mewmew/e/ecs"
)

// A System implements logic for rendering entities containing a
// render.Component.
type System struct {
}

// Update updates the rendering system. It is invoked once every frame, with dt
// being the duration since the previous update.
func (sys *System) Update(dt time.Duration) error {
	return nil
}

// Render renders visible entities to screen. It is invoked once every frame.
func (sys *System) Render(screen *ebiten.Image) error {
	return nil
}

// Ensure that System implements the ecs.System interface.
var _ ecs.System = &System{}

// Ensure that System implements the e.Renderer interface.
var _ e.Renderer = &System{}
