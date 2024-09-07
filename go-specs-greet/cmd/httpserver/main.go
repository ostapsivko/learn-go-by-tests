package main

import (
	"go-specs-greet/adapters/httpserver"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/greet", httpserver.HandleGreet)
	mux.HandleFunc("/curse", httpserver.HandleCurse)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
