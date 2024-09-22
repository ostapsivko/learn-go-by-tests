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
	dummyStdOut      = &bytes.Buffer{}
)

type GameSpy struct {
	StartedWith  int
	FinishedWith string
	StartCalled  bool
	FinishCalled bool
}

func (g *GameSpy) Start(numberOfPlayers int) {
	g.StartedWith = numberOfPlayers
	g.StartCalled = true
}

func (g *GameSpy) Finish(winner string) {
	g.FinishedWith = winner
	g.FinishCalled = true
}

func TestCLI(t *testing.T) {
	t.Run("starts game with 5 players and record Azdab win from user input", func(t *testing.T) {
		in := userSends("5", "Azdab wins")
		game := &GameSpy{}
		out := &bytes.Buffer{}

		cli := poker.NewCLI(in, out, game)
		cli.PlayPoker()

		assertMessagesSentToUser(t, out, poker.PlayerPrompt)
		assertWinner(t, game, "Azdab")
		assertGameStartedWith(t, game, 5)
	})

	t.Run("start game with 6 players and record Andrii win from user input", func(t *testing.T) {
		in := userSends("6", "Andrii wins")
		game := &GameSpy{}

		cli := poker.NewCLI(in, dummyStdOut, game)
		cli.PlayPoker()

		assertWinner(t, game, "Andrii")
		assertGameStartedWith(t, game, 6)
	})

	t.Run("prints an error when a non-numeric value is entered and does not start the game", func(t *testing.T) {
		out := &bytes.Buffer{}
		in := userSends("Abcdefg")
		game := &GameSpy{}

		cli := poker.NewCLI(in, out, game)
		cli.PlayPoker()

		assertGameNotStarted(t, game)
		assertMessagesSentToUser(t, out, poker.PlayerPrompt, poker.BadPlayerInputErrMsg)
	})

	t.Run("handles incorrect win message", func(t *testing.T) {
		out := &bytes.Buffer{}
		game := &GameSpy{}
		in := userSends("3", "Sinner is a killer")

		cli := poker.NewCLI(in, out, game)
		cli.PlayPoker()

		assertGameNotFinished(t, game)
		assertMessagesSentToUser(t, out, poker.PlayerPrompt, poker.BadWinnerInputErrMsg)
	})
}

func assertMessagesSentToUser(t testing.TB, out *bytes.Buffer, messages ...string) {
	t.Helper()

	want := strings.Join(messages, "")
	got := out.String()
	if got != want {
		t.Errorf("got %q sent to out but expected %+v", got, messages)
	}
}

func assertGameStartedWith(t testing.TB, game *GameSpy, players int) {
	t.Helper()

	if players != game.StartedWith {
		t.Errorf("players started with does not match expected value - got %q, want %q", game.StartedWith, players)
	}
}

func assertWinner(t testing.TB, game *GameSpy, winner string) {
	t.Helper()

	if winner != game.FinishedWith {
		t.Errorf("winner with does not match expected value - got %q, want %q", game.FinishedWith, winner)
	}
}

func assertGameNotStarted(t testing.TB, game *GameSpy) {
	t.Helper()

	if game.StartCalled {
		t.Errorf("game should not have started")
	}
}

func assertGameNotFinished(t testing.TB, game *GameSpy) {
	t.Helper()

	if game.FinishCalled {
		t.Errorf("game should not have ended")
	}
}

func userSends(messages ...string) *bytes.Buffer {
	input := strings.Join(messages, "\n")
	return bytes.NewBuffer([]byte(input))
}
