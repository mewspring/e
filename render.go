package e

import "github.com/hajimehoshi/ebiten"

// TODO: Update documentation to reflect the fact that the rendering destination
// does not have to be the screen.

// A RenderingSystem represents a system capable of rendering visible entities
// to screen.
type RenderingSystem interface {
	// Render renders visible entities to screen. It is invoked once every frame.
	Render(screen *ebiten.Image) error
}
