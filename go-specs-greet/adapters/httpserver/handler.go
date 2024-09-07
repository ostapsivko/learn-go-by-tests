package httpserver

import (
	"fmt"
	"go-specs-greet/domain/interactions"
	"net/http"
)

func HandleGreet(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Fprint(w, interactions.Greet(name))
}

func HandleCurse(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Fprint(w, interactions.Curse(name))
}
