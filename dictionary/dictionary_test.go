package dictionary

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "This is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		assertStrings(t, "This is just a test", got)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")

		if err == nil {
			t.Fatal("expected to get an error")
		}

		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	dictionary.Add("test", "This is just a test")

	got, err := dictionary.Search("test")
	want := "This is just a test"

	if err != nil {
		t.Fatal("should find added word: ", want)
	}

	assertStrings(t, want, got)
}

func assertError(t testing.TB, want, got error) {
	t.Helper()

	if want != got {
		t.Errorf("got error %q, want error %q", got, want)
	}
}

func assertStrings(t testing.TB, want, got string) {
	t.Helper()

	if got != want {
		t.Errorf("want %q, got %q", want, got)
	}
}
