package poker

import (
	"strings"
	"testing"
)

func TestCLI(t *testing.T) {
	in := strings.NewReader("Azdab wins\n")
	playerStore := &StubPlayerStore{}

	cli := &CLI{playerStore, in}
	cli.PlayPoker()

	assertPlayerWin(t, playerStore, "Azdab")
}

func assertPlayerWin(t testing.TB, store *StubPlayerStore, want string) {
	t.Helper()

	if len(store.winCalls) != 1 {
		t.Fatal("expected a win call but did not get any")
	}

	if store.winCalls[0] != want {
		t.Errorf("did not record correct winner, got %s want %s", store.winCalls[0], want)
	}
}
