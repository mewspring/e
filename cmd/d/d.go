// d is a game implemented in e, and an exploration into the ECS paradigm.
package main

import (
	"log"

	"github.com/mewmew/e"
	"github.com/mewmew/e/render"
)

func main() {
	// Add rendering system to game engine.
	renderer := &render.System{}
	e.AddSystem(renderer)

	// Launch game.
	if err := e.Main(640, 480, "d"); err != nil {
		log.Fatal(err)
	}
}
