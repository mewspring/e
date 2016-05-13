package ecs_test

import (
	"testing"

	"github.com/mewmew/e/ecs"
)

func BenchmarkNewID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ecs.NewID()
	}
}
