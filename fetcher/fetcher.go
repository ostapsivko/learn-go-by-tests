package fetcher

import (
	"context"
	"fmt"
	"net/http"
)

type Fetcher interface {
	Fetch(ctx context.Context) (string, error)
}

func Server(fetcher Fetcher) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		result, err := fetcher.Fetch(ctx)

		if err != nil {
			return
		}

		fmt.Fprint(w, result)
	}
}
