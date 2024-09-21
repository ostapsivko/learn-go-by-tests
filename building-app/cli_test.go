package poker_test

import (
	"poker"
	"strings"
	"testing"
)

func TestCLI(t *testing.T) {
	t.Run("record Azdab win from user input", func(t *testing.T) {
		in := strings.NewReader("Azdab wins\n")
		playerStore := &poker.StubPlayerStore{}

		cli := poker.NewCLI(playerStore, in)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Azdab")
	})

	t.Run("record Andrii from user input", func(t *testing.T) {
		in := strings.NewReader("Andrii wins\n")
		playerStore := &poker.StubPlayerStore{}

		cli := poker.NewCLI(playerStore, in)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Andrii")
	})
}
