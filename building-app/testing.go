package poker

import (
	"reflect"
	"testing"
)

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	league   League
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	return s.scores[name]
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func (s *StubPlayerStore) GetLeague() League {
	return s.league
}

func AssertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q, want %q", got, want)
	}
}

func AssertStatusCode(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("status codes do not match, got %d, want %d", got, want)
	}
}

func AssertLeague(t testing.TB, got, want League) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func AssertContentType(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response did not have content-type header %v, got %v", want, got)
	}
}

func AssertNoError(t testing.TB, err error) {
	t.Helper()

	if err != nil {
		t.Fatalf("did not expect an error but got one, %v", err)
	}
}

func AssertPlayerWin(t testing.TB, store *StubPlayerStore, want string) {
	t.Helper()

	if len(store.winCalls) != 1 {
		t.Fatal("expected a win call but did not get any")
	}

	if store.winCalls[0] != want {
		t.Errorf("did not record correct winner, got %s want %s", store.winCalls[0], want)
	}
}

func AssertScoreEquals(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
