package main

import (
	"io"
	"os"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("league from a reader", func(t *testing.T) {
		database, cleanData := createTempFile(t, `[
		{"Name": "Azdab", "Score": 10},
		{"Name": "Andrii", "Score": 33}]`)
		defer cleanData()

		store := NewFileSystemPlayerStore(database)

		got := store.GetLeague()

		want := League{
			{"Azdab", 10},
			{"Andrii", 33},
		}

		assertLeague(t, got, want)

		got = store.GetLeague()
		assertLeague(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		database, cleanData := createTempFile(t, `[
		{"Name": "Azdab", "Score": 10},
		{"Name": "Andrii", "Score": 33}]`)
		defer cleanData()

		store := NewFileSystemPlayerStore(database)

		got := store.GetPlayerScore("Azdab")
		want := 10
		assertScoreEquals(t, got, want)
	})

	t.Run("store wins for existing players", func(t *testing.T) {
		database, cleanData := createTempFile(t, `[
			{"Name": "Azdab", "Score": 10},
			{"Name": "Andrii", "Score": 33}]`)
		defer cleanData()

		store := NewFileSystemPlayerStore(database)

		store.RecordWin("Azdab")

		got := store.GetPlayerScore("Azdab")
		want := 11
		assertScoreEquals(t, got, want)
	})

	t.Run("store wins for new players", func(t *testing.T) {
		database, cleanData := createTempFile(t, `[
			{"Name": "Azdab", "Score": 10},
			{"Name": "Andrii", "Score": 33}]`)

		defer cleanData()

		store := NewFileSystemPlayerStore(database)

		store.RecordWin("Oleksandr")

		got := store.GetPlayerScore("Oleksandr")
		want := 1
		assertScoreEquals(t, got, want)
	})
}

func assertScoreEquals(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func createTempFile(t testing.TB, initialData string) (io.ReadWriteSeeker, func()) {
	t.Helper()

	tmpfile, err := os.CreateTemp("", "db")

	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}

	tmpfile.Write([]byte(initialData))

	removeFile := func() {
		tmpfile.Close()
		os.Remove(tmpfile.Name())
	}

	return tmpfile, removeFile
}
