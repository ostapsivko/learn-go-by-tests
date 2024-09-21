package poker

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

		AssertNoError(t, err)

		got := store.GetLeague()

		want := League{
			{"Andrii", 33},
			{"Azdab", 10},
		}

		AssertLeague(t, got, want)

		got = store.GetLeague()
		AssertLeague(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		database, cleanData := createTempFile(t, `[
		{"Name": "Azdab", "Score": 10},
		{"Name": "Andrii", "Score": 33}]`)
		defer cleanData()

		store, err := NewFileSystemPlayerStore(database)

		AssertNoError(t, err)

		got := store.GetPlayerScore("Azdab")
		want := 10
		AssertScoreEquals(t, got, want)
	})

	t.Run("store wins for existing players", func(t *testing.T) {
		database, cleanData := createTempFile(t, `[
			{"Name": "Azdab", "Score": 10},
			{"Name": "Andrii", "Score": 33}]`)
		defer cleanData()

		store, err := NewFileSystemPlayerStore(database)

		AssertNoError(t, err)

		store.RecordWin("Azdab")

		got := store.GetPlayerScore("Azdab")
		want := 11
		AssertScoreEquals(t, got, want)
	})

	t.Run("store wins for new players", func(t *testing.T) {
		database, cleanData := createTempFile(t, `[
			{"Name": "Azdab", "Score": 10},
			{"Name": "Andrii", "Score": 33}]`)

		defer cleanData()

		store, err := NewFileSystemPlayerStore(database)

		AssertNoError(t, err)

		store.RecordWin("Oleksandr")

		got := store.GetPlayerScore("Oleksandr")
		want := 1
		AssertScoreEquals(t, got, want)
	})

	t.Run("works with an empty file", func(t *testing.T) {
		database, cleanData := createTempFile(t, "")

		defer cleanData()

		_, err := NewFileSystemPlayerStore(database)

		AssertNoError(t, err)
	})

	t.Run("league sorted", func(t *testing.T) {
		database, cleanData := createTempFile(t, `[
			{"Name": "Azdab", "Score": 10},
			{"Name": "Andrii", "Score": 33}]`)

		defer cleanData()

		store, err := NewFileSystemPlayerStore(database)
		AssertNoError(t, err)

		got := store.GetLeague()

		want := League{
			{"Andrii", 33},
			{"Azdab", 10},
		}

		AssertLeague(t, got, want)

		//try again
		got = store.GetLeague()
		AssertLeague(t, got, want)
	})
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
