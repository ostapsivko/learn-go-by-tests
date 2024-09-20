package main

import (
	"os"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("league from a reader", func(t *testing.T) {
		database, cleanData := createTempFile(t, `[
		{"Name": "Azdab", "Score": 10},
		{"Name": "Andrii", "Score": 33}]`)
		defer cleanData()

		store, err := NewFileSystemPlayerStore(database)

		assertNoError(t, err)

		got := store.GetLeague()

		want := League{
			{"Andrii", 33},
			{"Azdab", 10},
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

		store, err := NewFileSystemPlayerStore(database)

		assertNoError(t, err)

		got := store.GetPlayerScore("Azdab")
		want := 10
		assertScoreEquals(t, got, want)
	})

	t.Run("store wins for existing players", func(t *testing.T) {
		database, cleanData := createTempFile(t, `[
			{"Name": "Azdab", "Score": 10},
			{"Name": "Andrii", "Score": 33}]`)
		defer cleanData()

		store, err := NewFileSystemPlayerStore(database)

		assertNoError(t, err)

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

		store, err := NewFileSystemPlayerStore(database)

		assertNoError(t, err)

		store.RecordWin("Oleksandr")

		got := store.GetPlayerScore("Oleksandr")
		want := 1
		assertScoreEquals(t, got, want)
	})

	t.Run("works with an empty file", func(t *testing.T) {
		database, cleanData := createTempFile(t, "")

		defer cleanData()

		_, err := NewFileSystemPlayerStore(database)

		assertNoError(t, err)
	})

	t.Run("league sorted", func(t *testing.T) {
		database, cleanData := createTempFile(t, `[
			{"Name": "Azdab", "Score": 10},
			{"Name": "Andrii", "Score": 33}]`)

		defer cleanData()

		store, err := NewFileSystemPlayerStore(database)
		assertNoError(t, err)

		got := store.GetLeague()

		want := League{
			{"Andrii", 33},
			{"Azdab", 10},
		}

		assertLeague(t, got, want)

		//try again
		got = store.GetLeague()
		assertLeague(t, got, want)
	})
}

func assertScoreEquals(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func createTempFile(t testing.TB, initialData string) (*os.File, func()) {
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
