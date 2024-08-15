package fetcher

import (
	"fmt"
	"net/http"
)

type Fetcher interface {
	Fetch() string
	Cancel()
}

func Server(fetcher Fetcher) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		data := make(chan string, 1)

		go func() {
			data <- fetcher.Fetch()
		}()

		select {
		case d := <-data:
			fmt.Fprint(w, d)
		case <-ctx.Done():
			fetcher.Cancel()
		}
	}
}
