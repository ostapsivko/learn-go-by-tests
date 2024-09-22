package poker_test

import (
	"bytes"
	"fmt"
	"poker"
	"strings"
	"testing"
	"time"
)

type scheduledAlert struct {
	at     time.Duration
	amount int
}

func (s *scheduledAlert) String() string {
	return fmt.Sprintf("%d chips at %v", s.amount, s.amount)
}

type SpyBlindAlerter struct {
	alerts []scheduledAlert
}

func (s *SpyBlindAlerter) ScheduleAlertAt(duration time.Duration, amount int) {
	s.alerts = append(s.alerts, scheduledAlert{duration, amount})
}

var (
	dummySpyAlerter  = &SpyBlindAlerter{}
	dummyPlayerStore = &poker.StubPlayerStore{}
	dummyStdIn       = &bytes.Buffer{}
	dummyStdOut      = &bytes.Buffer{}
)

type GameSpy struct {
	StartedWith  int
	FinishedWith string
}

func (g *GameSpy) Start(numberOfPlayers int) {
	g.StartedWith = numberOfPlayers
}

func (g *GameSpy) Finish(winner string) {
	g.FinishedWith = winner
}

func TestCLI(t *testing.T) {
	t.Run("prompts user to enter the number of players and starts the game", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("7\n")
		game := &GameSpy{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		gotPrompt := stdout.String()
		wantPrompt := poker.PlayerPrompt

		if gotPrompt != wantPrompt {
			t.Errorf("got %q, want %q", gotPrompt, wantPrompt)
		}

		if game.StartedWith != 7 {
			t.Errorf("wanted StartedWith to be equal 7 but got %d", game.StartedWith)
		}
	})

	t.Run("start game with 5 players and record Azdab win from user input", func(t *testing.T) {
		in := strings.NewReader("5\nAzdab wins\n")
		game := &GameSpy{}

		cli := poker.NewCLI(in, dummyStdOut, game)
		cli.PlayPoker()

		want := "Azdab"
		got := game.FinishedWith

		if got != want {
			t.Errorf("want winner to be %s, got %s", want, got)
		}

		if game.StartedWith != 5 {
			t.Errorf("got %q, want %q", game.StartedWith, 5)
		}
	})

	t.Run("start game with 6 players and record Andrii win from user input", func(t *testing.T) {
		in := strings.NewReader("6\nAndrii wins\n")
		game := &GameSpy{}

		cli := poker.NewCLI(in, dummyStdOut, game)
		cli.PlayPoker()

		players := 6
		winner := "Andrii"
		got := game.FinishedWith

		if got != winner {
			t.Errorf("want winner to be %s, got %s", winner, got)
		}

		if game.StartedWith != players {
			t.Errorf("got %q, want %q", game.StartedWith, players)
		}
	})
}
