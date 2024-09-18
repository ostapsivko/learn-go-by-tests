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
}
