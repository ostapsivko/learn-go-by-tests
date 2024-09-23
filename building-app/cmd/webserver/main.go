package main

import (
	"log"
	"net/http"
	"poker"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatalf("problem creating file system store: %v", err)
	}
	defer close()

	server, err := poker.NewPlayerServer(store)

	if err != nil {
		log.Fatalf("problem creating server: %v", err)
	}

	log.Fatal(http.ListenAndServe(":5000", server))
}
