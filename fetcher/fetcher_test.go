package fetcher

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type SpyFetcher struct {
	response string
}

func (s *SpyFetcher) Fetch() string {
	return s.response
}

func TestServer(t *testing.T) {
	data := "hello world"
	srv := Server(&SpyFetcher{data})

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()

	srv.ServeHTTP(response, request)

	if response.Body.String() != data {
		t.Errorf(`got "%s", want "%s"`, response.Body.String(), data)
	}
}
