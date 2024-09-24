package poker_test

import (
	"net/http"
	"net/http/httptest"
	"poker"
	"testing"
)

func TestRecoringWinsAndRetrievingThem(t *testing.T) {
	database, clearData := poker.CreateTempFile(t, "[]")
	defer clearData()

	store, err := poker.NewFileSystemPlayerStore(database)

	poker.AssertNoError(t, err)

	game := poker.NewTexasHoldem(store, &SpyBlindAlerter{})
	server, err := poker.NewPlayerServer(store, game)
	player := "Pepper"
	poker.AssertNoError(t, err)

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetScoreRequest(player))
		poker.AssertStatusCode(t, response.Code, http.StatusOK)
		poker.AssertResponseBody(t, response.Body.String(), "3")
	})

	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newLeagueRequest())
		poker.AssertStatusCode(t, response.Code, http.StatusOK)

		got := getLeagueFromResponse(t, response.Body)
		want := poker.League{
			{"Pepper", 3},
		}
		poker.AssertLeague(t, got, want)
	})
}
