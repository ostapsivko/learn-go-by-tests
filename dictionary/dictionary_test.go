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
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "This is just a test"

		_ = dictionary.Add(word, definition)

		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "This is just a test"
		dictionary := Dictionary{word: definition}

		err := dictionary.Add(word, "new definition")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
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

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word: ", err)
	}

	assertStrings(t, definition, got)
}
