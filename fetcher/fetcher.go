package fetcher

import (
	"fmt"
	"net/http"
)

type Fetcher interface {
	Fetch() string
}

func Server(fetcher Fetcher) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, fetcher.Fetch())
	}
}
