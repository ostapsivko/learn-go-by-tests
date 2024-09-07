package interactions_test

import (
	"go-specs-greet/domain/interactions"
	"go-specs-greet/specifications"
	"testing"

	"github.com/alecthomas/assert/v2"
)

func TestGreet(t *testing.T) {
	specifications.GreetSpecification(t, specifications.GreetAdapter(interactions.Greet))

	t.Run("empty name defaults to `World`", func(t *testing.T) {
		assert.Equal(t, "Hello, World!", interactions.Greet(""))
	})
}
