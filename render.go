package e

import "github.com/hajimehoshi/ebiten"

// A Renderer represents a system capable of rendering entities to screen.
type Renderer interface {
	// Render renders visible entities to screen. It is invoked once every frame.
	Render(screen *ebiten.Image) error
}
