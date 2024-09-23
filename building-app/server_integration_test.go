package poker

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecoringWinsAndRetrievingThem(t *testing.T) {
	database, clearData := createTempFile(t, "[]")
	defer clearData()

	store, err := NewFileSystemPlayerStore(database)

	AssertNoError(t, err)

	server, err := NewPlayerServer(store)
	player := "Pepper"
	AssertNoError(t, err)

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetScoreRequest(player))
		AssertStatusCode(t, response.Code, http.StatusOK)
		AssertResponseBody(t, response.Body.String(), "3")
	})

	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newLeagueRequest())
		AssertStatusCode(t, response.Code, http.StatusOK)

		got := getLeagueFromResponse(t, response.Body)
		want := League{
			{"Pepper", 3},
		}
		AssertLeague(t, got, want)
	})
}
