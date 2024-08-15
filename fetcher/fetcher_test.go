package fetcher

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyFetcher struct {
	response  string
	cancelled bool
	t         *testing.T
}

func (s *SpyFetcher) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return s.response
}

func (s *SpyFetcher) Cancel() {
	s.cancelled = true
}

func (s *SpyFetcher) assertWasCancelled() {
	s.t.Helper()
	if !s.cancelled {
		s.t.Error("fetcher was not told to cancel")
	}
}

func (s *SpyFetcher) assertWasNotCancelled() {
	s.t.Helper()
	if s.cancelled {
		s.t.Error("fetcher was not expected to cancel")
	}
}

func TestServer(t *testing.T) {
	t.Run("returns data from store without cancellation", func(t *testing.T) {
		data := "hello world"
		fetcher := &SpyFetcher{response: data, t: t}
		srv := Server(fetcher)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		srv.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf(`got "%s", want "%s"`, response.Body.String(), data)
		}

		fetcher.assertWasNotCancelled()
	})

	t.Run("tells fetcher to cancel work if request is cancelled", func(t *testing.T) {
		data := "hello world"
		fetcher := &SpyFetcher{response: data, t: t}
		srv := Server(fetcher)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancelCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)

		request = request.WithContext(cancelCtx)

		response := httptest.NewRecorder()
		srv.ServeHTTP(response, request)

		fetcher.assertWasCancelled()
	})
}
