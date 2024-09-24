package poker_test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"poker"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

var (
	dummyGame = &poker.GameSpy{}
)

func TestGETPlayers(t *testing.T) {
	store := &poker.StubPlayerStore{
		Scores: map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
	}

	server := mustMakePlayerServer(t, store, dummyGame)

	t.Run("returns Pepper's score", func(t *testing.T) {
		request := newGetScoreRequest("Pepper")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		poker.AssertResponseBody(t, response.Body.String(), "20")
		poker.AssertStatusCode(t, response.Code, http.StatusOK)
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		request := newGetScoreRequest("Floyd")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		poker.AssertResponseBody(t, response.Body.String(), "10")
		poker.AssertStatusCode(t, response.Code, http.StatusOK)
	})

	t.Run("returns 404 on missing players", func(t *testing.T) {
		request := newGetScoreRequest("Apollo")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		poker.AssertStatusCode(t, response.Code, http.StatusNotFound)
	})
}

func TestStoreWins(t *testing.T) {
	store := &poker.StubPlayerStore{
		Scores: map[string]int{},
	}

	server := mustMakePlayerServer(t, store, dummyGame)

	t.Run("it records wins when POST", func(t *testing.T) {
		player := "Pepper"

		request := newPostWinRequest(player)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		poker.AssertStatusCode(t, response.Code, http.StatusAccepted)
		poker.AssertPlayerWin(t, store, player)
	})
}

func TestLeague(t *testing.T) {

	t.Run("it returns league table as JSON", func(t *testing.T) {
		wantedLeague := poker.League{
			{"Azdab", 20},
			{"Andrii", 40},
			{"Oleksandr", 88},
		}

		store := &poker.StubPlayerStore{League: wantedLeague}
		server := mustMakePlayerServer(t, store, dummyGame)

		request := newLeagueRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := getLeagueFromResponse(t, response.Body)
		poker.AssertStatusCode(t, response.Code, http.StatusOK)
		poker.AssertLeague(t, got, wantedLeague)
		poker.AssertContentType(t, response.Result().Header.Get("content-type"), poker.JsonContentType)
	})
}

func TestGame(t *testing.T) {
	t.Run("GET /game returns 200", func(t *testing.T) {
		server := mustMakePlayerServer(t, &poker.StubPlayerStore{}, dummyGame)

		request := newGameRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		poker.AssertStatusCode(t, response.Code, http.StatusOK)
	})

	t.Run("start a game with 3 players, send a blind alert and declare Azdab as a winner", func(t *testing.T) {
		winner := "Azdab"
		wantedAlert := "Blind is 100"
		game := &poker.GameSpy{BlindAlert: []byte(wantedAlert)}

		playerServer := mustMakePlayerServer(t, dummyPlayerStore, game)
		server := httptest.NewServer(playerServer)
		defer server.Close()

		wsURL := "ws" + strings.TrimPrefix(server.URL, "http") + "/ws"
		ws := mustDialWs(t, wsURL)
		defer ws.Close()

		sendWsMessage(t, ws, "3")
		sendWsMessage(t, ws, winner)

		assertGameStartedWith(t, game, 3)
		assertWinner(t, game, winner)

		within(t, 1*time.Second, func() { assertWebsocketGotMsg(t, ws, wantedAlert) })

	})
}

func assertWebsocketGotMsg(t testing.TB, conn *websocket.Conn, want string) {
	t.Helper()

	_, msg, _ := conn.ReadMessage()

	if string(msg) != want {
		t.Errorf("got alert %s, want %s", string(msg), want)
	}
}

func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func newPostWinRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func newLeagueRequest() *http.Request {
	request, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return request
}

func newGameRequest() *http.Request {
	request, _ := http.NewRequest(http.MethodGet, "/game", nil)
	return request
}

func getLeagueFromResponse(t testing.TB, response io.Reader) poker.League {
	t.Helper()

	league, err := poker.NewLeague(response)

	if err != nil {
		t.Fatalf("Unable to parse response from server %q into the slice of Player, '%v'", response, err)
	}

	return league
}

func mustMakePlayerServer(t testing.TB, store poker.PlayerStore, game *poker.GameSpy) *poker.PlayerServer {
	t.Helper()

	server, err := poker.NewPlayerServer(store, game)

	if err != nil {
		t.Fatal("problem creating a server", err)
	}

	return server
}

func mustDialWs(t testing.TB, url string) *websocket.Conn {
	ws, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		t.Fatalf("could not open a ws connection on %s %v", url, err)
	}
	return ws
}

func sendWsMessage(t testing.TB, conn *websocket.Conn, message string) {
	if err := conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
		t.Fatalf("could not send message over ws connection %v", err)
	}
}

func within(t testing.TB, d time.Duration, assert func()) {
	t.Helper()

	done := make(chan struct{}, 1)

	go func() {
		assert()
		done <- struct{}{}
	}()

	select {
	case <-time.After(d):
		t.Error("timed out")
	case <-done:
	}
}

func retryUntil(d time.Duration, f func() bool) bool {
	deadline := time.Now().Add(d)
	for time.Now().Before(deadline) {
		if f() {
			return true
		}
	}

	return false
}
