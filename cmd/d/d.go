// d is a game implemented in e, and an exploration into the ECS paradigm.
package main

import (
	"log"

	"github.com/mewmew/e"
)

func main() {
	if err := e.Main(640, 480, "d"); err != nil {
		log.Fatal(err)
	}
}
