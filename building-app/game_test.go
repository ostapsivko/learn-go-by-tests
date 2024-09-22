package poker_test

import (
	"bytes"
	"fmt"
	"poker"
	"strings"
	"testing"
	"time"
)

func TestGame_Start(t *testing.T) {
	t.Run("schedules printing of blind values", func(t *testing.T) {
		store := &poker.StubPlayerStore{}
		blindAlerter := &SpyBlindAlerter{}
		game := poker.NewGame(store, blindAlerter)

		game.Start(5)

		cases := []scheduledAlert{
			{0 * time.Second, 100},
			{10 * time.Minute, 200},
			{20 * time.Minute, 300},
			{30 * time.Minute, 400},
			{40 * time.Minute, 500},
			{50 * time.Minute, 600},
			{60 * time.Minute, 800},
			{70 * time.Minute, 1000},
			{80 * time.Minute, 2000},
			{90 * time.Minute, 4000},
			{100 * time.Minute, 8000},
		}

		checkSchedulingCases(t, cases, blindAlerter)
	})

	t.Run("prompts user to enter the number of players", func(t *testing.T) {
		output := &bytes.Buffer{}
		input := strings.NewReader("7\n")
		blindAlerter := &SpyBlindAlerter{}
		game := poker.NewGame(dummyPlayerStore, blindAlerter)

		cli := poker.NewCLI(input, output, game)
		cli.PlayPoker()

		got := output.String()
		want := poker.PlayerPrompt

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

		cases := []scheduledAlert{
			{0 * time.Second, 100},
			{12 * time.Minute, 200},
			{24 * time.Minute, 300},
			{36 * time.Minute, 400},
		}

		checkSchedulingCases(t, cases, blindAlerter)
	})
}

func TestGame_Finish(t *testing.T) {
	store := &poker.StubPlayerStore{}
	game := poker.NewGame(store, dummySpyAlerter)
	winner := "Oleksandr"

	game.Finish(winner)
	poker.AssertPlayerWin(t, store, winner)
}

func checkSchedulingCases(t *testing.T, cases []scheduledAlert, alerter *SpyBlindAlerter) {
	for i, want := range cases {
		t.Run(fmt.Sprint(want), func(t *testing.T) {
			if len(alerter.alerts) <= i {
				t.Fatalf("alert %d was not scheduled for %v", i, alerter.alerts)
			}

			got := alerter.alerts[i]
			assertScheduledAlert(t, got, want)
		})
	}
}

func assertScheduledAlert(t testing.TB, got, want scheduledAlert) {
	t.Helper()

	if got.amount != want.amount {
		t.Errorf("got amount %d, want %d", got.amount, want.amount)
	}

	if got.at != want.at {
		t.Errorf("got scheduled time of %v, want %v", got.at, want.at)
	}
}
