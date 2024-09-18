package main

import (
	"strings"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("league from a reader", func(t *testing.T) {
		database := strings.NewReader(`[
		{"Name": "Azdab", "Score": 10},
		{"Name": "Andrii", "Score": 33}]`)

		store := FileSystemPlayerStore{database}

		got := store.GetLeague()

		want := []Player{
			{"Azdab", 10},
			{"Andrii", 33},
		}

		assertLeague(t, got, want)

		got = store.GetLeague()
		assertLeague(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		database := strings.NewReader(`[
		{"Name": "Azdab", "Score": 10},
		{"Name": "Andrii", "Score": 33}]`)

		store := FileSystemPlayerStore{database}

		got := store.GetPlayerScore("Azdab")

		want := 10
		assertScoreEquals(t, got, want)
	})
}

func assertScoreEquals(t *testing.T, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
