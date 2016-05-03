package e

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/mewmew/e/ecs"
)

// TODO: Update documentation to reflect the fact that the rendering destination
// does not have to be the screen.

// A RenderSystem periodically renders visible entities to screen.
type RenderSystem interface {
	ecs.System
	// Render renders visible entities to screen. It is invoked once every frame.
	Render(screen *ebiten.Image) error
}
