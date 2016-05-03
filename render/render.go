// Package render implements a simple rendering system.
package render

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/mewkiz/pkg/errutil"
	"github.com/mewmew/e"
	"github.com/mewmew/e/ecs"
)

// A System periodically renders visible entities to screen. A visible entity is
// an entity containing a render.Component.
type System struct {
}

// Remove removes the given entity from the system.
func (sys *System) Remove(id ecs.ID) {
}

// Render renders visible entities to screen. It is invoked once every frame.
func (sys *System) Render(screen *ebiten.Image) error {
	if err := screen.Fill(color.RGBA{0, 0, 0, 0xFF}); err != nil {
		return errutil.Err(err)
	}
	return nil
}

// Ensure that System implements the e.RenderSystem interface.
var _ e.RenderSystem = &System{}
